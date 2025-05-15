package enums

type Role string

const (
	AdminRole            Role = "Admin"
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
		string(TrackmaniaPlayerRole): "Connected with Trackmania account, can access Trackmania-specific features",
		string(UserRole):             "Registered user with basic access to features",
	}
}

func GetPermissionsForRole(role Role) []string {
	switch role {
	case AdminRole:
		return AllPermissions()
	case TrackmaniaPlayerRole:
		return []string{
			string(ViewLeaguePermission),
			string(ViewTeamPermission),
			string(ViewResultPermission),
			string(ViewPlayerPermission),
			string(ViewQueuePermission),
			string(ViewMatchPermission),

			string(ManageLeaguePermission),
			string(ManageTeamPermission),
			string(ManageQueuePermission),
			string(ManageMatchPermission),
		}
	case UserRole:
		return []string{
			string(ViewLeaguePermission),
			string(ViewTeamPermission),
			string(ViewResultPermission),
			string(ViewPlayerPermission),
			string(ViewQueuePermission),
			string(ViewMatchPermission),
		}
	default:
		return []string{}
	}
}
