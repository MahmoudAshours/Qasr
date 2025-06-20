package handler

import (
	"qasr/internal/app/shortener"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *shortener.ShortenerService
}

func NewHandler(service *shortener.ShortenerService) *Handler {
	return &Handler{Service: service}
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
	c.Redirect(302, url)
}
