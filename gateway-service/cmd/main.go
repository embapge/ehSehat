package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ehSehat/gateway-service/config"
	_ "ehSehat/gateway-service/docs"
	"ehSehat/gateway-service/handler"
	ownMiddleware "ehSehat/gateway-service/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	grpcClients := config.NewGRPCClients()
	h := handler.NewGatewayHandler(grpcClients)
	e := echo.New()
	e.POST("/login", h.ProxyToAuthService, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))
	e.POST("/register", h.ProxyToAuthService)
	e.Static("/docs", "docs")

	e.POST("/xendit-payment-webhook", func(c echo.Context) error {
		paymentID := c.FormValue("payment_id")
		captureID := c.FormValue("capture_id")
		callbackToken := c.Request().Header.Get("x-callback-token")

		fmt.Println("Received xendit payment webhook:")
		fmt.Println("Payment ID:", paymentID)
		fmt.Println("Capture ID:", captureID)
		fmt.Println("X-Callback-Token:", callbackToken)

		return c.JSON(http.StatusOK, map[string]string{
			"payment_id":       paymentID,
			"capture_id":       captureID,
			"x-callback-token": callbackToken,
		})
	})

	consultationRoute := e.Group("/consultations")
	paymentRoute := e.Group("/payments")
	clinicRoute := e.Group("/clinics")
	queueRoute := e.Group("/queues")
	appointmentRoute := e.Group("/appointments")

	queueRoute.Any("", h.ProxyToAppointmentQueueService)
	queueRoute.Any("/*", h.ProxyToAppointmentQueueService)
	appointmentRoute.Use(ownMiddleware.JWTAuth, ownMiddleware.AccessRole("patient"))
	appointmentRoute.Any("", h.ProxyToAppointmentService)
	appointmentRoute.Any("/*", h.ProxyToAppointmentService)

	clinicRoute.Use(ownMiddleware.JWTAuth, ownMiddleware.AccessRole("admin", "receptionist"))
	clinicRoute.Any("", h.ProxyToClinicDataService)
	clinicRoute.Any("/*", h.ProxyToClinicDataService)

	consultationRoute.Use(ownMiddleware.JWTAuth, ownMiddleware.AccessRole("admin", "doctor"))
	consultationRoute.Any("", h.ProxyToConsultationService)
	consultationRoute.Any("/*", h.ProxyToConsultationService)

	paymentRoute.POST("/xendit-webhook", h.ProxyToPaymentWebhook)
	paymentRoute.Use(ownMiddleware.JWTAuth, ownMiddleware.AccessRole("admin", "receptionist"))
	paymentRoute.Any("", h.ProxyToPaymentService)
	paymentRoute.Any("/*", h.ProxyToPaymentService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Gateway running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
