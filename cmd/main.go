package main

import (
	handler "qasr/internal/handler/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define root route
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Qasr is live 👑")
	})
	r.POST("/shorten", handler.Shorten)
	r.GET("/r/:slug", handler.Redirect)
	// Start server
	r.Run(":8080") // Default is localhost:8080
}
