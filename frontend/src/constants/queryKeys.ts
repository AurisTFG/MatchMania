export const QUERY_KEYS = {
  AUTH: {
    ROOT: ['auth'],
    ME: ['auth', 'me'],
  },

  USERS: {
    ROOT: ['users'],
    ALL: ['users', 'all'],
    BY_ID: (id: string) => ['users', 'byId', id],
  },

  SEASONS: {
    ROOT: ['seasons'],
    ALL: ['seasons', 'all'],
    BY_ID: (id: string) => ['seasons', 'byId', id],
  },

  TEAMS: {
    ROOT: ['teams'],
    ALL: (seasonId: string) => ['teams', 'all', 'season', seasonId],
    BY_ID: (seasonId: string, teamId: string) => [
      'teams',
      'byId',
      'season',
      seasonId,
      'team',
      teamId,
    ],
  },

  RESULTS: {
    ROOT: ['results'],
    ALL: (seasonId: string, teamId: string) => [
      'results',
      'all',
      'season',
      seasonId,
      'team',
      teamId,
    ],
    BY_ID: (seasonId: string, teamId: string, resultId: string) => [
      'results',
      'byId',
      'season',
      seasonId,
      'team',
      teamId,
      'result',
      resultId,
    ],
  },
};
