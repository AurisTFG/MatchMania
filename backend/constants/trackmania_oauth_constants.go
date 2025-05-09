package constants

const (
	TrackmaniaOAuthBaseURL = "https://api.trackmania.com/"

	TrackmaniaOAuthURL          = TrackmaniaOAuthBaseURL + "oauth/authorize"
	TrackmaniaOAuthTokenURL     = TrackmaniaOAuthBaseURL + "api/access_token"
	TrackmaniaOAuthUserURL      = TrackmaniaOAuthBaseURL + "api/user"
	TrackmaniaOAuthFavoritesURL = TrackmaniaOAuthBaseURL + "api/user/maps/favorite"
)
