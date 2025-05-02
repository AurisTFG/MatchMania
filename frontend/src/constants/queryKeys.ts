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
    ROOT: ['leagues'],
    ALL: ['leagues', 'all'],
    BY_ID: (id: string) => ['leagues', 'byId', id],
  },

  TEAMS: {
    ROOT: ['teams'],
    ALL: (leagueId: string) => ['teams', 'all', 'league', leagueId],
    BY_ID: (leagueId: string, teamId: string) => [
      'teams',
      'byId',
      'league',
      leagueId,
      'team',
      teamId,
    ],
  },

  RESULTS: {
    ROOT: ['results'],
    ALL: (leagueId: string, teamId: string) => [
      'results',
      'all',
      'league',
      leagueId,
      'team',
      teamId,
    ],
    BY_ID: (leagueId: string, teamId: string, resultId: string) => [
      'results',
      'byId',
      'league',
      leagueId,
      'team',
      teamId,
      'result',
      resultId,
    ],
  },

  TRACKMANIA: {
    ROOT: ['trackmania'],
  },

  MATCHMAKING: {
    ROOT: ['matchmaking'],
    QUEUE_TEAMS_COUNT: (leagueId: string) => [
      'matchmaking',
      'queue',
      'teamsCount',
      leagueId,
    ],
    CHECK_MATCH_STATUS: (teamId: string) => [
      'matchmaking',
      'queue',
      'status',
      teamId,
    ],
  },
};
