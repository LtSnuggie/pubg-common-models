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

type MatchInfoLite struct {
	MatchGlobal
	DBNOs        int     `json:"DBNOs"`
	Assists      int     `json:"assists"`
	DamageDealt  float32 `json:"damageDealt"`
	Kills        int     `json:"kills"`
	LongestKill  float32 `json:"longestKill"`
	Name         string  `json:"name"`
	PlayerID     string  `json:"playerId"`
	Revives      int     `json:"revives"`
	TeamKills    int     `json:"teamKills"`
	TimeSurvived float32 `json:"timeSurvived"`
	WinPlace     int     `json:"winPlace"`
}

type MatchHeader struct {
	ID    string `json:"id"`
	Shard string `json:"shard"`
}

func MatchResponseToMatchInfo(mr pubgo.MatchResponse) (mi MatchInfo) {
	return MatchInfo{
		Stats:       mr.GetStats(),
		MatchGlobal: MatchResponseToMatchGlobal(mr),
	}
}

func MatchResponseToMatchGlobal(mr pubgo.MatchResponse) (glbl MatchGlobal) {
	return MatchGlobal{
		GameMode:    mr.GetGameMode(),
		MapName:     mr.GetMapName(),
		CustomMatch: mr.GetCustomMatch(),
		Date:        mr.GetCreatedAt(),
		Shard:       mr.GetShardID(),
		SeasonState: mr.GetSeasonState(),
		MatchID:     mr.GetMatchID(),
	}
}
