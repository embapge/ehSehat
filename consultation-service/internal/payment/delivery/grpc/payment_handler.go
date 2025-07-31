package grpc

import (
	"context"
	paymentPb "ehSehat/consultation-service/internal/payment/delivery/grpc/pb"
	"ehSehat/consultation-service/internal/payment/domain"
	"ehSehat/libs/utils"
	"ehSehat/libs/utils/rabbitmqown"
	"encoding/json"
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type paymentHandler struct {
	paymentPb.UnimplementedPaymentServiceServer
	app domain.PaymentService
	ch  *amqp.Channel
}

func NewPaymentHandler(app domain.PaymentService, ch *amqp.Channel) *paymentHandler {
	return &paymentHandler{
		app: app,
		ch:  ch,
	}
}

func (h *paymentHandler) CreatePaymentGRPC(ctx context.Context, req *paymentPb.PaymentRequest) (*paymentPb.PaymentResponse, error) {
	if req == nil {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Request cannot be nil"))
	}
	if req.ConsultationId == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Consultation ID is required"))
	}
	if req.Amount <= 0 {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Amount must be greater than zero"))
	}

	payment := &domain.CreatePaymentRequest{
		ConsultationID: req.ConsultationId,
		Amount:         req.Amount,
	}

	paymentResponse, err := h.app.CreatePayment(ctx, payment)
	if err != nil {
		return nil, utils.GRPCErrorToHTTPError(err)
	}

	var pbPaymentLogs []*paymentPb.PaymentLog
	for _, log := range paymentResponse.PaymentLogs {
		var responseStr string
		if str, ok := log.Response.(string); ok {
			responseStr = str
		} else if log.Response != nil {
			responseStr = fmt.Sprintf("%v", log.Response)
		}
		pbPaymentLogs = append(pbPaymentLogs, &paymentPb.PaymentLog{
			Id:        log.ID,
			PaymentId: log.PaymentID,
			Response:  wrapperspb.String(responseStr),
		})
	}

	// Example: Extract "invoice_url" from the marshaled response if needed
	var invoiceURL string
	if len(paymentResponse.PaymentLogs) > 0 {
		if paymentResponse.PaymentLogs[0].Response != nil {
			responseBytes, err := json.Marshal(paymentResponse.PaymentLogs[0].Response)
			if err == nil {
				var responseMap map[string]interface{}
				if err := json.Unmarshal(responseBytes, &responseMap); err == nil {
					if url, ok := responseMap["invoice_url"].(string); ok {
						invoiceURL = url
					}
				}
			}
		}
	}

	consultationContext, _ := json.Marshal(paymentResponse)
	payload := rabbitmqown.NotificationPayload{
		Channel: "email",
		// Recipient:     consultation.Patient.ID, // Assuming recipient is the patient ID
		Recipient:     "baratagusti.bg@gmail.com", // Assuming recipient is the patient ID
		TemplateName:  "paymentCreated",           // Example template name
		Subject:       "Payment Created!",
		Body:          fmt.Sprintf("Your payment has been created. Please refer to this link to paid consultation: %v, sebesar %v", invoiceURL, paymentResponse.Amount),
		SourceService: "consultationService",
		Context:       consultationContext, // Additional context can be added here if needed
		Status:        "pending",           // Initial status
		ErrorMessage:  "",                  // No error message initially
		RetryCount:    0,                   // Initial retry count
	}

	payloadBytes, _ := json.Marshal(payload)

	err = h.ch.Publish(
		"",                        // default exchange
		os.Getenv("RABBIT_QUEUE"), // routing key (queue name)
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payloadBytes,
		},
	)

	if err != nil {
		return nil, err
	}

	return &paymentPb.PaymentResponse{
		Id:               paymentResponse.ID,
		ConsultationId:   paymentResponse.ConsultationID,
		ConsultationDate: timestamppb.New(*paymentResponse.ConsultationDate),
		PatientId:        paymentResponse.PatientID,
		PatientName:      wrapperspb.String(*paymentResponse.PatientName),
		DoctorId:         paymentResponse.DoctorID,
		DoctorName:       wrapperspb.String(*paymentResponse.DoctorName),
		Amount:           paymentResponse.Amount,
		Method:           wrapperspb.String(paymentResponse.Method),
		Gateway:          wrapperspb.String(paymentResponse.Gateway),
		Status:           wrapperspb.String(paymentResponse.Status),
		PaymentLog:       pbPaymentLogs,
	}, nil
}

func (h *paymentHandler) GetPaymentByIdGRPC(ctx context.Context, req *wrapperspb.StringValue) (*paymentPb.PaymentResponse, error) {
	if req == nil || req.GetValue() == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Payment ID is required"))
	}

	payment, err := h.app.FindByIDPayment(ctx, req.GetValue())
	if err != nil {
		return nil, utils.GRPCErrorToHTTPError(err)
	}

	return &paymentPb.PaymentResponse{
		Id:             payment.ID,
		ConsultationId: payment.ConsultationID,
		Amount:         payment.Amount,
		Method:         wrapperspb.String(payment.Method),
		Gateway:        wrapperspb.String(payment.Gateway),
		Status:         wrapperspb.String(payment.Status),
	}, nil
}

func (h *paymentHandler) HandlePaymentWebhookGRPC(ctx context.Context, req *paymentPb.PaymentWebhookRequest) (*wrapperspb.StringValue, error) {
	if req == nil || req.ExternalId == "" || req.PaymentId == "" || req.EventType == "" {
		return nil, utils.GRPCErrorToHTTPError(utils.NewBadRequestError("Invalid webhook request"))
	}

	webhook := &domain.PaymentWebhook{
		ExternalID: req.ExternalId,
		PaymentID:  req.PaymentId,
		EventType:  req.EventType,
		Payload:    req.Payload.GetValue(),
	}

	err := h.app.HandlePaymentWebhook(ctx, webhook)
	if err != nil {
		return nil, utils.GRPCErrorToHTTPError(err)
	}

	if webhook.EventType == "PAID" {
		// webhook.Payload is a string, so unmarshal it to get the amount
		var payloadMap map[string]interface{}
		if err := json.Unmarshal([]byte(webhook.Payload), &payloadMap); err != nil {
			return nil, utils.NewBadRequestError("Invalid payload format")
		}
		amountFloat, ok := payloadMap["amount"].(float64)
		if !ok {
			return nil, utils.NewBadRequestError("Amount field missing or invalid in payload")
		}

		paymentContext, _ := json.Marshal(payloadMap)
		payload := rabbitmqown.NotificationPayload{
			Channel: "email",
			// Recipient:     consultation.Patient.ID, // Assuming recipient is the patient ID
			Recipient:     "baratagusti.bg@gmail.com", // Assuming recipient is the patient ID
			TemplateName:  "paymentPaid",              // Example template name
			Subject:       "Payment Paid!",
			Body:          fmt.Sprintf("Thanks for consultation with us. Your payment Rp. %v has been processed successfully.", amountFloat),
			SourceService: "consultationService",
			Context:       paymentContext, // Additional context can be added here if needed
			Status:        "pending",      // Initial status
			ErrorMessage:  "",             // No error message initially
			RetryCount:    0,              // Initial retry count
		}

		payloadBytes, _ := json.Marshal(payload)

		err = h.ch.Publish(
			"",                        // default exchange
			os.Getenv("RABBIT_QUEUE"), // routing key (queue name)
			false,                     // mandatory
			false,                     // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        payloadBytes,
			},
		)

		if err != nil {
			return nil, err
		}
	}

	return wrapperspb.String("Webhook processed successfully"), nil
}
