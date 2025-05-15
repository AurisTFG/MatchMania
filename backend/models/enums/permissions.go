package enums

type Permission string

const (
	ViewLeaguePermission   Permission = "LeaguesView"
	ManageLeaguePermission Permission = "LeaguesManage"

	ViewTeamPermission   Permission = "TeamsView"
	ManageTeamPermission Permission = "TeamsManage"

	ViewResultPermission   Permission = "ResultsView"
	ManageResultPermission Permission = "ResultsManage"

	ViewUserPermission   Permission = "UsersView"
	ManageUserPermission Permission = "UsersManage"

	ViewPlayerPermission   Permission = "PlayersView"
	ManagePlayerPermission Permission = "PlayersManage"

	ViewRolePermission   Permission = "RolesView"
	ManageRolePermission Permission = "RolesManage"

	ViewQueuePermission   Permission = "QueuesView"
	ManageQueuePermission Permission = "QueuesManage"

	ViewMatchPermission   Permission = "MatchesView"
	ManageMatchPermission Permission = "MatchesManage"

	// Special permissions
	AdminPermission    Permission = "Admin"    // For all administrator rights
	LoggedInPermission Permission = "LoggedIn" // For all logged-in users, user in auth middleware and not stored in DB
)

func (p *Permission) Scan(value any) error {
	*p = Permission(value.(string))
	return nil
}

func (p *Permission) Value() (any, error) {
	return string(*p), nil
}

func AllPermissions() []string {
	return []string{
		string(AdminPermission),

		string(ViewLeaguePermission),
		string(ManageLeaguePermission),

		string(ViewTeamPermission),
		string(ManageTeamPermission),

		string(ViewResultPermission),
		string(ManageResultPermission),

		string(ViewUserPermission),
		string(ManageUserPermission),

		string(ViewPlayerPermission),
		string(ManagePlayerPermission),

		string(ViewRolePermission),
		string(ManageRolePermission),

		string(ViewQueuePermission),
		string(ManageQueuePermission),

		string(ViewMatchPermission),
		string(ManageMatchPermission),
	}
}

func AllPermissionsWithDesc() map[string]string {
	return map[string]string{
		string(AdminPermission): "Full access to all features",

		string(ViewLeaguePermission):   "View leagues",
		string(ManageLeaguePermission): "Manage leagues",

		string(ViewTeamPermission):   "View teams",
		string(ManageTeamPermission): "Manage teams",

		string(ViewResultPermission):   "View results",
		string(ManageResultPermission): "Manage results",

		string(ViewUserPermission):   "View users",
		string(ManageUserPermission): "Manage users",

		string(ViewPlayerPermission):   "View players",
		string(ManagePlayerPermission): "Manage players",

		string(ViewRolePermission):   "View roles",
		string(ManageRolePermission): "Manage roles",

		string(ViewQueuePermission):   "View queues",
		string(ManageQueuePermission): "Manage queues",

		string(ViewMatchPermission):   "View matches",
		string(ManageMatchPermission): "Manage matches",
	}
}
