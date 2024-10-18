package models

type BadRequestResponse struct {
	Error string `json:"error" example:"JSON parsing error"`
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
