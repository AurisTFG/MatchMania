export type CreateTeamDto = {
  name: string;
  logoUrl: string | null;

  leagueIds: string[];
  playerIds: string[];
};
