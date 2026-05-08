package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	SupabaseID string   `gorm:"unique" json:"supabase_id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type Agent struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Config      string    `json:"config"` // JSON for agent config (inspired by Claude Code agents)
	CreatedAt   time.Time `json:"created_at"`
}

type Session struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `json:"user_id"`
	AgentID   uint      `json:"agent_id"`
	Status    string    `json:"status"` // running, completed, etc.
	Worktree  string    `json:"worktree"` // simulated worktree path
	CreatedAt time.Time `json:"created_at"`
}

type Payment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `json:"user_id"`
	Amount    int64     `json:"amount"`
	StripeID  string    `json:"stripe_id"`
	CreatedAt time.Time `json:"created_at"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Agent{}, &Session{}, &Payment{})
}