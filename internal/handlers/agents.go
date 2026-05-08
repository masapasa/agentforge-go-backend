package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masapasa/agentforge-go-backend/internal/config"
	"github.com/masapasa/agentforge-go-backend/internal/models"
	"gorm.io/gorm"
)

func SetupAgentRoutes(api fiber.Router, cfg *config.Config) {
	// In real impl, inject DB
	api.Post("/agents", createAgent)
	api.Get("/agents", listAgents)
}

// Placeholder handlers - production would use DB service
func createAgent(c *fiber.Ctx) error {
	// Parse user from locals, create agent
	return c.JSON(fiber.Map{"message": "Agent created - inspired by Claude Code subagents and plugins"})
}

func listAgents(c *fiber.Ctx) error {
	return c.JSON([]fiber.Map{{"id": 1, "name": "CodeReviewer", "description": "Multi-agent code review like /ultrareview"}})
}