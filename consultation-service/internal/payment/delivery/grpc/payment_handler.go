package grpc

import (
	"context"
	paymentPb "ehSehat/consultation-service/internal/payment/delivery/grpc/pb"
	"ehSehat/consultation-service/internal/payment/domain"
	"ehSehat/libs/utils"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// message PaymentLog {
//   string id = 1;
//   string payment_id = 2;
//   google.protobuf.StringValue response = 3; // Use StringValue to allow nullability
//   google.protobuf.Timestamp created_at = 4;
//   google.protobuf.Timestamp updated_at = 5;
// }

// message PaymentRequest {
//   string consultation_id = 1;
//   double amount = 2;
//   string method = 3; // e.g., "credit_card", "bank_transfer"
// }

// message PaymentResponse {
//   string id = 1;
//   string consultation_id = 2;
//   google.protobuf.Timestamp consultation_date = 3;
//   string patient_id = 4;
//   google.protobuf.StringValue patient_name = 5;
//   string doctor_id = 6;
//   google.protobuf.StringValue doctor_name = 7;
//   double amount = 8;
//   google.protobuf.StringValue method = 9;
//   google.protobuf.StringValue gateway = 10;
//   repeated PaymentLog payment_log = 11;
//   google.protobuf.StringValue status = 12;
//   google.protobuf.StringValue created_by = 13;
//   google.protobuf.StringValue created_name = 14;
//   google.protobuf.StringValue created_email = 15;
//   google.protobuf.StringValue created_role = 16;
//   google.protobuf.Timestamp created_at = 17;
//   google.protobuf.StringValue updated_by = 18;
//   google.protobuf.StringValue updated_name = 19;
//   google.protobuf.StringValue updated_email = 20;
//   google.protobuf.StringValue updated_role = 21;
//   google.protobuf.Timestamp updated_at = 22;
// }

// message PaymentUpdateRequest {
//   string id = 1; // unique identifier for the payment
//   double amount = 2; // updated amount
//   string status = 3; // updated status
// }

// service PaymentService {
//   rpc CreatePaymentGRPC(PaymentRequest) returns (PaymentResponse);
//   rpc GetPaymentByIdGRPC(google.protobuf.StringValue) returns (PaymentResponse);
//   rpc UpdatePaymentGRPC(PaymentUpdateRequest) returns (PaymentResponse);
// }

type paymentHandler struct {
	paymentPb.UnimplementedPaymentServiceServer
	app domain.PaymentService
}

func NewPaymentHandler(app domain.PaymentService) *paymentHandler {
	return &paymentHandler{
		app: app,
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
