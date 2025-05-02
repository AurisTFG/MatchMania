export enum Permission {
  ManageLeague = 'LeaguesManage',
  ManageTeam = 'TeamsManage',
  ManageResult = 'ResultsManage',
  ManageUser = 'UsersManage',
  ManageRole = 'RolesManage',
  ManageQueue = 'QueueManage',

  LoggedIn = 'LoggedIn',
}

export const AllPermissions: Permission[] = [
  Permission.ManageLeague,
  Permission.ManageTeam,
  Permission.ManageResult,
  Permission.ManageUser,
  Permission.ManageRole,
  Permission.ManageQueue,
];

export const AllPermissionsWithDesc: Record<Permission, string> = {
  [Permission.ManageLeague]: 'Manage leagues',
  [Permission.ManageTeam]: 'Manage teams',
  [Permission.ManageResult]: 'Manage results',
  [Permission.ManageUser]: 'Manage users',
  [Permission.ManageRole]: 'Manage roles',
  [Permission.ManageQueue]: 'Manage queue',
  [Permission.LoggedIn]: 'Logged in',
};
