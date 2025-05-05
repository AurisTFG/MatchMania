import { LeagueMinimalDto } from '../leagues/leagueMinimalDto';
import { PlayerMinimalDto } from '../players/playerMinimalDto';
import { UserMinimalDto } from '../users/userMinimalDto';

export type TeamDto = {
  id: string;
  name: string;
  logoUrl: string;
  elo: number;

  user: UserMinimalDto;
  players: PlayerMinimalDto[];
  leagues: LeagueMinimalDto[];
};
