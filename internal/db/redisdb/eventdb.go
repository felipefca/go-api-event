package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"go-api-event/configs"
	"go-api-event/internal/appctx"
	"go-api-event/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type EventRedisDB interface {
	GetRecentEvents(ctx context.Context) ([]models.Event, error)
}

type eventRedisDB struct {
	redisConn *redis.Client
}

func NewEventRedisDB(redisConn *redis.Client) EventRedisDB {
	return &eventRedisDB{
		redisConn: redisConn,
	}
}

func (db eventRedisDB) GetRecentEvents(ctx context.Context) ([]models.Event, error) {
	logger := appctx.FromContext(ctx)
	cfg := configs.GetConfig().RedisDB

	prefix := cfg.KeyPrefix + "*"
	keys, err := db.redisConn.Keys(ctx, prefix).Result()
	if err != nil {
		return nil, err
	}

	var events []models.Event
	for _, key := range keys {
		ttl, err := db.redisConn.TTL(ctx, key).Result()
		if err != nil {
			logger.Error(fmt.Sprintf("Error getting TTL for key %s: %v", key, err))
			continue
		}

		if ttl > 0 && ttl <= 5*time.Minute {
			val, err := db.redisConn.Get(ctx, key).Result()
			if err != nil {
				logger.Error(fmt.Sprintf("Error getting value for key %s: %v", key, err))
				continue
			}

			var event models.Event
			err = json.Unmarshal([]byte(val), &event)
			if err != nil {
				logger.Error(fmt.Sprintf("Error unmarshaling value for key %s: %v", key, err))
				continue
			}

			events = append(events, event)
		}
	}

	return events, nil
}
