package requests

type UpdateTeamDto struct {
	Name string `example:"BIG Clan" json:"name" validate:"required,min=3,max=100"`
}
