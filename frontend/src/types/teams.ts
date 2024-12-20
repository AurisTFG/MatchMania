export interface Team {
  id: number;
  name: string;
  elo: number;
  seasonId: number;
  userUUID: string;
}

export interface TeamsResponse {
  teams?: Team[];
}

export interface TeamResponse {
  team?: Team;
}
