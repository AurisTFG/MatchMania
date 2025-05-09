package ubisoftapi

type UbisoftSessionDto struct {
	PlatformType                  string  `json:"platformType"`
	Ticket                        string  `json:"ticket"`
	TwoFactorAuthenticationTicket *string `json:"twoFactorAuthenticationTicket"`
	ProfileID                     string  `json:"profileId"`
	UserID                        string  `json:"userId"`
	NameOnPlatform                string  `json:"nameOnPlatform"`
	Environment                   string  `json:"environment"`
	Expiration                    string  `json:"expiration"`
	SpaceID                       string  `json:"spaceId"`
	ClientIP                      string  `json:"clientIp"`
	ClientIPCountry               string  `json:"clientIpCountry"`
	ServerTime                    string  `json:"serverTime"`
	SessionID                     string  `json:"sessionId"`
	SessionKey                    string  `json:"sessionKey"`
	RememberMeTicket              *string `json:"rememberMeTicket"`
}
