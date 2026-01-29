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

func RegisterRouters(mux *http.ServeMux, dbQueries *database.Queries) {
	apiCfg := apiConfig{
		DBQueries:   dbQueries,
		RedisClient: redisClient.Get(),
	}

	go simulation.UpdateUserRating(apiCfg.DBQueries, apiCfg.RedisClient)
	mux.HandleFunc("GET /api/leaderboard", apiCfg.handlerLeaderboard)
	mux.HandleFunc("GET /api/search", apiCfg.handlerSearch)
}
