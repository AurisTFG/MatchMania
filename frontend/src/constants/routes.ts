export const ROUTES = {
  HOME: '/',
  LOGIN: '/login',
  SIGNUP: '/signup',
  PROFILE: '/profile',
  SEASONS: '/seasons',
  TEAMS: (seasonId: string) => `/seasons/${seasonId}/teams`,
  RESULTS: (seasonId: string, teamId: string) =>
    `/seasons/${seasonId}/teams/${teamId}/results`,
  NOT_FOUND: '*',
};
