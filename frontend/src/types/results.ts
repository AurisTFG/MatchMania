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

export interface ResultsResponse {
  results?: Result[];
}

export interface ResultResponse {
  result?: Result;
}
