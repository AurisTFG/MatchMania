package enums

type Permission string

const (
	ManageLeaguePermission Permission = "LeaguesManage"
	ManageTeamPermission   Permission = "TeamsManage"
	ManageResultPermission Permission = "ResultsManage"
	ManageUserPermission   Permission = "UsersManage"
	ManageRolePermission   Permission = "RolesManage"
	ManageQueuePermission  Permission = "QueueManage"

	// Special permissions
	AdminPermission     Permission = "Admin"     // For all administrator rights
	ModeratorPermission Permission = "Moderator" // For all moderator rights
	LoggedInPermission  Permission = "LoggedIn"  // For all logged-in users, user in auth middleware and not stored in DB
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
		string(ManageLeaguePermission),
		string(ManageTeamPermission),
		string(ManageResultPermission),
		string(ManageUserPermission),
		string(ManageRolePermission),
		string(ManageQueuePermission),

		string(AdminPermission),
		string(ModeratorPermission),
	}
}

func AllPermissionsWithDesc() map[string]string {
	return map[string]string{
		string(AdminPermission):     "Full access to all features",
		string(ModeratorPermission): "Moderate content and manage users",

		string(ManageLeaguePermission): "Manage leagues",
		string(ManageTeamPermission):   "Manage teams",
		string(ManageResultPermission): "Manage results",
		string(ManageUserPermission):   "Manage users",
		string(ManageRolePermission):   "Manage roles",
		string(ManageQueuePermission):  "Manage queue",
	}
}
