package handlers

import (
	"strings"
	"time"

	"be-lp-go/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AnalyticsHandler struct {
	DB *gorm.DB
}

type AnalyticsPayload struct {
	EventType string  `json:"event_type"`
	PageName  *string `json:"page_name"`
	SessionID string  `json:"session_id"`
}

func (h *AnalyticsHandler) Store(c *fiber.Ctx) error {
	var payload AnalyticsPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	payload.EventType = strings.TrimSpace(payload.EventType)
	payload.SessionID = strings.TrimSpace(payload.SessionID)

	if payload.EventType == "" || payload.SessionID == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  "error",
			"message": "event_type and session_id are required",
		})
	}

	ip := c.IP()
	now := time.Now()
	event := models.AnalyticsEvent{
		EventType: payload.EventType,
		PageName:  payload.PageName,
		SessionID: payload.SessionID,
		IPAddress: &ip,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := h.DB.Create(&event).Error; err != nil {
		return serverError(c, err)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Event logged successfully",
	})
}

