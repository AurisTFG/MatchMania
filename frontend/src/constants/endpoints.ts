const ENDPOINTS = {
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

  PLAYERS: {
    ROOT: `players`,
    BY_ID: (id: string) => `players/${id}`,
  },

  LEAGUES: {
    ROOT: `leagues`,
    BY_ID: (id: string) => `leagues/${id}`,
  },

  RESULTS: {
    ROOT: `results`,
    BY_ID: (id: string) => `results/${id}`,
  },

  TEAMS: {
    ROOT: `teams`,
    BY_ID: (teamId: string) => `teams/${teamId}`,
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

export default ENDPOINTS;
