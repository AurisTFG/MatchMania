export enum Permission {
  ManageSeason = 'SeasonsManage',
  ManageTeam = 'TeamsManage',
  ManageResult = 'ResultsManage',
  ManageUser = 'UsersManage',
  ManageRole = 'RolesManage',
  ManageQueue = 'QueueManage',

  LoggedIn = 'LoggedIn',
}

export const AllPermissions: Permission[] = [
  Permission.ManageSeason,
  Permission.ManageTeam,
  Permission.ManageResult,
  Permission.ManageUser,
  Permission.ManageRole,
  Permission.ManageQueue,
];

export const AllPermissionsWithDesc: Record<Permission, string> = {
  [Permission.ManageSeason]: 'Manage seasons',
  [Permission.ManageTeam]: 'Manage teams',
  [Permission.ManageResult]: 'Manage results',
  [Permission.ManageUser]: 'Manage users',
  [Permission.ManageRole]: 'Manage roles',
  [Permission.ManageQueue]: 'Manage queue',
  [Permission.LoggedIn]: 'Logged in',
};
