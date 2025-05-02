import { UserMinimalDto } from '../users/userMinimalDto';

export type LeagueDto = {
  id: string;
  name: string;
  startDate: Date;
  endDate: Date;

  user: UserMinimalDto;
};
