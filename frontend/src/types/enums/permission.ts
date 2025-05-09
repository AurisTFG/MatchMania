export enum Permission {
  ManageLeague = 'LeaguesManage',
  ManageTeam = 'TeamsManage',
  ViewResult = 'ResultsView',
  ManageResult = 'ResultsManage',
  ManageUser = 'UsersManage',
  ManageRole = 'RolesManage',
  ManageQueue = 'QueueManage',

  // Special permissions
  Admin = 'Admin', // For all administrator rights
  Moderator = 'Moderator', // For all moderator rights
  LoggedIn = 'LoggedIn', // For all logged-in users, used in auth middleware and not stored in DB
  LoggedOut = 'LoggedOut', // For all logged-out users, used in auth middleware and not stored in DB
}
