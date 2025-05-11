export enum Permission {
  ViewLeague = 'LeaguesView',
  ManageLeague = 'LeaguesManage',

  ViewTeam = 'TeamsView',
  ManageTeam = 'TeamsManage',

  ViewResult = 'ResultsView',
  ManageResult = 'ResultsManage',

  ViewUser = 'UsersView',
  ManageUser = 'UsersManage',

  ViewPlayer = 'PlayersView',
  ManagePlayer = 'PlayersManage',

  ViewRole = 'RolesView',
  ManageRole = 'RolesManage',

  ViewQueue = 'QueuesView',
  ManageQueue = 'QueuesManage',

  ViewMatch = 'MatchesView',
  ManageMatch = 'MatchesManage',

  // Special permissions
  Admin = 'Admin', // For all administrator rights
  LoggedIn = 'LoggedIn', // For all logged-in users, used in auth middleware and not stored in DB
  LoggedOut = 'LoggedOut', // For all logged-out users, used in auth middleware and not stored in DB
}
