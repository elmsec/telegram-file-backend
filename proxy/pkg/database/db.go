package database

import (
	"fmt"
	"os"

	"github.com/elmsec/telegram-file-backend/pkg/config"

	"github.com/go-redis/redis/v8"
)

var appConfig = config.InitConfig()

func NewRedisClient() *redis.Client {
	redisConf := appConfig.Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		DB:       redisConf.DB,
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	return redisDB
}
