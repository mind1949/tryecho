package db

import (
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client

const (
	redisAddr     = "192.168.1.121:6379"
	redisPassword = ""
	redisDB       = 1
	SidKey        = "else.sid"
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
