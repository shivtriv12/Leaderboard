package api

import (
	"net/http"

	"github.com/redis/go-redis/v9"
	"github.com/shivtriv12/Leaderboard/internal/database"
	"github.com/shivtriv12/Leaderboard/internal/redisClient"
	"github.com/shivtriv12/Leaderboard/internal/simulation"
)

type apiConfig struct {
	DBQueries   *database.Queries
	RedisClient *redis.Client
}

const leaderboardKey = "leaderboard"

func RegisterRouters(mux *http.ServeMux, dbQueries *database.Queries) {
	apiCfg := apiConfig{
		DBQueries:   dbQueries,
		RedisClient: redisClient.Get(),
	}

	go simulation.UpdateUserRating(apiCfg.DBQueries, apiCfg.RedisClient)

	mux.HandleFunc("GET /api/leaderboard", corsMiddleware(apiCfg.handlerLeaderboard))
	mux.HandleFunc("GET /api/search", corsMiddleware(apiCfg.handlerSearch))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // change when deploy
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		next(w, r)
	}
}
