package types

type Leaderboard struct {
	GlobalRank int    `json:"global_rank"`
	Username   string `json:"username"`
	Rating     int    `json:"rating"`
}
