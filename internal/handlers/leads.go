package handlers

import (
	"strings"
	"time"

	"be-lp-go/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LeadsHandler struct {
	DB *gorm.DB
}

type LeadPayload struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Message   string `json:"message"`
	SessionID string `json:"session_id"`
}

func (h *LeadsHandler) Store(c *fiber.Ctx) error {
	var payload LeadPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	payload.Name = strings.TrimSpace(payload.Name)
	payload.Phone = strings.TrimSpace(payload.Phone)
	payload.Message = strings.TrimSpace(payload.Message)

	if payload.Name == "" || payload.Phone == "" || payload.Message == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "Semua field (nama, telepon, pesan) wajib diisi.",
		})
	}

	now := time.Now()
	lead := models.Lead{
		Name:      payload.Name,
		Phone:     payload.Phone,
		Message:   payload.Message,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := h.DB.Create(&lead).Error; err != nil {
		return serverError(c, err)
	}

	// Also log analytics event (lead_submitted) like Laravel
	ip := c.IP()
	pageName := "Contact Form"
	sid := payload.SessionID
	if sid == "" {
		sid = "unknown"
	}
	event := models.AnalyticsEvent{
		EventType: "lead_submitted",
		PageName:  &pageName,
		SessionID: sid,
		IPAddress: &ip,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	_ = h.DB.Create(&event).Error

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Pesan Anda berhasil terkirim. Admin kami akan segera menghubungi Anda!",
		"data":    lead,
	})
}

