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
    BY_ID: (id: string | number) => `users/${String(id)}`,
  },
  SEASONS: {
    ROOT: `seasons`,
    BY_ID: (id: string | number) => `seasons/${String(id)}`,
  },
  TEAMS: {
    ROOT: (seasonId: string | number) => `seasons/${String(seasonId)}/teams`,
    BY_ID: (seasonId: string | number, teamId: string | number) =>
      `seasons/${String(seasonId)}/teams/${String(teamId)}`,
  },
  RESULTS: {
    ROOT: (seasonId: string | number, teamId: string | number) =>
      `seasons/${String(seasonId)}/teams/${String(teamId)}/results`,
    BY_ID: (
      seasonId: string | number,
      teamId: string | number,
      resultId: string | number,
    ) =>
      `seasons/${String(seasonId)}/teams/${String(teamId)}/results/${String(
        resultId,
      )}`,
  },
};
