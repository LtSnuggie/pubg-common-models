package common

import "github.com/ltsnuggie/pubgo"

type SeasonStats struct {
	DuoTPP   pubgo.PlayerSeasonGameModeStats `json:"duo"`
	DuoFPP   pubgo.PlayerSeasonGameModeStats `json:"duo-fpp"`
	SoloTPP  pubgo.PlayerSeasonGameModeStats `json:"solo"`
	SoloFPP  pubgo.PlayerSeasonGameModeStats `json:"solo-fpp"`
	SquadTPP pubgo.PlayerSeasonGameModeStats `json:"squad"`
	SquadFPP pubgo.PlayerSeasonGameModeStats `json:"squad-fpp"`
}

type SeasonInfo struct {
	pubgo.PlayerSeasonGameModeStats
	SeasonID string `json:"seasonid"`
	PlayerID string `json:"playerid"`
	GameMode string `json:"gamemode"`
}
