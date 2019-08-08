package redis

import (
	"github.com/vasarostik/go_blog/pkg/utl/config"
	"github.com/go-redis/redis"
	"log"
)

func New(cfg *config.Redis) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	err := redisClient.Ping().Err()
	if err != nil {
		log.Println(err)
	}

	return redisClient, err
}
