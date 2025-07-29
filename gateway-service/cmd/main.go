package main

import (
	"log"
	"os"

	"ehSehat/gateway-service/config"
	_ "ehSehat/gateway-service/docs"
	"ehSehat/gateway-service/handler"
	"ehSehat/gateway-service/middleware"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Gateway API
// @version 1.0
// @description API Gateway for Shopping and Payment Services
// @host 34.101.116.9:8000
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// ===== Swagger Documentation for Shopping Service =====

// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param product body CreateProductRequest true "Create Product"
// @Success 201 {object} main.ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /products [post]
func _SwaggerCreateProduct() {}

// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {array} main.ProductResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /products [get]
func _SwaggerGetAllProducts() {}

// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} main.ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /products/{id} [get]
func _SwaggerGetProductByID() {}

// @Summary Update a product
// @Description Update a product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body main.UpdateProductRequest true "Update Product"
// @Success 200 {object} main.ProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /products/{id} [put]
func _SwaggerUpdateProduct() {}

// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /products/{id} [delete]
func _SwaggerDeleteProduct() {}

// @Summary Create a new transaction
// @Description Create a new transaction and payment
// @Tags transactions
// @Accept json
// @Produce json
// @Param transaction body CreateTransactionRequest true "Create Transaction"
// @Success 201 {object} TransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Security ApiKeyAuth
// @Router /transactions [post]
func _SwaggerCreateTransaction() {}

// @Summary Get all transactions
// @Description Get all transactions
// @Tags transactions
// @Produce json
// @Success 200 {array} TransactionResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions [get]
// @Security ApiKeyAuth
func _SwaggerGetAllTransactions() {}

// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} TransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{id} [get]
// @Security ApiKeyAuth
func _SwaggerGetTransactionByID() {}

// @Summary Update a transaction
// @Description Update a transaction with the input payload
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param transaction body CreateTransactionRequest true "Update Transaction"
// @Success 200 {object} TransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{id} [put]
// @Security ApiKeyAuth
func _SwaggerUpdateTransaction() {}

// @Summary Delete a transaction
// @Description Delete a transaction by ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 204 "No Content"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transactions/{id} [delete]
// @Security ApiKeyAuth
func _SwaggerDeleteTransaction() {}

// ===== Swagger Documentation for Payment Service =====
// @Summary Create payment
// @Description Create payment via payment-service
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body CreatePaymentRequest true "Create Payment"
// @Success 201 {object} PaymentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /payments [post]
// @Security ApiKeyAuth
func _SwaggerCreatePayment() {}

// @Summary Login via Auth Service
// @Description Login user via auth-service
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Login credentials"
// @Success 200 {object} AuthResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func _SwaggerLogin() {}

// @Summary Register via Auth Service
// @Description Register new user via auth-service
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration data"
// @Success 201 {object} AuthResponse
// @Failure 400 {object} ErrorResponse
// @Router /register [post]
func _SwaggerRegister() {}

// @Summary Get all payments
// @Description Retrieve all payments via payment-service
// @Tags payments
// @Produce json
// @Success 200 {array} PaymentResponse
// @Router /payments [get]
// @Security ApiKeyAuth
func _SwaggerGetAllPayments() {}

// @Summary Get payment by ID
// @Description Retrieve payment by ID via payment-service
// @Tags payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} PaymentResponse
// @Failure 404 {object} ErrorResponse
// @Router /payments/{id} [get]
// @Security ApiKeyAuth
func _SwaggerGetPaymentByID() {}

// @Summary Delete payment by ID
// @Description Delete payment via payment-service
// @Tags payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} ErrorResponse
// @Router /payments/{id} [delete]
// @Security ApiKeyAuth
func _SwaggerDeletePayment() {}

// ===== DTOs for Swagger =====
// Auth Service DTOs
// LoginRequest represents user login payload
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret"`
}

// RegisterRequest represents user registration payload
type RegisterRequest struct {
	Name     string `json:"name" example:"John Doe"`
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret"`
}

// AuthResponse represents authentication response from auth-service
type AuthResponse struct {
	Id    string `json:"id" example:"uuid-string"`
	Name  string `json:"name" example:"John Doe"`
	Email string `json:"email" example:"user@example.com"`
	Token string `json:"token" example:"jwt-token-string"`
}

// DTOs for Swagger
type CreateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
type UpdateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
type ProductResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type CreateTransactionRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type TransactionResponse struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	PaymentID string  `json:"payment_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type CreatePaymentRequest struct {
	Amount float64 `json:"amount"`
}
type PaymentResponse struct {
	ID        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

// PaymentListResponse represents list of payments from payment-service
// swagger:response PaymentListResponse
type PaymentListResponse struct {
	Payments []PaymentResponse `json:"payments"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	grpcClients := config.NewGRPCClients()

	// Handler HTTP to service
	h := handler.NewGatewayHandler(grpcClients)

	e := echo.New()

	e.POST("/login", h.ProxyToAuthService)
	e.POST("/register", h.ProxyToAuthService)

	// Define route groups by prefix
	consultationRoute := e.Group("/consultations")
	paymentRoute := e.Group("/payments")
	// transactionRoute := e.Group("/transactions")
	// paymentRoute := e.Group("/payments")

	// e.GET("/swagger/*", echoSwagger.WrapHandler)

	consultationRoute.Use(middleware.JWTAuth)
	consultationRoute.Any("", h.ProxyToConsultationService)
	consultationRoute.Any("/*", h.ProxyToConsultationService)

	paymentRoute.Use(middleware.JWTAuth)
	paymentRoute.Any("", h.ProxyToPaymentService)
	paymentRoute.Any("/*", h.ProxyToPaymentService)

	// productRoute.Use(middleware.JWTAuth)
	// transactionRoute.Use(middleware.JWTAuth)
	// paymentRoute.Use(middleware.JWTAuth)

	// // Proxy all /products and subpaths to shopping-service
	// productRoute.Any("", h.ProxyToShoppingService)
	// productRoute.Any("/*", h.ProxyToShoppingService)
	// // Proxy all /transactions and subpaths to shopping-service
	// transactionRoute.Any("", h.ProxyToShoppingService)
	// transactionRoute.Any("/*", h.ProxyToShoppingService)
	// // Proxy all /payments and subpaths to payment-service
	// paymentRoute.Any("", h.ProxyToPaymentService)
	// paymentRoute.Any("/*", h.ProxyToPaymentService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Gateway running on port", port)
	e.Logger.Fatal(e.Start(":" + port))
}
