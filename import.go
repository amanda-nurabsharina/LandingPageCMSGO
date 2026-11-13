package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"be-lp-go/internal/config"
	"be-lp-go/internal/database"
	"be-lp-go/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// importSQLite copies all landing page data (and admin users) from the old
// Laravel SQLite database into PostgreSQL, and copies uploaded files from the
// Laravel public disk into STORAGE_PATH.
func importSQLite(cfg *config.Config, args []string) error {
	fset := flag.NewFlagSet("import-sqlite", flag.ExitOnError)
	dbPath := fset.String("db", "../BE LP/database/database.sqlite", "path to Laravel SQLite database")
	storagePath := fset.String("storage", "../BE LP/storage/app/public", "path to Laravel public storage (empty to skip)")
	force := fset.Bool("force", false, "truncate target tables before importing")
	_ = fset.Parse(args)

	if _, err := os.Stat(*dbPath); err != nil {
		return fmt.Errorf("sqlite database not found: %s", *dbPath)
	}

	src, err := gorm.Open(sqlite.Open(*dbPath), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return fmt.Errorf("open sqlite: %w", err)
	}
	dst, err := database.Connect(cfg)
	if err != nil {
		return err
	}
	if err := database.Migrate(dst); err != nil {
		return err
	}

	steps := []struct {
		table string
		rows  func() any
	}{
		{"users", func() any { return &[]models.User{} }},
		{"site_configs", func() any { return &[]models.SiteConfig{} }},
		{"hero_sections", func() any { return &[]models.HeroSection{} }},
		{"hero_backgrounds", func() any { return &[]models.HeroBackground{} }},
		{"hero_carousels", func() any { return &[]models.HeroCarousel{} }},
		{"why_choose_us", func() any { return &[]models.WhyChooseUs{} }},
		{"cta_sections", func() any { return &[]models.CtaSection{} }},
		{"statistics", func() any { return &[]models.Statistic{} }},
		{"services", func() any { return &[]models.Service{} }},
		{"order_steps", func() any { return &[]models.OrderStep{} }},
		{"portfolios", func() any { return &[]models.Portfolio{} }},
		{"testimonials", func() any { return &[]models.Testimonial{} }},
		{"landing_sections", func() any { return &[]models.LandingSection{} }},
		{"about_items", func() any { return &[]models.AboutItem{} }},
		{"service_premiums", func() any { return &[]models.ServicePremium{} }},
		{"work_steps", func() any { return &[]models.WorkStep{} }},
		{"news", func() any { return &[]models.News{} }},
		{"activities", func() any { return &[]models.Activity{} }},
		{"branches", func() any { return &[]models.Branch{} }},
		{"clients", func() any { return &[]models.Client{} }},
		{"leads", func() any { return &[]models.Lead{} }},
		{"analytics_events", func() any { return &[]models.AnalyticsEvent{} }},
	}

	for _, s := range steps {
		if !src.Migrator().HasTable(s.table) {
			log.Printf("[skip] %s: table not found in sqlite", s.table)
			continue
		}

		var existing int64
		if err := dst.Table(s.table).Count(&existing).Error; err != nil {
			return err
		}
		if existing > 0 {
			if !*force {
				log.Printf("[skip] %s: target already has %d rows (use -force to overwrite)", s.table, existing)
				continue
			}
			if err := dst.Exec(fmt.Sprintf(`TRUNCATE TABLE "%s" RESTART IDENTITY CASCADE`, s.table)).Error; err != nil {
				return err
			}
		}

		rows := s.rows()
		if err := src.Table(s.table).Order("id").Find(rows).Error; err != nil {
			return fmt.Errorf("read %s: %w", s.table, err)
		}
		n := sliceLen(rows)
		if n == 0 {
			log.Printf("[ok]   %s: 0 rows", s.table)
			continue
		}
		if err := dst.CreateInBatches(rows, 100).Error; err != nil {
			return fmt.Errorf("write %s: %w", s.table, err)
		}
		if err := database.ResetSequence(dst, s.table); err != nil {
			return err
		}
		log.Printf("[ok]   %s: %d rows imported", s.table, n)
	}

	if *storagePath != "" {
		copied, err := copyDir(*storagePath, cfg.StoragePath)
		if err != nil {
			return fmt.Errorf("copy storage: %w", err)
		}
		log.Printf("[ok]   storage: %d files copied to %s", copied, cfg.StoragePath)
	}

	log.Println("import finished")
	return nil
}

func sliceLen(v any) int {
	switch s := v.(type) {
	case *[]models.User:
		return len(*s)
	case *[]models.SiteConfig:
		return len(*s)
	case *[]models.HeroSection:
		return len(*s)
	case *[]models.HeroBackground:
		return len(*s)
	case *[]models.HeroCarousel:
		return len(*s)
	case *[]models.WhyChooseUs:
		return len(*s)
	case *[]models.CtaSection:
		return len(*s)
	case *[]models.Statistic:
		return len(*s)
	case *[]models.Service:
		return len(*s)
	case *[]models.OrderStep:
		return len(*s)
	case *[]models.Portfolio:
		return len(*s)
	case *[]models.Testimonial:
		return len(*s)
	case *[]models.LandingSection:
		return len(*s)
	case *[]models.AboutItem:
		return len(*s)
	case *[]models.ServicePremium:
		return len(*s)
	case *[]models.WorkStep:
		return len(*s)
	case *[]models.News:
		return len(*s)
	case *[]models.Activity:
		return len(*s)
	case *[]models.Branch:
		return len(*s)
	case *[]models.Client:
		return len(*s)
	case *[]models.Lead:
		return len(*s)
	case *[]models.AnalyticsEvent:
		return len(*s)
	}
	return 0
}

func copyDir(srcDir, dstDir string) (int, error) {
	if _, err := os.Stat(srcDir); err != nil {
		log.Printf("[skip] storage: %s not found", srcDir)
		return 0, nil
	}
	count := 0
	err := filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dstDir, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if d.Name() == ".gitignore" {
			return nil
		}
		if _, err := os.Stat(target); err == nil {
			return nil // already exists
		}
		if err := copyFile(path, target); err != nil {
			return err
		}
		count++
		return nil
	})
	return count, err
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		return err
	}
	return out.Close()
}

