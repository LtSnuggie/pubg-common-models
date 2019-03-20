package common

import (
	"errors"

	"github.com/ltsnuggie/pubgo"
)

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

func PlayerSeasonResponseToSeasonInfo(psr pubgo.PlayerSeasonResponse) (info []SeasonInfo, err error) {
	info = make([]SeasonInfo, 0)
	season := psr.Data.Relationships.Season.Data.ID
	player := psr.Data.Relationships.Player.Data.ID
	if season == "" || player == "" {
		return info, errors.New("invalid PlayerSeasonResponse")
	}
	info = append(info, SeasonInfo{psr.Data.Attributes.GameModeStats.DuoFPP, season, player, "duo-fpp"})
	info = append(info, SeasonInfo{psr.Data.Attributes.GameModeStats.DuoTPP, season, player, "duo"})
	info = append(info, SeasonInfo{psr.Data.Attributes.GameModeStats.SoloFPP, season, player, "solo-fpp"})
	info = append(info, SeasonInfo{psr.Data.Attributes.GameModeStats.SoloTPP, season, player, "solo"})
	info = append(info, SeasonInfo{psr.Data.Attributes.GameModeStats.SquadFPP, season, player, "squad-fpp"})
	info = append(info, SeasonInfo{psr.Data.Attributes.GameModeStats.SquadTPP, season, player, "squad"})
	return
}
