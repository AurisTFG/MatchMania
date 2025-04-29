export const ROUTE_PATHS = {
  HOME: '/',
  LOGIN: '/login',
  SIGNUP: '/signup',
  PROFILE: '/profile',
  SEASONS: '/seasons',
  TEAMS: '/teams',
  RESULTS: '/results',
  MATCHMAKING: '/matchmaking',

  UNAUTHORIZED: '/unauthorized',
  FORBIDDEN: '/forbidden',
  NOT_FOUND: '*',
};

export const PARAMS = {
  SEASON_ID: 'seasonId',
  TEAM_ID: 'teamId',
};

export const getTeamsLink = (seasonId: string) => ({
  pathname: ROUTE_PATHS.TEAMS,
  search: `?${PARAMS.SEASON_ID}=${seasonId}`,
});

export const getResultsLink = (seasonId: string, teamId: string) => ({
  pathname: ROUTE_PATHS.RESULTS,
  search: `?${PARAMS.SEASON_ID}=${seasonId}&${PARAMS.TEAM_ID}=${teamId}`,
});
