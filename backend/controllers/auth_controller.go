package controllers

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	authService    services.AuthService
	sessionService services.SessionService
	env            *config.Env
}

func NewAuthController(authService services.AuthService, sessionService services.SessionService, env *config.Env) AuthController {
	return AuthController{authService: authService, sessionService: sessionService, env: env}
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

	r.Created(ctx, r.AuthSignUpResponse{User: user.ToDto()})
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
		r.UnprocessableEntity(ctx, "Username or password was incorrect.")
		return
	}

	if !user.ComparePassword(bodyDto.Password) {
		r.UnprocessableEntity(ctx, "Username or password was incorrect.")
		return
	}

	sessionUUID := uuid.New()

	accessToken, err := c.authService.CreateAccessToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	refreshToken, err := c.authService.CreateRefreshToken(sessionUUID.String(), user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	err = c.sessionService.CreateSession(sessionUUID, user.UUID, refreshToken, time.Now().AddDate(0, 0, c.env.JWTRefreshTokenExpirationDays))
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("RefreshToken", refreshToken, c.env.JWTRefreshTokenExpirationDays*24*60*60, "/", "", false, true)

	r.OK(ctx, r.AuthLoginResponse{AccessToken: accessToken})
}

func (c *AuthController) AuthLogOut(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("RefreshToken")
	if err != nil {
		r.UnprocessableEntity(ctx, "Refresh token not found")
		return
	}

	_, sessionId, err := c.authService.VerifyRefreshToken(tokenString)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if !c.sessionService.IsSessionValid(sessionId, tokenString) {
		r.UnprocessableEntity(ctx, "Session is not valid")
		return
	}

	c.sessionService.InvalidateSession(sessionId)

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("RefreshToken", "", -1, "/", "", false, true)

	r.Deleted(ctx)
}

func (c *AuthController) AuthRefreshToken(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("RefreshToken")
	if err != nil {
		r.UnprocessableEntity(ctx, "Refresh token not found")
		return
	}

	user, sessionId, err := c.authService.VerifyRefreshToken(tokenString)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if !c.sessionService.IsSessionValid(sessionId, tokenString) {
		r.UnprocessableEntity(ctx, "Session is not valid")
		return
	}

	accessToken, err := c.authService.CreateAccessToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	refreshToken, err := c.authService.CreateRefreshToken(sessionId, user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.sessionService.ExtendSession(sessionId, refreshToken, time.Now().AddDate(0, 0, c.env.JWTRefreshTokenExpirationDays))

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("RefreshToken", refreshToken, c.env.JWTRefreshTokenExpirationDays*24*60*60, "/", "", false, true)

	r.OK(ctx, r.AuthRefreshTokenResponse{AccessToken: accessToken})
}
