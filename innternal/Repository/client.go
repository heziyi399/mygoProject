package Repository

import (
	"github.com/redis/go-redis/v9"
	"testProject/config"
)

type ClientStorage struct {
	redis  *redis.Client
	config *config.Config
}
