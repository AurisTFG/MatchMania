package requests

type CreateTeamDto struct {
	Name string `example:"BIG Clan" json:"name" validate:"required,min=3,max=100"`
}
