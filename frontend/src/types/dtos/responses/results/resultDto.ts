import { LeagueMinimalDto } from '../leagues/leagueMinimalDto';
import { TeamMinimalDto } from '../teams/teamMinimalDto';
import { UserMinimalDto } from '../users/userMinimalDto';

export type ResultDto = {
  id: string;
  startDate: Date;
  endDate: Date;
  score: number;
  opponentScore: number;
  eloDiff: number;
  opponentEloDiff: number;
  newElo: number;
  newOpponentElo: number;

  league: LeagueMinimalDto;
  team: TeamMinimalDto;
  opponentTeam: TeamMinimalDto;
  user: UserMinimalDto;
};
