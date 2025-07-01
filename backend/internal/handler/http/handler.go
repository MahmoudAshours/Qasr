package handler

import (
	"qasr/backend/internal/app/analytics"
	"qasr/backend/internal/app/shortener"
	model "qasr/backend/internal/domain"
	"qasr/backend/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service   *shortener.ShortenerService
	Analytics *analytics.AnalyticsService
}

func NewHandler(shortenerService *shortener.ShortenerService, analyticsService *analytics.AnalyticsService) *Handler {
	return &Handler{Service: shortenerService, Analytics: analyticsService}
}

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func (h *Handler) Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid"})
		return
	}

	slug := h.Service.CreateShortLink(req.URL)
	c.JSON(200, gin.H{
		"slug":      slug,
		"short_url": "http://localhost:8080/r/" + slug,
	})
}

func (h *Handler) Redirect(c *gin.Context) {
	slug := c.Param("slug")
	url, err := h.Service.GetOriginalURL(slug)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	go func() {
		req := c.Request

		// Basic info
		ip := c.ClientIP()
		ua := req.UserAgent()
		ref := req.Referer()
		lang := req.Header.Get("Accept-Language")

		// Use your own parser util or library here
		browser, deviceType, isBot := utils.ParseUserAgent(ua)

		// Optional: use geo resolver (e.g., MaxMind, IPinfo, or custom)
		// country, city, timezone := LookupGeo(ip)

		click := &model.Click{
			Slug:       slug,
			Timestamp:  time.Now(),
			IP:         ip,
			Country:    "EG",
			City:       "Cairo",
			Timezone:   "UTC+2",
			UserAgent:  ua,
			Browser:    browser,
			DeviceType: deviceType,
			IsBot:      isBot,
			Referrer:   ref,
			Language:   lang,
		}

		_ = h.Analytics.Repo.SaveClick(click)
	}()

	c.Redirect(302, url)
}

func (h *Handler) Dashboard(c *gin.Context) {
	slug := c.Param("slug")
	stats, err := h.Analytics.GetAnalytics(slug)

	if err != nil {
		c.JSON(500, stats)
		return
	}
	c.JSON(200, stats)
}
