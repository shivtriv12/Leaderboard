package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/shivtriv12/Leaderboard/internal/database"
	"github.com/shivtriv12/Leaderboard/internal/redisClient"
)

func buildRedis(dbQueries *database.Queries) error {
	log.Println("starting redis build")

	const leaderboardKey = "leaderboard"
	const batchSize = 1000
	redisClient := redisClient.Get()

	users, err := dbQueries.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	pipe := redisClient.Pipeline()
	count := 0
	for _, user := range users {
		pipe.ZAdd(context.Background(), leaderboardKey, redis.Z{
			Score:  float64(user.Ratings),
			Member: user.Username,
		})
		count++
		if count%batchSize == 0 {
			if _, err := pipe.Exec(context.Background()); err != nil {
				return err
			}
			pipe = redisClient.Pipeline()
		}
	}
	if _, err := pipe.Exec(context.Background()); err != nil {
		return err
	}

	log.Println("Redis built")
	return nil
}
