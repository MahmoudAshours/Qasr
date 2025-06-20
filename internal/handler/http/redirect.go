package handler

import (
	"net/http"

	"qasr/internal/app/shortener"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context) {
	slug := c.Param("slug")
	originalURL, err := shortener.GetOriginalURL(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}
	c.Redirect(http.StatusFound, originalURL)
}
