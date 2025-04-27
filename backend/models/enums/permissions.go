package enums

type Permission string

const (
	ManageSeasonPermission Permission = "SeasonsManage"
	ManageTeamPermission   Permission = "TeamsManage"
	ManageResultPermission Permission = "ResultsManage"
	ManageUserPermission   Permission = "UsersManage"
	ManageRolePermission   Permission = "RolesManage"
	ManageQueuePermission  Permission = "QueueManage"

	LoggedInPermission Permission = "LoggedIn" // this is a special permission for logged in users
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
		string(ManageSeasonPermission),
		string(ManageTeamPermission),
		string(ManageResultPermission),
		string(ManageUserPermission),
		string(ManageRolePermission),
		string(ManageQueuePermission),
	}
}

func AllPermissionsWithDesc() map[string]string {
	return map[string]string{
		string(ManageSeasonPermission): "Manage seasons",
		string(ManageTeamPermission):   "Manage teams",
		string(ManageResultPermission): "Manage results",
		string(ManageUserPermission):   "Manage users",
		string(ManageRolePermission):   "Manage roles",
		string(ManageQueuePermission):  "Manage queue",
	}
}
