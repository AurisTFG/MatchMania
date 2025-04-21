import { UserMinimalDto } from '../users/userMinimalDto';

export type SeasonDto = {
  id: string;
  name: string;
  startDate: Date;
  endDate: Date;

  user: UserMinimalDto;
};
