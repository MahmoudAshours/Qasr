package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"qasr/internal/app/analytics"
	"qasr/internal/app/shortener"
	handler "qasr/internal/handler/http"
	"qasr/internal/repo/mongodb"
	redisCache "qasr/internal/repo/redis"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("qasr")
	redisCache := redisCache.NewRedisCache("localhost:6379")
	linkRepo := mongodb.NewLinkRepository(db)
	clickRepo := mongodb.NewClickRepo(db)

	shortenerService := shortener.NewShortenerService(linkRepo, redisCache)
	analyticsService := analytics.NewAnalyticsService(clickRepo)

	h := handler.NewHandler(shortenerService, analyticsService)

	r := gin.Default()
	r.POST("/shorten", h.Shorten)
	r.GET("/r/:slug", h.Redirect)
	r.GET("/dashboard/:slug", h.Dashboard)
	r.POST("/ai/slug", h.SuggestSlug)
	r.GET("/ai/describe/:slug", h.DescribeSlug)

	r.Run(":8080")
}
