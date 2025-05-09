package constants

const (
	TrackmaniaApiBaseURL                = "https://meet.trackmania.nadeo.club/api/"
	TrackmaniaApiCreateCompetitionURL   = TrackmaniaApiBaseURL + "competitions/web/create"
	TrackmaniaApiDeleteCompetitionURL   = TrackmaniaApiBaseURL + "competitions/%d/delete"
	TrackmaniaApiAddTeamsURL            = TrackmaniaApiBaseURL + "competitions/%d/teams"
	TrackmaniaApiGetTeamsLeaderboardURL = TrackmaniaApiBaseURL + "competitions/%d/teams/leaderboard"

	LeaderboardTypeBracket  = "BRACKET"
	LeaderboardTypeSumscore = "SUMSCORE"
	LeaderboardTypeSum      = "SUM"

	TeamLeaderboardTypeScore = "TEAM_SCORE"

	ParticipantPlayer = "PLAYER"
	ParticipantTeam   = "TEAM"

	ScriptCup        = "TrackMania/TM_Cup_Online.Script.txt"
	ScriptRounds     = "TrackMania/TM_Rounds_Online.Script.txt"
	ScriptTimeAttack = "TrackMania/TM_TimeAttack_Online.Script.txt"
	ScriptKnockout   = "TrackMania/TM_Knockout_Online.Script.txt"
	ScriptCupClassic = "TrackMania/Legacy/TM_CupClassic_Online.Script.txt"
	ScriptCupLong    = "TrackMania/Legacy/TM_CupLong_Online.Script.txt"
	ScriptCupShort   = "TrackMania/Legacy/TM_CupShort_Online.Script.txt"
	ScriptTMWT2023   = "TrackMania/TM_TMWC2023_Online.Script.txt"
	ScriptTMWT2025   = "TrackMania/TM_TMWT2025_Online.Script.txt"

	PluginClub = "server-plugins/Club/ClubPlugin.Script.txt"

	MatchGeneratorType = "spot_filler"

	SpotQualification = "round_challenge_participant"
	SpotSeed          = "competition_participant"
	SpotCompetition   = "competition_leaderboard"
	SpotTeam          = "competition_team"
	SpotMatch         = "match_participant"

	AutoStartLight = "light"
	AutoStartDelay = "delay"

	RespawnDefault       = 0
	RespawnNormal        = 1
	RespawnIgnore        = 2
	RespawnGiveUpFirstCP = 3
	RespawnAlwaysGiveUp  = 4
	RespawnNeverGiveUp   = 5
)
