package handler

import (
	"context"
	"ehSehat/gateway-service/config"
	authpb "ehSehat/gateway-service/handler/pb"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
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
			var req authpb.LoginRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
			}
			resp, err := h.GRPC.AuthClient.Login(ctx, &req)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		} else if strings.HasPrefix(path, "/register") {
			var req authpb.RegisterRequest
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
			var req authpb.ConsultationRequest
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
			resp, err := h.GRPC.ConsultationClient.FindByIDConsultation(ctx, &authpb.ConsultationIDRequest{Id: id})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusOK, resp)
		}
	case http.MethodPut:
		if strings.HasPrefix(path, "/consultations/") && len(path) > len("/consultations/") {
			id := strings.TrimPrefix(path, "/consultations/")
			var req authpb.ConsultationRequest
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
	default:
		return c.NoContent(http.StatusMethodNotAllowed)
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
			var req authpb.PaymentRequest
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
