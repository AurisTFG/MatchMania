import { LeagueMinimalDto } from '../leagues/leagueMinimalDto';
import { TeamMinimalDto } from '../teams/teamMinimalDto';

export type QueueDto = {
  id: string;
  gameMode: string;
  league: LeagueMinimalDto;
  teams: TeamMinimalDto[];
};
