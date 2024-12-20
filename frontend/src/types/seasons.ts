export interface Season {
  id: number;
  name: string;
  startDate: Date;
  endDate: Date;
  userUUID: string;
}

export interface SeasonsResponse {
  seasons?: Season[];
}

export interface SeasonResponse {
  season?: Season;
}
