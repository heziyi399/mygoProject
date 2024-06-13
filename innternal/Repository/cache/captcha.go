package cache

import (
	"github.com/redis/go-redis/v9"
)

type CaptchaStorage struct {
	redis *redis.Client
}

func NewCaptchaStorage(redis *redis.Client) *CaptchaStorage {
	return &CaptchaStorage{
		redis: redis}
}
