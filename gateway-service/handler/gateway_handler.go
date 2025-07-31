package handler

import (
	"context"
	"ehSehat/gateway-service/config"
	allPb "ehSehat/gateway-service/handler/pb"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"

	// "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type GatewayHandler struct {
	GRPC *config.GRPCClients
}

func NewGatewayHandler(grpcClients *config.GRPCClients) *GatewayHandler {
	return &GatewayHandler{GRPC: grpcClients}
}

func (h *GatewayHandler) ProxyToAuthService(c echo.Context) error {
	ctx := c.Request().Context()
	method := c.Request().Method
	path := c.Request().URL.Path

	switch method {
	case http.MethodPost:
		if strings.HasPrefix(path, "/login") {
			var req allPb.LoginRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.AuthClient.Login(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/register") {
			var req allPb.RegisterRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.AuthClient.Register(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		}
	default:
		return c.NoContent(http.StatusMethodNotAllowed)
	}

	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *GatewayHandler) ProxyToConsultationService(c echo.Context) error {
	ctx := ForwardTSUserHeadersToGRPC(c)
	method := c.Request().Method
	path := c.Request().URL.Path

	switch method {
	case http.MethodPost:
		if strings.HasPrefix(path, "/consultations") {
			var req allPb.ConsultationRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.ConsultationClient.CreateConsultation(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		}
	case http.MethodGet:
		if strings.HasPrefix(path, "/consultations/") && len(path) > len("/consultations/") {
			id := strings.TrimPrefix(path, "/consultations/")
			resp, err := h.GRPC.ConsultationClient.FindByIDConsultation(ctx, &allPb.ConsultationIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	case http.MethodPut:
		if strings.HasPrefix(path, "/consultations/") && len(path) > len("/consultations/") {
			id := strings.TrimPrefix(path, "/consultations/")
			var req allPb.ConsultationRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			req.Id = id
			resp, err := h.GRPC.ConsultationClient.UpdateConsultation(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	}

	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *GatewayHandler) ProxyToPaymentService(c echo.Context) error {
	ctx := ForwardTSUserHeadersToGRPC(c)
	method := c.Request().Method
	path := c.Request().URL.Path

	switch method {
	case http.MethodPost:
		if strings.HasPrefix(path, "/payments") {
			var req allPb.PaymentRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.PaymentClient.CreatePaymentGRPC(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		}
	case http.MethodGet:
		if strings.HasPrefix(path, "/payments/") && len(path) > len("/payments/") {
			id := strings.TrimPrefix(path, "/payments/")
			resp, err := h.GRPC.PaymentClient.GetPaymentByIdGRPC(ctx, wrapperspb.String(id))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	default:
		return c.NoContent(http.StatusMethodNotAllowed)
	}

	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *GatewayHandler) ProxyToPaymentWebhook(c echo.Context) error {
	ctx := ForwardTSUserHeadersToGRPC(c)
	path := c.Request().URL.Path
	if c.Request().Header.Get("X-CALLBACK-TOKEN") != os.Getenv("XENDIT_WEBHOOK_TOKEN") {
		return c.JSON(http.StatusForbidden, map[string]string{"error": "invalid request"})
	}

	if strings.HasPrefix(path, "/payments/xendit-webhook") {
		fmt.Println("Received xendit payment webhook:")
		// Bind seluruh payload ke map agar fleksibel
		var payload map[string]interface{}
		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
		}

		// Ambil field utama dari payload
		id, _ := payload["id"].(string)
		externalID, _ := payload["external_id"].(string)

		// Event type bisa diambil dari status atau field lain sesuai kebutuhan
		eventType, _ := payload["status"].(string)

		// Marshal seluruh payload ke JSON string untuk dikirim sebagai payload
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to marshal payload"})
		}

		req := allPb.PaymentWebhookRequest{
			ExternalId: id,
			PaymentId:  externalID,
			EventType:  eventType,
			Payload:    wrapperspb.String(string(payloadBytes)),
		}

		fmt.Println("Forwarding webhook to gRPC service:", req.ExternalId, req.PaymentId, eventType)
		fmt.Println("\n\nPayload:", req.Payload.GetValue())

		resp, err := h.GRPC.PaymentClient.HandlePaymentWebhookGRPC(ctx, &req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, resp)
	}

	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *GatewayHandler) ProxyToClinicDataService(c echo.Context) error {
	ctx := ForwardTSUserHeadersToGRPC(c)
	method := c.Request().Method
	path := c.Request().URL.Path
	switch method {
	case http.MethodPost:
		if strings.HasPrefix(path, "/clinics/patients") {
			var req allPb.CreatePatientRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.ClinicDataClient.CreatePatient(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		} else if strings.HasPrefix(path, "/clinics/doctors") {
			var req allPb.CreateDoctorRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.ClinicDataClient.CreateDoctor(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		} else if strings.HasPrefix(path, "/clinics/specializations") {
			var req allPb.CreateSpecializationRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.ClinicDataClient.CreateSpecialization(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		} else if strings.HasPrefix(path, "/clinics/rooms") {
			var req allPb.CreateRoomRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.ClinicDataClient.CreateRoom(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		} else if strings.HasPrefix(path, "/clinics/schedule-fixed") {
			var req allPb.CreateScheduleFixedRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.ClinicDataClient.CreateScheduleFixed(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		}
	case http.MethodGet:
		// Berikan method untuk GetAllPatients, GetAllDoctors, GetAllSpecializations, GetAllRooms, GetAllScheduleFixed
		if path == "/clinics/patients" {
			resp, err := h.GRPC.ClinicDataClient.GetAllPatients(ctx, &allPb.Empty{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if path == "/clinics/doctors" {
			resp, err := h.GRPC.ClinicDataClient.GetAllDoctors(ctx, &allPb.Empty{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if path == "/clinics/specializations" {
			resp, err := h.GRPC.ClinicDataClient.GetAllSpecializations(ctx, &allPb.Empty{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if path == "/clinics/rooms" {
			resp, err := h.GRPC.ClinicDataClient.GetAllRooms(ctx, &allPb.Empty{})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/patients/") && len(path) > len("/clinics/patients/") {
			id := strings.TrimPrefix(path, "/clinics/patients/")
			resp, err := h.GRPC.ClinicDataClient.GetPatientByID(ctx, &allPb.GetPatientByIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/doctors/") && len(path) > len("/clinics/doctors/") {
			id := strings.TrimPrefix(path, "/clinics/doctors/")
			resp, err := h.GRPC.ClinicDataClient.GetDoctorByID(ctx, &allPb.GetDoctorByIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/specializations/") && len(path) > len("/clinics/specializations/") {
			id := strings.TrimPrefix(path, "/clinics/specializations/")
			resp, err := h.GRPC.ClinicDataClient.GetSpecializationByID(ctx, &allPb.GetSpecializationByIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/rooms/") && len(path) > len("/clinics/rooms/") {
			id := strings.TrimPrefix(path, "/clinics/rooms/")
			resp, err := h.GRPC.ClinicDataClient.GetRoomByID(ctx, &allPb.GetRoomByIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/schedule-fixed/") && len(path) > len("/clinics/schedule-fixed/") {
			id := strings.TrimPrefix(path, "/clinics/schedule-fixed/")
			resp, err := h.GRPC.ClinicDataClient.GetFixedSchedulesByDoctorID(ctx, &allPb.GetFixedSchedulesByDoctorIDRequest{DoctorId: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	case http.MethodPut:
		if strings.HasPrefix(path, "/clinics/patients/") && len(path) > len("/clinics/patients/") {
			id := strings.TrimPrefix(path, "/clinics/patients/")
			var req allPb.UpdatePatientRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			req.Id = id
			resp, err := h.GRPC.ClinicDataClient.UpdatePatient(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/doctors/") && len(path) > len("/clinics/doctors/") {
			id := strings.TrimPrefix(path, "/clinics/doctors/")
			var req allPb.UpdateDoctorRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			req.Id = id
			resp, err := h.GRPC.ClinicDataClient.UpdateDoctor(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/specializations/") && len(path) > len("/clinics/specializations/") {
			id := strings.TrimPrefix(path, "/clinics/specializations/")
			var req allPb.UpdateSpecializationRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			req.Id = id
			resp, err := h.GRPC.ClinicDataClient.UpdateSpecialization(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/clinics/schedule-fixed/") && len(path) > len("/clinics/schedule-fixed/") {
			id := strings.TrimPrefix(path, "/clinics/schedule-fixed/")
			var req allPb.UpdateScheduleFixedRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			req.Id = id
			resp, err := h.GRPC.ClinicDataClient.UpdateScheduleFixed(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	default:
		return c.NoContent(http.StatusMethodNotAllowed)
	}
	return c.NoContent(http.StatusMethodNotAllowed)
}

func (h *GatewayHandler) ProxyToAppointmentService(c echo.Context) error {
	ctx := ForwardTSUserHeadersToGRPC(c)
	method := c.Request().Method
	path := c.Request().URL.Path

	switch method {
	case http.MethodGet:
		if strings.HasPrefix(path, "/appointments/") && len(path) > len("/appointments/") {
			id := strings.TrimPrefix(path, "/appointments/")
			var idUint uint32
			_, err := fmt.Sscanf(id, "%d", &idUint)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid appointment id"})
			}
			resp, err := h.GRPC.AppointmentClient.GetAppointmentByID(ctx, &allPb.AppointmentIDRequest{Id: idUint})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/appointments/user/") && len(path) > len("/appointments/user/") {
			userID := strings.TrimPrefix(path, "/appointments/user/")
			var userIDUint uint32
			_, err := fmt.Sscanf(userID, "%d", &userIDUint)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid user id"})
			}
			resp, err := h.GRPC.AppointmentClient.GetAppointmentsByUserID(ctx, &allPb.UserAppointmentsRequest{UserId: userIDUint})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	case http.MethodPost:
		if path == "/appointments" {
			var req allPb.CreateAppointmentRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.AppointmentClient.CreateAppointment(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusCreated, resp)
		}
	case http.MethodPut:
		if strings.HasPrefix(path, "/appointments/") && len(path) > len("/appointments/") {
			id := strings.TrimPrefix(path, "/appointments/")
			var req allPb.UpdateAppointmentRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			var idUint uint32
			_, err := fmt.Sscanf(id, "%d", &idUint)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid appointment id"})
			}
			req.Id = idUint
			resp, err := h.GRPC.AppointmentClient.UpdateAppointment(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	case http.MethodDelete:
		if strings.HasPrefix(path, "/appointments/") && len(path) > len("/appointments/") {
			id := strings.TrimPrefix(path, "/appointments/")
			var idUint uint32
			_, err := fmt.Sscanf(id, "%d", &idUint)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid appointment id"})
			}
			resp, err := h.GRPC.AppointmentClient.MarkAppointmentAsPaid(ctx, &allPb.MarkAsPaidRequest{Id: idUint})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	}
	return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
}

func (h *GatewayHandler) ProxyToAppointmentQueueService(c echo.Context) error {
	method := c.Request().Method
	path := c.Request().URL.Path

	// Copy TS-USER-* headers to downstream
	copyTSUserHeaders := func(req *http.Request, c echo.Context) {
		for _, h := range []string{"TS-USER-ID", "TS-USER-NAME", "TS-USER-ROLE", "TS-USER-EMAIL"} {
			if val := c.Request().Header.Get(h); val != "" {
				req.Header.Set(h, val)
			}
		}
	}

	switch method {
	case http.MethodGet:
		if strings.HasPrefix(path, "/queues/") && len(path) > len("/queues/") {
			id := strings.TrimPrefix(path, "/queues/")
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://appointment-queue-service:8080/queues/%s", id), nil)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			copyTSUserHeaders(req, c)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer resp.Body.Close()
			return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
		} else if strings.HasPrefix(path, "/queues-today/") && len(path) > len("/queues-today/") {
			doctorID := strings.TrimPrefix(path, "/queues-today/")
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://appointment-queue-service:8080/queues-today/%s", doctorID), nil)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			copyTSUserHeaders(req, c)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer resp.Body.Close()
			return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
		}
	case http.MethodPost:
		if path == "/queues" {
			req, err := http.NewRequest(http.MethodPost, "http://appointment-queue-service:8080/queues", c.Request().Body)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			req.Header.Set("Content-Type", c.Request().Header.Get("Content-Type"))
			copyTSUserHeaders(req, c)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer resp.Body.Close()
			return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
		} else if path == "/queues/generate" {
			req, err := http.NewRequest(http.MethodPost, "http://appointment-queue-service:8080/queues/generate", c.Request().Body)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			req.Header.Set("Content-Type", c.Request().Header.Get("Content-Type"))
			copyTSUserHeaders(req, c)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer resp.Body.Close()
			return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
		}
	case http.MethodPut:
		if strings.HasPrefix(path, "/queues/") && len(path) > len("/queues/") {
			id := strings.TrimPrefix(path, "/queues/")
			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://appointment-queue-service:8080/queues/%s", id), c.Request().Body)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			req.Header.Set("Content-Type", c.Request().Header.Get("Content-Type"))
			copyTSUserHeaders(req, c)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			defer resp.Body.Close()
			return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
		}
	}

	return c.NoContent(http.StatusMethodNotAllowed)
}

func ForwardTSUserHeadersToGRPC(ctx echo.Context) (newCtx context.Context) {
	headers := []string{"TS-USER-ID", "TS-USER-NAME", "TS-USER-ROLE", "TS-USER-EMAIL"}
	mdMap := map[string]string{}
	for _, hname := range headers {
		val := ctx.Request().Header.Get(hname)
		fmt.Println("Header:", hname, "Value:", val)
		if val != "" {
			mdMap[strings.ToLower(hname)] = val
		}
	}
	if len(mdMap) > 0 {
		md := metadata.New(mdMap)
		return metadata.NewOutgoingContext(ctx.Request().Context(), md)
	}
	return ctx.Request().Context()
}
