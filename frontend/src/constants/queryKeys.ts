export const QUERY_KEYS = {
  AUTH: {
    ROOT: ["auth"],
    ME: ["auth", "me"],
  },

  USERS: {
    ROOT: ["users"],
    ALL: ["users", "all"],
    BY_ID: (id: string | number) => ["users", "byId", id],
  },

  SEASONS: {
    ROOT: ["seasons"],
    ALL: ["seasons", "all"],
    BY_ID: (id: string | number) => ["seasons", "byId", id],
  },

  TEAMS: {
    ROOT: ["teams"],
    ALL: (seasonId: string | number) => ["teams", "all", "season", seasonId],
    BY_ID: (seasonId: string | number, teamId: string | number) => [
      "teams",
      "byId",
      "season",
      seasonId,
      "team",
      teamId,
    ],
  },

  RESULTS: {
    ROOT: ["results"],
    ALL: (seasonId: string | number, teamId: string | number) => [
      "results",
      "all",
      "season",
      seasonId,
      "team",
      teamId,
    ],
    BY_ID: (
      seasonId: string | number,
      teamId: string | number,
      resultId: string | number,
    ) => [
      "results",
      "byId",
      "season",
      seasonId,
      "team",
      teamId,
      "result",
      resultId,
    ],
  },
};
