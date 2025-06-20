package handler

import (
	"net/http"

	"qasr/internal/app/shortener"

	"github.com/gin-gonic/gin"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

func Shorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	slug := shortener.CreateShortLink(req.URL)
	c.JSON(http.StatusOK, gin.H{
		"slug":      slug,
		"short_url": "http://localhost:8080/r/" + slug,
	})
}
