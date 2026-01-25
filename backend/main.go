package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shivtriv12/Leaderboard/internal/database"
	"github.com/shivtriv12/Leaderboard/internal/redisClient"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env")
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("error opening postgres")
	}
	dbQueries := database.New(db)

	// seedDB(dbQueries) ---------use this on first run to seed db----------
	redisClient.Init()
	defer redisClient.Close()
	err = buildRedis(dbQueries)
	if err != nil {
		log.Fatal(err)
	}
}
