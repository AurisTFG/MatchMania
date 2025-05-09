package trackmaniaapi

type TeamsResultsDto struct {
	Teams []TeamResultDto `json:"teams"`
}

type TeamResultDto struct {
	TeamId string `json:"team"`
	Rank   int    `json:"rank"`
	Score  int    `json:"score"`
	Zone   string `json:"zone"`
}
