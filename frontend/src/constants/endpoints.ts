export const ENDPOINTS = {
  AUTH: {
    LOGIN: `auth/login`,
    LOGOUT: `auth/logout`,
    ME: `auth/me`,
    REFRESH: `auth/refresh`,
    SIGNUP: `auth/signup`,
  },
  USERS: {
    ROOT: `users`,
    BY_ID: (id: string) => `users/${id}`,
  },
  SEASONS: {
    ROOT: `leagues`,
    BY_ID: (id: string) => `leagues/${id}`,
  },
  TEAMS: {
    ROOT: (leagueId: string) => `leagues/${leagueId}/teams`,
    BY_ID: (leagueId: string, teamId: string) =>
      `leagues/${leagueId}/teams/${teamId}`,
  },
  RESULTS: {
    ROOT: (leagueId: string, teamId: string) =>
      `leagues/${leagueId}/teams/${teamId}/results`,
    BY_ID: (leagueId: string, teamId: string, resultId: string) =>
      `leagues/${leagueId}/teams/${teamId}/results/${resultId}`,
  },
  TRACKMANIA: {
    OAUTH: {
      URL: `trackmania/oauth/url`,
    },
  },
  MATCHMAKING: {
    JOIN_QUEUE: `matchmaking/queue/join`,
    LEAVE_QUEUE: `matchmaking/queue/leave`,
    GET_QUEUED_TEAMS_COUNT: (leagueId: string) =>
      `matchmaking/queue/teams-count/${leagueId}`,
    GET_QUEUE_STATUS: (teamId: string) => `matchmaking/queue/status/${teamId}`,
  },
};
