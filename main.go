package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go-service-template/internal/auth"
	"go-service-template/internal/config"
	"go-service-template/internal/httptransport"
	"go-service-template/internal/integration/cart"
	"go-service-template/internal/integration/post"
	"go-service-template/internal/integration/product"
	"go-service-template/internal/logger"
	"go-service-template/internal/middleware"
	"go-service-template/internal/postgres"
	"go-service-template/internal/transaction"
	"go-service-template/internal/validator"
)

func main() {
	// config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// logger
	logger := logger.New()

	// db
	postgresDatabase, err := postgres.New(cfg.Database.Postgres)
	if err != nil {
		log.Fatalf("connect postgres: %v", err)
	}
	defer postgresDatabase.Close()

	// fiber
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		BodyLimit:    cfg.Server.BodyLimitMB * 1024 * 1024,
		ErrorHandler: httptransport.NewErrorHandler(logger),
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.CORS.AllowOrigins,
		AllowMethods: cfg.CORS.AllowMethods,
		AllowHeaders: cfg.CORS.AllowHeaders,
	}))

	// middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(logger))

	authMiddleware := auth.Middleware(
		cfg.Auth.CookieName,
		cfg.Auth.JWTSecret,
	)

	// routes
	api := app.Group("/api", authMiddleware)

	requestValidator := validator.New()

	httpClient := &http.Client{
		Timeout: cfg.HTTPClient.Timeout,
	}

	// health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// cart
	cartClient := cart.NewClient(
		cfg,
		httpClient,
	)
	cartService := cart.NewService(cartClient)
	cartHandler := cart.NewHandler(cartService, requestValidator)
	api.Get("/carts", cartHandler.GetCarts)

	// product
	productClient := product.NewClient(
		cfg,
		httpClient,
	)
	productService := product.NewService(productClient)
	productHandler := product.NewHandler(productService, requestValidator)
	api.Get("/products", productHandler.GetProducts)

	// post
	postClient := post.NewClient(
		cfg,
		httpClient,
	)
	postService := post.NewService(postClient)
	postHandler := post.NewHandler(postService)
	api.Get("/posts", postHandler.GetPosts)
	api.Get("/posts/:id", postHandler.GetPostByID)

	// transaction
	transactionRepository := transaction.NewRepository(postgresDatabase)
	transactionService := transaction.NewService(transactionRepository)
	transactionHandler := transaction.NewHandler(
		transactionService,
		requestValidator,
	)
	api.Post("/transactions/search", transactionHandler.SearchTransactions)
	api.Get("/transactions/:id", transactionHandler.GetTransactionByID)

	// server
	address := fmt.Sprintf(":%d", cfg.Server.Port)

	logger.Info(
		"server started",
		"app", cfg.App.Name,
		"address", address,
	)

	if err := app.Listen(address); err != nil {
		log.Fatal(err)
	}
}
