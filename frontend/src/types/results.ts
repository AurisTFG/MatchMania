export interface Result {
  id: number;
  matchStartDate: Date;
  matchEndDate: Date;
  score: string;
  opponentScore: string;
  seasonId: number;
  teamId: number;
  opponentTeamId: number;
  userUUID: string;
}
