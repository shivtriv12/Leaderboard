package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shivtriv12/Leaderboard/internal/database"
)

func seedDB(dbQueries database.Queries) {
	log.Print("started database seeding")

	for i := range 10000 {
		i += 1
		user := fmt.Sprintf("user_%v", i)
		err := dbQueries.CreateUser(context.Background(), database.CreateUserParams{
			Username: user,
			Ratings:  int32(i),
		})
		if err != nil {
			log.Fatalf("error seeding database %v", err)
		}
		if i%1000 == 0 {
			log.Printf("%v users seeded", i)
		}
	}

	log.Print("database seeded")
}
