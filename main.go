package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/masapasa/agentforge-go-backend/internal/config"
	"github.com/masapasa/agentforge-go-backend/internal/handlers"
	"github.com/masapasa/agentforge-go-backend/internal/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()

	app := fiber.New()

	app.Use(cors.New())

	// Auth middleware
	app.Use(middleware.AuthMiddleware(cfg.SupabaseJWTSecret))

	// Routes
	api := app.Group("/api/v1")
	handlers.SetupAgentRoutes(api, cfg)
	handlers.SetupSessionRoutes(api, cfg)
	handlers.SetupStripeRoutes(api, cfg)

	log.Fatal(app.Listen(":8080"))
}