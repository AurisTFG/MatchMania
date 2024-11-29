package responses

import (
	"MatchManiaAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}

func Created(c *gin.Context, obj any) {
	c.JSON(http.StatusCreated, obj)
}

func Deleted(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func BadRequest(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, BadRequestResponse{Error: errorMessage})
}

func Unauthorized(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnauthorized, UnauthorizedResponse{Error: errorMessage})
	c.Abort()
}

func Forbidden(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusForbidden, ForbiddenResponse{Error: errorMessage})
}

func NotFound(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusNotFound, NotFoundResponse{Error: errorMessage})
}

func UnprocessableEntity(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnprocessableEntity, UnprocessableEntityResponse{Error: errorMessage})
}

func InternalServerError(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusInternalServerError, InternalServerErrorResponse{Error: errorMessage})
}

type BadRequestResponse struct {
	Error string `json:"error" example:"JSON parsing error"`
}

type UnauthorizedResponse struct {
	Error string `json:"error" example:"Unauthorized"`
}

type ForbiddenResponse struct {
	Error string `json:"error" example:"Forbidden"`
}

type NotFoundResponse struct {
	Error string `json:"error" example:"Resource was not found"`
}

type UnprocessableEntityResponse struct {
	Error string `json:"error" example:"Validation error"`
}

type InternalServerErrorResponse struct {
	Error string `json:"error" example:"Internal server error"`
}

type AuthSignUpResponse struct {
	User models.UserDto `json:"user"`
}

type AuthLoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type AuthRefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type SeasonResponse struct {
	Season models.SeasonDto `json:"season"`
}

type SeasonsResponse struct {
	Seasons []models.SeasonDto `json:"seasons"`
}

type TeamResponse struct {
	Team models.TeamDto `json:"team"`
}

type TeamsResponse struct {
	Teams []models.TeamDto `json:"teams"`
}

type ResultResponse struct {
	Result models.ResultDto `json:"result"`
}

type ResultsResponse struct {
	Results []models.ResultDto `json:"results"`
}
