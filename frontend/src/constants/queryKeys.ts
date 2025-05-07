const QUERY_KEYS = {
  AUTH: {
    ROOT: ['auth'],
    ME: ['auth', 'me'],
  },

  USERS: {
    ROOT: ['users'],
    ALL: ['users', 'all'],
    BY_ID: (id: string) => ['users', 'byId', id],
  },

  PLAYERS: {
    ROOT: ['players'],
    ALL: ['players', 'all'],
    BY_ID: (id: string) => ['players', 'byId', id],
  },

  LEAGUES: {
    ROOT: ['leagues'],
    ALL: ['leagues', 'all'],
    BY_ID: (id: string) => ['leagues', 'byId', id],
  },

  RESULTS: {
    ROOT: ['results'],
    ALL: ['results', 'all'],
    BY_ID: (id: string) => ['results', 'byId', id],
  },

  TEAMS: {
    ROOT: ['teams'],
    ALL: ['teams', 'all'],
    BY_ID: (teamId: string) => ['teams', 'byId', teamId],
  },

  TRACKMANIA: {
    ROOT: ['trackmania'],
  },

  MATCHMAKING: {
    ROOT: ['matchmaking'],
    QUEUES: {
      ROOT: ['matchmaking', 'queues'],
      ALL: ['matchmaking', 'queues', 'all'],
      BY_ID: (id: string) => ['matchmaking', 'queues', 'byId', id],
    },
    MATCHES: {
      ROOT: ['matchmaking', 'matches'],
      ALL: ['matchmaking', 'matches', 'all'],
      BY_ID: (id: string) => ['matchmaking', 'matches', 'byId', id],
    },
  },
};

export default QUERY_KEYS;
