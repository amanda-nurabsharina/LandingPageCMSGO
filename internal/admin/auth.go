package admin

import (
	"errors"
	"strings"

	"be-lp-go/internal/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// dummyHash is compared against when the email is unknown, so response time
// does not reveal which emails exist.
var dummyHash, _ = bcrypt.GenerateFromPassword([]byte("not-a-real-password"), 12)

// requireAuth only lets authenticated admin users through.
func (a *Admin) requireAuth(c *fiber.Ctx) error {
	sess, err := a.store.Get(c)
	if err != nil {
		return c.Redirect("/admin/login")
	}
	uid := toUint(sess.Get(sessUserID))
	if uid == 0 {
		return c.Redirect("/admin/login")
	}
	var user models.User
	if err := a.db.First(&user, uid).Error; err != nil {
		_ = sess.Destroy()
		return c.Redirect("/admin/login")
	}
	c.Locals("user", &user)
	return c.Next()
}

func (a *Admin) loginPage(c *fiber.Ctx) error {
	if sess, err := a.store.Get(c); err == nil && toUint(sess.Get(sessUserID)) != 0 {
		return c.Redirect("/admin")
	}
	return a.render(c, "login", nil)
}

func (a *Admin) login(c *fiber.Ctx) error {
	email := strings.ToLower(strings.TrimSpace(c.FormValue("email")))
	password := c.FormValue("password")

	fail := func() error {
		return a.render(c.Status(fiber.StatusUnprocessableEntity), "login", fiber.Map{
			"Error": "Email atau password salah.",
			"Email": email,
		})
	}
	if email == "" || password == "" {
		return fail()
	}

	var user models.User
	if err := a.db.Where("LOWER(email) = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// constant-ish time: still run a bcrypt compare
			_ = bcrypt.CompareHashAndPassword(dummyHash, []byte(password))
		}
		return fail()
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return fail()
	}

	sess, err := a.store.Get(c)
	if err != nil {
		return err
	}
	if err := sess.Regenerate(); err != nil {
		return err
	}
	sess.Set(sessUserID, user.ID)
	if err := sess.Save(); err != nil {
		return err
	}
	return c.Redirect("/admin")
}

func (a *Admin) logout(c *fiber.Ctx) error {
	if sess, err := a.store.Get(c); err == nil {
		_ = sess.Destroy()
	}
	return c.Redirect("/admin/login")
}

// HashPassword hashes a password with bcrypt (same algorithm as Laravel, so
// imported Laravel password hashes keep working).
func HashPassword(pw string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
	return string(b), err
}

func toUint(v any) uint {
	switch n := v.(type) {
	case uint:
		return n
	case int:
		if n > 0 {
			return uint(n)
		}
	case int64:
		if n > 0 {
			return uint(n)
		}
	case uint64:
		return uint(n)
	}
	return 0
}
