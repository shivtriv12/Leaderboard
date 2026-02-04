package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shivtriv12/Leaderboard/internal/api"
	"github.com/shivtriv12/Leaderboard/internal/database"
	"github.com/shivtriv12/Leaderboard/internal/redisClient"
)

func main() {
	_ = godotenv.Load()
	// if err != nil {
	// 	log.Fatal("error loading env")
	// }

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

	mux := http.NewServeMux()
	api.RegisterRouters(mux, dbQueries)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Printf("server starting on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
