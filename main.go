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
	"go-service-template/internal/integration/dummyjson"
	"go-service-template/internal/integration/jsonplaceholder"
	"go-service-template/internal/logger"
	"go-service-template/internal/middleware"
	"go-service-template/internal/validator"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.New()

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

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(logger))

	authMiddleware := auth.Middleware(
		cfg.Auth.CookieName,
		cfg.Auth.JWTSecret,
	)

	api := app.Group("/api", authMiddleware)

	requestValidator := validator.New()

	httpClient := &http.Client{
		Timeout: cfg.HTTPClient.Timeout,
	}

	// dummyJSON
	dummyJSONClient := dummyjson.NewClient(
		cfg.DummyJSON.BaseURL,
		httpClient,
	)

	dummyJSONService := dummyjson.NewService(dummyJSONClient)
	dummyJSONHandler := dummyjson.NewHandler(dummyJSONService, requestValidator)
	api.Get("/carts", dummyJSONHandler.GetCarts)
	api.Get("/products", dummyJSONHandler.GetProducts)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// jsonPlaceholder
	jsonPlaceholderClient := jsonplaceholder.NewClient(
		cfg.JSONPlaceholder.BaseURL,
		httpClient,
	)

	jsonPlaceholderService := jsonplaceholder.NewService(
		jsonPlaceholderClient,
	)

	jsonPlaceholderHandler := jsonplaceholder.NewHandler(
		jsonPlaceholderService,
	)

	api.Get("/posts", jsonPlaceholderHandler.GetPosts)
	api.Get("/posts/:id", jsonPlaceholderHandler.GetPostByID)

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
