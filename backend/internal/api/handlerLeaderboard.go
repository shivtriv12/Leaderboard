package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/shivtriv12/Leaderboard/internal/types"
)

func (apiCfg *apiConfig) handlerLeaderboard(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Data       []types.Leaderboard `json:"data"`
		NextCursor string              `json:"next_cursor"`
	}

	cursorUser := r.URL.Query().Get("cursor")
	limit := 50
	if limitStr := r.URL.Query().Get("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
			if limit < 0 && limit > 100 {
				limit = 50
			}
		}
	}

	var startIndex int64 = 0
	if cursorUser != "" {
		rank, err := apiCfg.RedisClient.ZRevRank(r.Context(), leaderboardKey, cursorUser).Result()
		if err == nil {
			startIndex = rank + 1
		} else {
			respondWithError(w, http.StatusInternalServerError, "redis error", err)
			return
		}
	}

	users, err := apiCfg.RedisClient.ZRevRangeWithScores(r.Context(), leaderboardKey, startIndex, startIndex+int64(limit-1)).Result()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "redis error", err)
		return
	}
	if len(users) == 0 {
		respondWithJson(w, http.StatusOK, []types.Leaderboard{})
		return
	}

	firstUserScore := users[0].Score
	count, _ := apiCfg.RedisClient.ZCount(r.Context(), leaderboardKey, fmt.Sprintf("(%f", firstUserScore), "+inf").Result()
	currentRank := count + 1

	respData := []types.Leaderboard{}
	nextCursor := ""

	for i, user := range users {
		if i > 0 {
			if user.Score == users[i-1].Score {
				//same score so same rank
			} else {
				currentRank = startIndex + int64(i) + 1
			}
		}

		respData = append(respData, types.Leaderboard{
			GlobalRank: int(currentRank),
			Username:   user.Member.(string),
			Rating:     int(user.Score),
		})

		nextCursor = user.Member.(string)
	}

	resp := response{
		Data:       respData,
		NextCursor: nextCursor,
	}

	respondWithJson(w, http.StatusOK, resp)
}
