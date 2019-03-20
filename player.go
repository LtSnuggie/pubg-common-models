package common

type PlayerAccount struct {
	Platform string `json:"platform"`
	Name     string `json:"name"`
	ID       string `json:"id"`
}

type PlayerInfo struct {
	Matches []MatchInfoLite `json:"matches"`
	Seasons []string        `json:"seasons"`
}
