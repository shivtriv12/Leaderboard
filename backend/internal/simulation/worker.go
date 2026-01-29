package simulation

import (
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shivtriv12/Leaderboard/internal/database"
)

func UpdateUserRating(dbOueries *database.Queries, redisClient *redis.Client) {
	ticker := time.NewTicker(10 * time.Second)

	for ; ; <-ticker.C {
		users, err := getRandomUsers(dbOueries)
		if err != nil {
			fmt.Printf("Error getting random users %v", err)
			continue
		}

		ratings := make([]int32, randomUsersSize)
		for i := range randomUsersSize {
			ratings[i] = getRandomScore()
		}

		updateRedis(users, ratings, redisClient)
		log.Println("user ratings updated in redis")

		err = BatchUpdateUserRating(dbOueries, users, ratings)
		if err != nil {
			log.Printf("Error updating users score to db %v", err)
		} else {
			log.Println("user ratings updated in db")
		}
	}
}
