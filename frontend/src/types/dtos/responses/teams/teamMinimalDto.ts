import { PlayerMinimalDto } from '../players/playerMinimalDto';

export type TeamMinimalDto = {
  id: string;
  name: string;
  logoUrl: string | null;

  players: PlayerMinimalDto[];
};
