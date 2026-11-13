package handlers

import (
	"log"

	"be-lp-go/internal/models"
	"be-lp-go/internal/seed"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// LandingPage serves the consolidated landing page payload for the General branch.
type LandingPage struct {
	DB *gorm.DB
}

func (h *LandingPage) Index(c *fiber.Ctx) error {
	site, hero, heroBg, heroCarousel, why, cta, err := seed.Singletons(h.DB)
	if err != nil {
		return serverError(c, err)
	}
	if err := seed.ListDefaults(h.DB); err != nil {
		return serverError(c, err)
	}

	statistics := make([]models.Statistic, 0)
	services := make([]models.Service, 0)
	orderSteps := make([]models.OrderStep, 0)
	portfolios := make([]models.Portfolio, 0)
	testimonials := make([]models.Testimonial, 0)
	sections := make([]models.LandingSection, 0)
	aboutItems := make([]models.AboutItem, 0)
	servicePremiums := make([]models.ServicePremium, 0)
	workSteps := make([]models.WorkStep, 0)
	latestNews := make([]models.News, 0)
	latestActivities := make([]models.Activity, 0)
	branches := make([]models.Branch, 0)
	clients := make([]models.Client, 0)

	queries := []struct {
		dest  any
		order string
	}{
		{&statistics, "sort_order, id"},
		{&services, "sort_order, id"},
		{&orderSteps, "step_number, id"},
		{&portfolios, "sort_order, id"},
		{&testimonials, "sort_order, id"},
		{&sections, "sort_order, id"},
		{&aboutItems, "sort_order, id"},
		{&servicePremiums, "sort_order, id"},
		{&workSteps, "sort_order, id"},
	}
	for _, q := range queries {
		if err := h.DB.Order(q.order).Find(q.dest).Error; err != nil {
			return serverError(c, err)
		}
	}

	_ = h.DB.Order("created_at desc").Limit(5).Find(&latestNews).Error
	_ = h.DB.Order("created_at desc").Limit(5).Find(&latestActivities).Error
	_ = h.DB.Where("is_active = ?", true).Order("sort_order, id").Find(&branches).Error
	_ = h.DB.Where("is_active = ?", true).Order("sort_order, id").Find(&clients).Error

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"site_config":      site,
			"hero_section":     hero,
			"hero_background":  heroBg,
			"hero_carousel":    heroCarousel,
			"why_choose_us":    why,
			"cta_section":      cta,
			"statistics":       statistics,
			"services":         services,
			"order_steps":      orderSteps,
			"portfolios":       portfolios,
			"testimonials":     testimonials,
			"sections":         sections,
			"about_items":      aboutItems,
			"service_premiums": servicePremiums,
			"work_steps":       workSteps,
			"news":             latestNews,
			"activities":       latestActivities,
			"branches":         branches,
			"clients":          clients,
		},
	})
}

func serverError(c *fiber.Ctx, err error) error {
	log.Printf("server error: %v", err)
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "error",
		"message": "Internal server error",
	})
}
