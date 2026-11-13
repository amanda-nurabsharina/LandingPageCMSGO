package database

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"be-lp-go/internal/config"
	"be-lp-go/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect opens a PostgreSQL connection using GORM.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := strings.Join([]string{
		kv("host", cfg.DB.Host),
		kv("port", cfg.DB.Port),
		kv("user", cfg.DB.User),
		kv("password", cfg.DB.Password),
		kv("dbname", cfg.DB.Name),
		kv("sslmode", cfg.DB.SSLMode),
		kv("TimeZone", cfg.TimeZone),
	}, " ")

	level := logger.Warn
	if cfg.AppDebug {
		level = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             500 * time.Millisecond,
			LogLevel:                  level,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
	if err != nil {
		return nil, fmt.Errorf("connect postgres: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}

// Migrate creates / updates all tables.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(models.All()...)
}

// ResetSequence moves the serial sequence of a table past MAX(id). Needed after
// inserting rows with explicit IDs (e.g. singleton id=1 or a data import).
// Table names come from internal constants only.
func ResetSequence(db *gorm.DB, table string) error {
	q := fmt.Sprintf(
		`SELECT setval(pg_get_serial_sequence('%[1]s', 'id'), COALESCE((SELECT MAX(id) FROM "%[1]s"), 0) + 1, false)`,
		table,
	)
	return db.Exec(q).Error
}

// kv builds a libpq keyword/value pair.
func kv(key, value string) string {
	if strings.ContainsAny(value, " '\"") {
		v := strings.ReplaceAll(value, `\`, `\\`)
		v = strings.ReplaceAll(v, `'`, `\'`)
		return fmt.Sprintf("%s='%s'", key, v)
	}
	return fmt.Sprintf("%s=%s", key, value)
}

