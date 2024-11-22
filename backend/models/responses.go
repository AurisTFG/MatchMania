package models

type BadRequestResponse struct {
	Error string `json:"error" example:"JSON parsing error"`
}

type UnauthorizedResponse struct {
	Error string `json:"error" example:"Unauthorized"`
}

type NotFoundResponse struct {
	Error string `json:"error" example:"Resource was not found"`
}

type UnprocessableEntityResponse struct {
	Error string `json:"error" example:"Validation error"`
}

type BadGatewayResponse struct {
	Error string `json:"error" example:"Unable to communicate with the database"`
}

type UserSignUpResponse struct {
	User UserDto `json:"user"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserRefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type SeasonResponse struct {
	Season SeasonDto `json:"season"`
}

type SeasonsResponse struct {
	Seasons []SeasonDto `json:"seasons"`
}

type TeamResponse struct {
	Team TeamDto `json:"team"`
}

type TeamsResponse struct {
	Teams []TeamDto `json:"teams"`
}

type ResultResponse struct {
	Result ResultDto `json:"result"`
}

type ResultsResponse struct {
	Results []ResultDto `json:"results"`
}
