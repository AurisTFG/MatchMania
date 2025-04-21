import { TeamMinimalDto } from '../teams/teamMinimalDto';
import { UserMinimalDto } from '../users/userMinimalDto';

export type ResultDto = {
  id: string;
  startDate: Date;
  endDate: Date;
  score: string;
  opponentScore: string;

  team: TeamMinimalDto;
  opponentTeam: TeamMinimalDto;
  user: UserMinimalDto;
};
