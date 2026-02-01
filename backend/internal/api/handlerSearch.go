package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/shivtriv12/Leaderboard/internal/database"
	"github.com/shivtriv12/Leaderboard/internal/types"
)

func (apiCfg *apiConfig) handlerSearch(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data       []types.Leaderboard `json:"data"`
		NextCursor string              `json:"next_cursor"`
	}

	q := r.URL.Query().Get("q")
	if q == "" {
		respondWithError(w, http.StatusBadRequest, "no query", nil)
		return
	}
	cursor := r.URL.Query().Get("cursor")
	limit := 25
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
			if limit < 0 && limit > 100 {
				limit = 25
			}
		}
	}

	users, err := apiCfg.DBQueries.GetUsersByUsername(r.Context(), database.GetUsersByUsernameParams{
		Column1:  sql.NullString{String: q, Valid: true},
		Username: cursor,
		Limit:    int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "DB error", err)
		return
	}

	pipe := apiCfg.RedisClient.Pipeline()
	for _, user := range users {
		pipe.ZCount(r.Context(), leaderboardKey, fmt.Sprintf("(%d", user.Ratings), "+inf").Result()
	}
	cmders, err := pipe.Exec(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Redis Error", err)
		return
	}

	respData := []types.Leaderboard{}
	nextCursor := ""
	for i, cmder := range cmders {
		count := cmder.(*redis.IntCmd).Val()
		rank := count + 1

		respData = append(respData, types.Leaderboard{
			GlobalRank: int(rank),
			Username:   users[i].Username,
			Rating:     int(users[i].Ratings),
		})
		nextCursor = users[i].Username
	}

	resp := response{
		Data:       respData,
		NextCursor: nextCursor,
	}

	respondWithJson(w, http.StatusCreated, resp)
}
