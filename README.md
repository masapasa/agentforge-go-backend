# AgentForge Backend

Production-ready Golang backend for **AgentForge** - Multi-agent AI coding platform.

## Inspired by Claude Code Changelog
- Agent teams & subagents
- Worktree isolation
- Plugin/MCP system
- Hooks & permissions
- Session management

## Features
- Supabase Auth (JWT)
- PostgreSQL with GORM
- Stripe Subscriptions & Webhooks
- REST API for agents, sessions, orchestration

## Quick Start
1. `cp .env.example .env` and fill values
2. `docker compose up --build`
3. API ready at http://localhost:8080

## Endpoints
- POST /api/v1/agents
- GET /api/v1/sessions
- Stripe webhooks

Built as the perfect backend companion to Claude Code's agentic features.