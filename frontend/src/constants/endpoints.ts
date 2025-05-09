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
    QUEUES: {
      ROOT: `matchmaking/queues`,
      BY_ID: (id: string) => `matchmaking/queues/${id}`,
      JOIN: `matchmaking/queues/join`,
      LEAVE: `matchmaking/queues/leave`,
    },

    MATCHES: {
      ROOT: `matchmaking/matches`,
      BY_ID: (id: string) => `matchmaking/matches/${id}`,
      END: `matchmaking/matches/end`,
    },
  },
};

export default ENDPOINTS;
