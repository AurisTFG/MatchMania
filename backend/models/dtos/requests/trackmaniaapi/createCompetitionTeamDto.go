package trackmaniaapi

import "MatchManiaAPI/models"

type CreateCompetitionTeamDto struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Seed    int          `json:"seed"`
	Members []TeamMember `json:"members"`
}

type TeamMember struct {
	Member string `json:"member"`
}

func MakeCompetitionTeams(teamA *models.Team, teamB *models.Team) []CreateCompetitionTeamDto {
	return []CreateCompetitionTeamDto{
		{
			ID:   teamA.Id.String(),
			Name: teamA.Name,
			Seed: 1,
			Members: []TeamMember{
				{Member: teamA.Players[0].TrackmaniaId.String()},
				{Member: teamA.Players[1].TrackmaniaId.String()},
			},
		},
		{
			ID:   teamB.Id.String(),
			Name: teamB.Name,
			Seed: 2,
			Members: []TeamMember{
				{Member: teamB.Players[0].TrackmaniaId.String()},
				{Member: teamB.Players[1].TrackmaniaId.String()},
			},
		},
	}
}
