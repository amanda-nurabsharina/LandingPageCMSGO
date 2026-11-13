package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"be-lp-go/internal/admin"
	"be-lp-go/internal/config"
	"be-lp-go/internal/database"
	"be-lp-go/internal/handlers"
	"be-lp-go/internal/models"
	"be-lp-go/internal/seed"
	"be-lp-go/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/gorm"
)

const usage = `BE LP GO - General Landing Page Backend (Go + Fiber + GORM + PostgreSQL)

Usage:
  go run . [command] [flags]

Commands:
  serve           Start the HTTP server (default)
  migrate         Run database migrations (create/update tables)
  seed            Insert default landing page content (only into empty tables)
  create-admin    Create / update an admin user
                    -name "Admin" -email admin@example.com -password secret
  import-sqlite   Import data from the old Laravel SQLite database + uploaded files
                    -db "../BE LP/database/database.sqlite"
                    -storage "../BE LP/storage/app/public"
                    -force   (truncate target tables before import)
`

func main() {
	cfg := config.Load()

	cmd := "serve"
	args := []string{}
	if len(os.Args) > 1 {
		cmd = os.Args[1]
		args = os.Args[2:]
	}

	var err error
	switch cmd {
	case "serve":
		err = serve(cfg)
	case "migrate":
		err = withDB(cfg, func(db *gorm.DB) error {
			if err := database.Migrate(db); err != nil {
				return err
			}
			log.Println("migration completed")
			return nil
		})
	case "seed":
		err = withDB(cfg, func(db *gorm.DB) error {
			if _, _, _, _, _, _, err := seed.Singletons(db); err != nil {
				return err
			}
			if err := seed.ListDefaults(db); err != nil {
				return err
			}
			log.Println("seed completed")
			return nil
		})
	case "create-admin":
		err = createAdmin(cfg, args)
	case "import-sqlite":
		err = importSQLite(cfg, args)
	case "help", "-h", "--help":
		fmt.Print(usage)
	default:
		fmt.Print(usage)
		os.Exit(2)
	}
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func withDB(cfg *config.Config, fn func(db *gorm.DB) error) error {
	db, err := database.Connect(cfg)
	if err != nil {
		return err
	}
	return fn(db)
}

func serve(cfg *config.Config) error {
	db, err := database.Connect(cfg)
	if err != nil {
		return err
	}
	if cfg.AutoMigrate {
		if err := database.Migrate(db); err != nil {
			return fmt.Errorf("migrate: %w", err)
		}
	}
	if err := ensureAdminFromEnv(cfg, db); err != nil {
		return err
	}

	st, err := storage.New(cfg.StoragePath)
	if err != nil {
		return fmt.Errorf("storage: %w", err)
	}

	app := fiber.New(fiber.Config{
		AppName:      cfg.AppName,
		BodyLimit:    25 * 1024 * 1024,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorHandler: errorHandler,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	corsMw := cors.New(cors.Config{
		AllowOrigins: cfg.CORSOrigins,
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":   cfg.AppName,
			"status": "ok",
			"endpoints": fiber.Map{
				"landing_page": "/api/landing-page",
				"news":         "/api/news",
				"activities":   "/api/activities",
				"analytics":    "/api/analytics",
				"leads":        "/api/leads",
				"admin":        "/admin",
				"storage":      "/storage/{path}",
			},
		})
	})
	app.Get("/up", func(c *fiber.Ctx) error { return c.SendString("OK") })

	// Public REST API (consumed by FE LP)
	api := app.Group("/api", corsMw)
	lp := &handlers.LandingPage{DB: db}
	newsH := &handlers.NewsHandler{DB: db}
	actH := &handlers.ActivitiesHandler{DB: db}
	analyticsH := &handlers.AnalyticsHandler{DB: db}
	leadsH := &handlers.LeadsHandler{DB: db}

	api.Get("/landing-page", lp.Index)
	api.Get("/news", newsH.Index)
	api.Get("/news/:slug", newsH.Show)
	api.Get("/activities", actH.Index)
	api.Get("/activities/:slug", actH.Show)
	api.Post("/analytics", analyticsH.Store)
	api.Post("/leads", leadsH.Store)

	// Public uploaded files (Laravel "public" disk / storage:link)
	app.Use("/storage", corsMw)
	app.Static("/storage", st.Root, fiber.Static{Browse: false, MaxAge: 86400})

	// Admin panel (replacement of the Filament panel)
	if err := admin.Register(app, db, cfg, st); err != nil {
		return fmt.Errorf("admin: %w", err)
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit
		log.Println("shutting down...")
		_ = app.ShutdownWithTimeout(10 * time.Second)
	}()

	log.Printf("server listening on :%s (APP_URL=%s)", cfg.Port, cfg.AppURL)
	return app.Listen(":" + cfg.Port)
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	} else {
		log.Printf("error: %v", err)
	}
	if strings.HasPrefix(c.Path(), "/api") {
		return c.Status(code).JSON(fiber.Map{"status": "error", "message": msg})
	}
	return c.Status(code).SendString(fmt.Sprintf("%d - %s", code, msg))
}

// ensureAdminFromEnv creates the first admin from ADMIN_EMAIL/ADMIN_PASSWORD
// when the users table is empty.
func ensureAdminFromEnv(cfg *config.Config, db *gorm.DB) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	if cfg.AdminEmail == "" || cfg.AdminPassword == "" {
		log.Println("WARNING: belum ada user admin. Jalankan: go run . create-admin -email you@example.com -password secret")
		return nil
	}
	if err := upsertAdmin(db, cfg.AdminName, cfg.AdminEmail, cfg.AdminPassword); err != nil {
		return err
	}
	log.Printf("admin user %s created from environment", cfg.AdminEmail)
	return nil
}

func createAdmin(cfg *config.Config, args []string) error {
	fs := flag.NewFlagSet("create-admin", flag.ExitOnError)
	name := fs.String("name", "Administrator", "admin name")
	email := fs.String("email", "", "admin email (required)")
	password := fs.String("password", "", "admin password (required, min 8 chars)")
	_ = fs.Parse(args)

	if *email == "" || *password == "" {
		fs.Usage()
		return fmt.Errorf("-email and -password are required")
	}
	if len(*password) < 8 {
		return fmt.Errorf("password minimal 8 karakter")
	}

	return withDB(cfg, func(db *gorm.DB) error {
		if err := database.Migrate(db); err != nil {
			return err
		}
		if err := upsertAdmin(db, *name, *email, *password); err != nil {
			return err
		}
		log.Printf("admin user %s saved", *email)
		return nil
	})
}

func upsertAdmin(db *gorm.DB, name, email, password string) error {
	hash, err := admin.HashPassword(password)
	if err != nil {
		return err
	}
	email = strings.ToLower(strings.TrimSpace(email))
	now := time.Now()

	var user models.User
	err = db.Where("LOWER(email) = ?", email).First(&user).Error
	if err == nil {
		return db.Model(&user).Updates(map[string]any{"name": name, "password": hash, "updated_at": now}).Error
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	user = models.User{Name: name, Email: email, Password: hash, EmailVerifiedAt: &now, CreatedAt: &now, UpdatedAt: &now}
	return db.Create(&user).Error
}

