import { Permission } from 'types/enums/permission';

export const AllPermissions: Permission[] = [
  Permission.ManageLeague,
  Permission.ManageTeam,
  Permission.ViewResult,
  Permission.ManageResult,
  Permission.ManageUser,
  Permission.ManageRole,
  Permission.ManageQueue,

  Permission.Admin,
  Permission.Moderator,
];

export const AllPermissionsWithDesc: Record<Permission, string> = {
  [Permission.ManageLeague]: 'Manage leagues',
  [Permission.ManageTeam]: 'Manage teams',
  [Permission.ViewResult]: 'View results',
  [Permission.ManageResult]: 'Manage results',
  [Permission.ManageUser]: 'Manage users',
  [Permission.ManageRole]: 'Manage roles',
  [Permission.ManageQueue]: 'Manage queue',

  [Permission.Admin]: 'Full access to all features',
  [Permission.Moderator]: 'Moderate content and manage users',
  [Permission.LoggedIn]: 'Logged in',
  [Permission.LoggedOut]: 'Logged out',
};
