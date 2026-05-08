package config

import "os"

type Config struct {
	SupabaseJWTSecret string
	PostgresDSN       string
	StripeSecretKey   string
	StripeWebhookKey  string
}

func Load() *Config {
	return &Config{
		SupabaseJWTSecret: os.Getenv("SUPABASE_JWT_SECRET"),
		PostgresDSN:       os.Getenv("DATABASE_URL"),
		StripeSecretKey:   os.Getenv("STRIPE_SECRET_KEY"),
		StripeWebhookKey:  os.Getenv("STRIPE_WEBHOOK_SECRET"),
	}
}