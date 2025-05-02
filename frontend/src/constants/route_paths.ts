export const ROUTE_PATHS = {
  HOME: '/',
  LOGIN: '/login',
  SIGNUP: '/signup',
  PROFILE: '/profile',
  SEASONS: '/leagues',
  TEAMS: '/teams',
  RESULTS: '/results',
  MATCHMAKING: '/matchmaking',

  UNAUTHORIZED: '/unauthorized',
  FORBIDDEN: '/forbidden',
  NOT_FOUND: '*',
};

export const PARAMS = {
  SEASON_ID: 'leagueId',
  TEAM_ID: 'teamId',
};

export const getTeamsLink = (leagueId: string) => ({
  pathname: ROUTE_PATHS.TEAMS,
  search: `?${PARAMS.SEASON_ID}=${leagueId}`,
});

export const getResultsLink = (leagueId: string, teamId: string) => ({
  pathname: ROUTE_PATHS.RESULTS,
  search: `?${PARAMS.SEASON_ID}=${leagueId}&${PARAMS.TEAM_ID}=${teamId}`,
});
