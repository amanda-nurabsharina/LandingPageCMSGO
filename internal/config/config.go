package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// DBConfig holds PostgreSQL connection settings.
type DBConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SSLMode  string
}

// Config holds the whole application configuration (loaded from env / .env).
type Config struct {
	AppName     string
	AppEnv      string
	AppDebug    bool
	AppURL      string
	Port        string
	TimeZone    string
	DB          DBConfig
	CORSOrigins string
	StoragePath string
	AutoMigrate bool

	// Optional: auto-create first admin user when the users table is empty.
	AdminName     string
	AdminEmail    string
	AdminPassword string
}

// Load reads configuration from environment variables (and .env if present).
func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		AppName:     env("APP_NAME", "General CMS"),
		AppEnv:      env("APP_ENV", "local"),
		AppDebug:    envBool("APP_DEBUG", true),
		AppURL:      strings.TrimRight(env("APP_URL", "http://127.0.0.1:8000"), "/"),
		Port:        env("APP_PORT", "8000"),
		TimeZone:    env("APP_TIMEZONE", "Asia/Jakarta"),
		CORSOrigins: env("CORS_ALLOW_ORIGINS", "*"),
		StoragePath: env("STORAGE_PATH", "./storage/public"),
		AutoMigrate: envBool("DB_AUTO_MIGRATE", true),
		DB: DBConfig{
			Host:     env("DB_HOST", "127.0.0.1"),
			Port:     env("DB_PORT", "5432"),
			Name:     env("DB_DATABASE", "suppliyam_lp"),
			User:     env("DB_USERNAME", "postgres"),
			Password: env("DB_PASSWORD", ""),
			SSLMode:  env("DB_SSLMODE", "disable"),
		},
		AdminName:     env("ADMIN_NAME", "Administrator"),
		AdminEmail:    env("ADMIN_EMAIL", ""),
		AdminPassword: env("ADMIN_PASSWORD", ""),
	}
}

// IsProduction reports whether the app runs in production mode.
func (c *Config) IsProduction() bool {
	return strings.EqualFold(c.AppEnv, "production")
}

func env(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func envBool(key string, def bool) bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	switch v {
	case "1", "true", "yes", "on":
		return true
	case "0", "false", "no", "off":
		return false
	}
	return def
}

