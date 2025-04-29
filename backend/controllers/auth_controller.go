package controllers

import (
	"MatchManiaAPI/constants"
	"MatchManiaAPI/models/dtos/requests"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpresponses"
	"MatchManiaAPI/validators"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(
	authService services.AuthService,
	userService services.UserService,
) AuthController {
	return AuthController{
		authService: authService,
		userService: userService,
	}
}

// @Summary Get current user
// @Description Get current user
// @Tags auth
// @Success 200 {object} responses.UserDto
// @Failure 422 {object} responses.ErrorDto
// @Router /auth/me [get]
func (c *AuthController) GetMe(ctx *gin.Context) {
	userId := utils.MustGetRequestingUserId(ctx)

	user, err := c.userService.GetUserById(userId)
	if err != nil {
		r.NotFound(ctx, err.Error())
		return
	}

	dto := utils.MustCopy[responses.UserDto](user)

	permissions, err := c.userService.GetDistinctPermissionsByUserId(userId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	dto.Permissions = permissions

	r.OK(ctx, dto)
}

// @Summary Sign up
// @Description Sign up
// @Tags auth
// @Accept json
// @Produce json
// @Param signUpDto body requests.SignupDto true "Sign up details"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /auth/signup [post]
func (c *AuthController) SignUp(ctx *gin.Context) {
	var bodyDto requests.SignupDto

	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if err := c.userService.CreateUser(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Log in
// @Description Log in
// @Tags auth
// @Accept json
// @Produce json
// @Param loginDto body requests.LoginDto true "Log in details"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /auth/login [post]
func (c *AuthController) LogIn(ctx *gin.Context) {
	var bodyDto requests.LoginDto

	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user, err := c.userService.GetUserByEmail(bodyDto.Email)
	if err != nil {
		r.UnprocessableEntity(ctx, "Username or password was incorrect.")
		return
	}

	if !user.ComparePassword(bodyDto.Password) {
		r.UnprocessableEntity(ctx, "Username or password was incorrect.")
		return
	}

	sessionId := uuid.New()

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

	err = c.authService.CreateSession(sessionId, user.Id, refreshToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.authService.SetCookies(ctx, accessToken, refreshToken)

	r.NoContent(ctx)
}

// @Summary Log out
// @Description Log out
// @Tags auth
// @Success 204
// @Failure 422 {object} responses.ErrorDto
// @Router /auth/logout [post]
func (c *AuthController) LogOut(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(constants.RefreshTokenName)
	if err != nil {
		r.UnprocessableEntity(ctx, "Already logged out")
		c.authService.DeleteCookies(ctx)
		return
	}

	_, sessionId, err := c.authService.VerifyRefreshToken(tokenString)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		c.authService.DeleteCookies(ctx)
		return
	}

	if !c.authService.IsSessionValid(sessionId, tokenString) {
		r.UnprocessableEntity(ctx, "Session is not valid")
		c.authService.DeleteCookies(ctx)
		return
	}

	err = c.authService.InvalidateSession(sessionId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		c.authService.DeleteCookies(ctx)
		return
	}

	c.authService.DeleteCookies(ctx)

	r.NoContent(ctx)
}

// @Summary Refresh token
// @Description Refresh token
// @Tags auth
// @Success 204
// @Failure 422 {object} responses.ErrorDto
// @Router /auth/refresh [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(constants.RefreshTokenName)
	if err != nil {
		r.UnprocessableEntity(ctx, "Refresh token not found")
		return
	}

	user, sessionId, err := c.authService.VerifyRefreshToken(tokenString)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if !c.authService.IsSessionValid(sessionId, tokenString) {
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

	err = c.authService.ExtendSession(sessionId, refreshToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.authService.SetCookies(ctx, accessToken, refreshToken)

	r.NoContent(ctx)
}
