import { LeagueMinimalDto } from '../leagues/leagueMinimalDto';
import { TeamMinimalDto } from '../teams/teamMinimalDto';

export type MatchDto = {
  id: string;
  gameMode: string;
  league: LeagueMinimalDto;
  teams: TeamMinimalDto[];
};
