package db

import (
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client

const (
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = 1
)

func init() {
	RedisClient = defaultRedisClient()
}

func defaultRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	pong, err := client.Ping().Result()
	log.Println(pong, err)
	return client
}
