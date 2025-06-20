package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	repo := mongodb.NewLinkRepository(db)
	service := shortener.NewShortenerService(repo, redisCache)
	h := handler.NewHandler(service)

	r := gin.Default()
	r.POST("/shorten", h.Shorten)
	r.GET("/r/:slug", h.Redirect)
	r.Run(":8080")
}
