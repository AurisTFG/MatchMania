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
    ROOT: `seasons`,
    BY_ID: (id: string) => `seasons/${id}`,
  },
  TEAMS: {
    ROOT: (seasonId: string) => `seasons/${seasonId}/teams`,
    BY_ID: (seasonId: string, teamId: string) =>
      `seasons/${seasonId}/teams/${teamId}`,
  },
  RESULTS: {
    ROOT: (seasonId: string, teamId: string) =>
      `seasons/${seasonId}/teams/${teamId}/results`,
    BY_ID: (seasonId: string, teamId: string, resultId: string) =>
      `seasons/${seasonId}/teams/${teamId}/results/${resultId}`,
  },
  MATCHMAKING: {
    JOIN_QUEUE: `matchmaking/queue/join`,
    LEAVE_QUEUE: `matchmaking/queue/leave`,
    GET_QUEUED_TEAMS_COUNT: (seasonId: string) =>
      `matchmaking/queue/teams-count/${seasonId}`,
    GET_QUEUE_STATUS: (teamId: string) => `matchmaking/queue/status/${teamId}`,
  },
};
