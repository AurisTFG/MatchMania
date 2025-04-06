package responses

import (
	"MatchManiaAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, obj any) { // 200
	c.JSON(http.StatusOK, obj)
}

func Created(c *gin.Context, obj any) { // 201
	c.JSON(http.StatusCreated, obj)
}

func NoContent(c *gin.Context) { // 204
	c.JSON(http.StatusNoContent, nil)
}

func BadRequest(c *gin.Context, errorMessage string) { // 400
	c.JSON(http.StatusBadRequest, BadRequestResponse{Error: errorMessage})
}

func Unauthorized(c *gin.Context, errorMessage string) { // 401
	c.JSON(http.StatusUnauthorized, UnauthorizedResponse{Error: errorMessage})
	c.Abort()
}

func Forbidden(c *gin.Context, errorMessage string) { // 403
	c.JSON(http.StatusForbidden, ForbiddenResponse{Error: errorMessage})
}

func NotFound(c *gin.Context, errorMessage string) { // 404
	c.JSON(http.StatusNotFound, NotFoundResponse{Error: errorMessage})
}

func UnprocessableEntity(c *gin.Context, errorMessage string) { // 422
	c.JSON(http.StatusUnprocessableEntity, UnprocessableEntityResponse{Error: errorMessage})
}

func InternalServerError(c *gin.Context, errorMessage string) { // 500
	c.JSON(http.StatusInternalServerError, InternalServerErrorResponse{Error: errorMessage})
}

type BadRequestResponse struct {
	Error string `example:"JSON parsing error" json:"error"`
}

type UnauthorizedResponse struct {
	Error string `example:"Unauthorized" json:"error"`
}

type ForbiddenResponse struct {
	Error string `example:"Forbidden" json:"error"`
}

type NotFoundResponse struct {
	Error string `example:"Resource was not found" json:"error"`
}

type UnprocessableEntityResponse struct {
	Error string `example:"Validation error" json:"error"`
}

type InternalServerErrorResponse struct {
	Error string `example:"Internal server error" json:"error"`
}

type AuthSignUpResponse struct {
	User models.UserDto `json:"user"`
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

type UsersResponse struct {
	Users []models.UserDto `json:"users"`
}

type UserResponse struct {
	User models.UserDto `json:"user"`
}
