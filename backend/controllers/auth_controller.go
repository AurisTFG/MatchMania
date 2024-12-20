package controllers

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"net/http"

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

// @Summary Sign up
// @Description Sign up
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.SignUpDto true "Sign up details"
// @Success 201 {object} responses.AuthSignUpResponse
// @Failure 400 {object} responses.BadRequestResponse
// @Failure 422 {object} responses.UnprocessableEntityResponse
// @Router /auth/signup [post]
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

// @Summary Log in
// @Description Log in
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.LoginDto true "Log in details"
// @Success 200 {object} responses.AuthLoginResponse
// @Failure 400 {object} responses.BadRequestResponse
// @Failure 422 {object} responses.UnprocessableEntityResponse
// @Router /auth/login [post]
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

	err = c.sessionService.CreateSession(sessionUUID, user.UUID, refreshToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.SetCookie(ctx, refreshToken)

	r.OK(ctx, r.AuthLoginResponse{AccessToken: accessToken})
}

// @Summary Log out
// @Description Log out
// @Tags auth
// @Success 204
// @Failure 422 {object} responses.UnprocessableEntityResponse
// @Router /auth/logout [post]
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

	err = c.sessionService.InvalidateSession(sessionId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.SetCookie(ctx, "refreshToken")

	r.Deleted(ctx)
}

// @Summary Refresh token
// @Description Refresh token
// @Tags auth
// @Success 200 {object} responses.AuthRefreshTokenResponse
// @Failure 422 {object} responses.UnprocessableEntityResponse
// @Router /auth/refresh-token [post]
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

	err = c.sessionService.ExtendSession(sessionId, refreshToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.SetCookie(ctx, refreshToken)

	r.OK(ctx, r.AuthRefreshTokenResponse{AccessToken: accessToken})
}

func (c *AuthController) SetCookie(ctx *gin.Context, refreshToken string) {
	ctx.SetSameSite(http.SameSiteLaxMode)

	name := "RefreshToken"
	value := refreshToken
	maxAge := -1
	if refreshToken != "" {
		maxAge = int(c.env.JWTRefreshTokenDuration.Seconds())
	}
	path := "/"
	domain := "localhost"
	secure := !c.env.IsDev
	httpOnly := true

	ctx.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}
