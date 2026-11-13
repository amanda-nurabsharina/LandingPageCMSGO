package handlers

import (
	"strconv"

	"be-lp-go/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type NewsHandler struct {
	DB *gorm.DB
}

func (h *NewsHandler) Index(c *fiber.Ctx) error {
	var items []models.News
	if err := h.DB.Order("created_at desc").Find(&items).Error; err != nil {
		return serverError(c, err)
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   items,
	})
}

func (h *NewsHandler) Show(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var item models.News
	err := h.DB.Where("slug = ?", slug).First(&item).Error
	if err != nil {
		if id, convErr := strconv.Atoi(slug); convErr == nil && id > 0 {
			err = h.DB.First(&item, id).Error
		}
	}
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Berita tidak ditemukan",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   item,
	})
}

