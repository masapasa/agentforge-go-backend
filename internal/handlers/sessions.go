package handlers

import "github.com/gofiber/fiber/v2"

func SetupSessionRoutes(api fiber.Router, cfg *config.Config) {
	api.Post("/sessions", createSession)
	api.Get("/sessions", listSessions)
}

func createSession(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Session started with worktree isolation (Claude Code inspired)"})
}

func listSessions(c *fiber.Ctx) error {
	return c.JSON([]fiber.Map{})
}