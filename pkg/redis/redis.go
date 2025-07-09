package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"log"
	"user-doc-collaborative-editor/pkg/logger"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Default Redis address
		DB:   0,                // Default DB
	})

	// Test the connection
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		logger.WithFields(logrus.Fields{
			"action": "InitRedis",
			"error":  err.Error(),
		}).Error("Redis connection failed")
		log.Fatalf("Redis connection failed: %v", err)
	}
	log.Println("Connected to Redis")
}
