package controllers

import (
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

type TrackmaniaOAuthController struct {
	trackmaniaOAuthService services.TrackmaniaOAuthService
	userService            services.UserService
}

func NewTrackmaniaOAuthController(
	trackmaniaOAuthService services.TrackmaniaOAuthService,
	userService services.UserService,
) TrackmaniaOAuthController {
	return TrackmaniaOAuthController{
		trackmaniaOAuthService: trackmaniaOAuthService,
		userService:            userService,
	}
}

// @Summary Start Trackmania OAuth flow
// @Description Start Trackmania OAuth flow
// @Tags trackmania
// @Success 200 {object} responses.TrackmaniaOAuthUrlDto
// @Failure 422 {object} responses.ErrorDto
// @Router /trackmania/oauth/url [get]
func (c *TrackmaniaOAuthController) GetAuthUrl(ctx *gin.Context) {
	state := c.trackmaniaOAuthService.GenerateRandomState()
	userId := utils.MustGetRequestingUserId(ctx)

	err := c.trackmaniaOAuthService.SaveState(state, userId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	authURL := c.trackmaniaOAuthService.GetAuthorizationUrl(state)
	if authURL == "" {
		r.UnprocessableEntity(ctx, "Failed to generate authorization URL")
		return
	}

	dto := responses.TrackmaniaOAuthUrlDto{
		Url: authURL,
	}

	r.OK(ctx, dto)
}

// @Summary Handle Trackmania OAuth callback
// @Description Handle Trackmania OAuth callback
// @Tags trackmania
// @Param code query string true "Authorization code"
// @Param state query string true "State parameter"
// @Success 302
// @Failure 422 {object} responses.ErrorDto
// @Router /trackmania/oauth/callback [get]
func (c *TrackmaniaOAuthController) HandleCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	state := ctx.Query("state")

	ok := c.trackmaniaOAuthService.VerifyCallbackResponse(state, code)
	if !ok {
		r.UnprocessableEntity(ctx, "Invalid state or code "+state)
		return
	}

	userId, err := c.trackmaniaOAuthService.GetUserIdByState(state)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	accessTokenDto, err := c.trackmaniaOAuthService.GetAccessToken(code)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	oauthUserDto, err := c.trackmaniaOAuthService.GetUserInfo(accessTokenDto.AccessToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	err = c.userService.UpdateUserWithTrackmaniaUser(userId, oauthUserDto)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	tracks, err := c.trackmaniaOAuthService.GetUserFavoriteMaps(accessTokenDto.AccessToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return

	}

	err = c.userService.UpdateUserWithTrackmaniaTracks(userId, tracks)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	profilePageUrl := c.trackmaniaOAuthService.GetProfilePageUrl()

	if profilePageUrl == "" {
		r.UnprocessableEntity(ctx, "Failed to get profile page URL")
		return
	}

	ctx.Redirect(302, profilePageUrl)
}
