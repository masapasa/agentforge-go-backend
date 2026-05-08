package handlers

import "github.com/gofiber/fiber/v2"

func SetupStripeRoutes(api fiber.Router, cfg *config.Config) {
	api.Post("/stripe/webhook", stripeWebhook)
}

func stripeWebhook(c *fiber.Ctx) error {
	// Production: verify signature with cfg.StripeWebhookKey
	return c.JSON(fiber.Map{"status": "ok"})
}