package simulation

import (
	"context"
	"log"
	"math/rand"

	"github.com/redis/go-redis/v9"
	"github.com/shivtriv12/Leaderboard/internal/database"
)

const randomUsersSize = 1000
const ratingRange = 10000
const leaderboardKey = "leaderboard"

func getRandomUsers(dbOueries *database.Queries) ([]string, error) {
	users, err := dbOueries.GetRandomUsers(context.Background(), randomUsersSize)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getRandomScore() int32 {
	score := rand.Intn(ratingRange) + 1 //1 to 10k
	return int32(score)
}

func BatchUpdateUserRating(dbOueries *database.Queries, users []string, ratings []int32) error {
	return dbOueries.BatchUpdateUserRating(context.Background(), database.BatchUpdateUserRatingParams{
		Column1: users,
		Column2: ratings,
	})
}

func updateRedis(usernames []string, ratings []int32, redisClient *redis.Client) {
	pipe := redisClient.Pipeline()

	for i := range randomUsersSize {
		pipe.ZAdd(context.Background(), leaderboardKey, redis.Z{
			Score:  float64(ratings[i]),
			Member: usernames[i],
		})
	}

	_, err := pipe.Exec(context.Background())
	if err != nil {
		log.Println("redis pipeline update failed:", err)
	}
}
