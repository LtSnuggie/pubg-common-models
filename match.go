package common

import (
	"time"

	"github.com/ltsnuggie/pubgo"
)

type MatchGlobal struct {
	GameMode    string    `json:"gameMode"`
	MapName     string    `json:"mapName"`
	CustomMatch bool      `json:"isCustom"`
	Date        time.Time `json:"createdAt"`
	Shard       string    `json:"shardId"`
	SeasonState string    `json:"seasonState"`
	MatchID     string    `json:"matchId"`
}

type MatchInfo struct {
	MatchGlobal
	Stats []pubgo.MatchStats `json:"stats"`
}
