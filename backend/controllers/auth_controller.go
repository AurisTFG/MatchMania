package controllers

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"MatchManiaAPI/responses"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
	env         *config.Env
}

func NewAuthController(authService services.AuthService, env *config.Env) AuthController {
	return AuthController{authService: authService, env: env}
}

func (c *AuthController) AuthSignUp(ctx *gin.Context) {
	var bodyDto models.SignUpDto

	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user, err := c.authService.CreateUser(&bodyDto)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.Created(ctx, responses.AuthSignUpResponse{User: user.ToDto()})
}

func (c *AuthController) AuthLogIn(ctx *gin.Context) {
	var bodyDto models.LoginDto

	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user, err := c.authService.GetUserByEmail(bodyDto.Email)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if !user.ComparePassword(bodyDto.Password) {
		r.UnprocessableEntity(ctx, "Invalid email or password")
		return
	}

	accessToken, err := c.authService.CreateAccessToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	refreshToken, err := c.authService.CreateRefreshToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("AccessToken", accessToken, c.env.JWTAccessTokenExpirationDays*24*60*60, "/", "", false, true)
	ctx.SetCookie("RefreshToken", refreshToken, c.env.JWTRefreshTokenExpirationDays*24*60*60, "/", "", false, true)

	r.OK(ctx, responses.AuthLoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}

func (c *AuthController) AuthLogOut(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("AccessToken", "", -1, "/", "", false, true)
	ctx.SetCookie("RefreshToken", "", -1, "/", "", false, true)

	r.Deleted(ctx)
}

func (c *AuthController) AuthRefreshToken(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("RefreshToken")
	if err != nil {
		r.Unauthorized(ctx, "Refresh token not found")
		return
	}

	user, err := c.authService.VerifyRefreshToken(tokenString)
	if err != nil {
		r.Unauthorized(ctx, err.Error())
		return
	}

	accessToken, err := c.authService.CreateAccessToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	refreshToken, err := c.authService.CreateRefreshToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("AccessToken", accessToken, c.env.JWTAccessTokenExpirationDays*24*60*60, "/", "", false, true)
	ctx.SetCookie("RefreshToken", refreshToken, c.env.JWTRefreshTokenExpirationDays*24*60*60, "/", "", false, true)

	r.OK(ctx, responses.AuthLoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}
