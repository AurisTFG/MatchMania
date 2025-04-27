package enums

type Role string

const (
	AdminRole          Role = "Admin"
	ModeratorRole      Role = "Moderator"
	TrackmaniaUserRole Role = "Trackmania User"
	UserRole           Role = "User"
	GuestRole          Role = "Guest"
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
		string(AdminRole):          "Full access to all administrative features and settings",
		string(ModeratorRole):      "Can manage users, moderate content, and oversee platform activity",
		string(TrackmaniaUserRole): "Connected with Trackmania account, can access Trackmania-specific features",
		string(UserRole):           "Registered user with basic access to features",
		string(GuestRole):          "Limited access, can only view public information",
	}
}

func GetPermissionsForRole(role Role) []string {
	switch role {
	case AdminRole:
		return AllPermissions()
	case ModeratorRole:
		return []string{
			string(ManageSeasonPermission),
			string(ManageTeamPermission),
			string(ManageResultPermission),
			string(ManageQueuePermission),
		}
	case TrackmaniaUserRole:
		return []string{
			string(ManageQueuePermission),
		}
	case UserRole:
		return []string{}
	case GuestRole:
		return []string{}
	default:
		return []string{}
	}
}
