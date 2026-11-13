package admin

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"strings"
	"time"

	"be-lp-go/internal/config"
	"be-lp-go/internal/models"
	"be-lp-go/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

//go:embed templates/*.html
var templatesFS embed.FS

const (
	sessUserID = "user_id"
	sessFlash  = "flash"
	csrfKey    = "csrf"
	perPage    = 15
)

// Admin is the server-rendered admin panel (replacement of the Filament panel at /admin).
type Admin struct {
	db      *gorm.DB
	cfg     *config.Config
	storage *storage.Storage
	store   *session.Store
	pages   map[string]*template.Template
}

// Register mounts the admin panel routes under /admin.
func Register(app *fiber.App, db *gorm.DB, cfg *config.Config, st *storage.Storage) error {
	a := &Admin{
		db:      db,
		cfg:     cfg,
		storage: st,
		store: session.New(session.Config{
			Expiration:     2 * time.Hour, // SESSION_LIFETIME=120
			KeyLookup:      "cookie:lp_admin_session",
			CookieHTTPOnly: true,
			CookieSecure:   cfg.IsProduction(),
			CookieSameSite: "Lax",
		}),
	}
	if err := a.loadTemplates(); err != nil {
		return err
	}

	g := app.Group("/admin", csrf.New(csrf.Config{
		KeyLookup:      "form:_csrf",
		CookieName:     "lp_admin_csrf",
		CookieSameSite: "Lax",
		CookieHTTPOnly: true,
		CookieSecure:   cfg.IsProduction(),
		Expiration:     2 * time.Hour,
		ContextKey:     csrfKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusForbidden).SendString("Sesi formulir kedaluwarsa (CSRF). Silakan kembali, muat ulang halaman, lalu coba lagi.")
		},
	}))

	loginLimiter := limiter.New(limiter.Config{
		Max:        10,
		Expiration: time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return a.render(c.Status(fiber.StatusTooManyRequests), "login", fiber.Map{
				"Error": "Terlalu banyak percobaan login. Coba lagi dalam 1 menit.",
			})
		},
	})

	g.Get("/login", a.loginPage)
	g.Post("/login", loginLimiter, a.login)
	g.Post("/logout", a.requireAuth, a.logout)

	g.Get("/", a.requireAuth, a.dashboard)
	g.Get("/:slug", a.requireAuth, a.list)
	g.Get("/:slug/create", a.requireAuth, a.createPage)
	g.Post("/:slug/bulk-delete", a.requireAuth, a.bulkDelete)
	g.Post("/:slug", a.requireAuth, a.store_)
	g.Get("/:slug/:id/edit", a.requireAuth, a.editPage)
	g.Post("/:slug/:id/delete", a.requireAuth, a.destroy)
	g.Post("/:slug/:id", a.requireAuth, a.update)

	return nil
}

// ---------------------------------------------------------------------------
// Templates
// ---------------------------------------------------------------------------

func (a *Admin) loadTemplates() error {
	funcs := template.FuncMap{
		"storageURL": func(p string) string { return "/storage/" + p },
		"initial": func(s string) string {
			for _, r := range strings.TrimSpace(s) {
				return strings.ToUpper(string(r))
			}
			return "?"
		},
	}
	a.pages = map[string]*template.Template{}

	login, err := template.New("login.html").Funcs(funcs).ParseFS(templatesFS, "templates/login.html")
	if err != nil {
		return err
	}
	a.pages["login"] = login

	for _, p := range []string{"dashboard", "list", "form"} {
		t, err := template.New("layout.html").Funcs(funcs).ParseFS(templatesFS, "templates/layout.html", "templates/"+p+".html")
		if err != nil {
			return err
		}
		a.pages[p] = t
	}
	return nil
}

func (a *Admin) render(c *fiber.Ctx, page string, data fiber.Map) error {
	if data == nil {
		data = fiber.Map{}
	}
	for _, k := range []string{"Title", "Error", "Email", "Flash", "FormError"} {
		if _, ok := data[k]; !ok && (k != "Flash" || page == "login") {
			data[k] = ""
		}
	}
	data["AppName"] = a.cfg.AppName
	data["CSRF"], _ = c.Locals(csrfKey).(string)
	data["Nav"] = Resources
	data["GroupedNav"] = GroupedNav()
	if u, ok := c.Locals("user").(*models.User); ok {
		data["User"] = u
	}
	if page != "login" {
		if _, ok := data["Flash"]; !ok {
			data["Flash"] = a.popFlash(c)
		}
	}

	t := a.pages[page]
	var buf bytes.Buffer
	name := "layout.html"
	if page == "login" {
		name = "login.html"
	}
	if err := t.ExecuteTemplate(&buf, name, data); err != nil {
		log.Printf("admin template %s: %v", page, err)
		return fiber.NewError(fiber.StatusInternalServerError, "template error")
	}
	c.Type("html", "utf-8")
	return c.Send(buf.Bytes())
}

// ---------------------------------------------------------------------------
// Session helpers
// ---------------------------------------------------------------------------

func (a *Admin) setFlash(c *fiber.Ctx, msg string) {
	sess, err := a.store.Get(c)
	if err != nil {
		return
	}
	sess.Set(sessFlash, msg)
	_ = sess.Save()
}

func (a *Admin) popFlash(c *fiber.Ctx) string {
	sess, err := a.store.Get(c)
	if err != nil {
		return ""
	}
	v, _ := sess.Get(sessFlash).(string)
	if v != "" {
		sess.Delete(sessFlash)
		_ = sess.Save()
	}
	return v
}
