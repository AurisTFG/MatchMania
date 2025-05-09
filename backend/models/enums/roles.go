package enums

type Role string

const (
	AdminRole            Role = "Admin"
	ModeratorRole        Role = "Moderator"
	TrackmaniaPlayerRole Role = "Trackmania Player"
	UserRole             Role = "User"
)

func (r *Role) Scan(value any) error {
	*r = Role(value.(string))
	return nil
}

func (r *Role) Value() (any, error) {
	return string(*r), nil
}

func AllRoles() map[string]string {
	return map[string]string{
		string(AdminRole):            "Full access to all administrative features and settings",
		string(ModeratorRole):        "Can manage users, moderate content, and oversee platform activity",
		string(TrackmaniaPlayerRole): "Connected with Trackmania account, can access Trackmania-specific features",
		string(UserRole):             "Registered user with basic access to features",
	}
}

func GetPermissionsForRole(role Role) []string {
	switch role {
	case AdminRole:
		return AllPermissions()
	case ModeratorRole:
		return []string{
			string(ManageLeaguePermission),
			string(ManageTeamPermission),
			string(ViewResultPermission),
			string(ManageResultPermission),
			string(ManageQueuePermission),

			string(ModeratorPermission),
		}
	case TrackmaniaPlayerRole:
		return []string{
			string(ManageLeaguePermission),
			string(ManageTeamPermission),
			string(ViewResultPermission),
			string(ManageQueuePermission),
		}
	case UserRole:
		return []string{}
	default:
		return []string{}
	}
}
