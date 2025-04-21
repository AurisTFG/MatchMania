import { UserMinimalDto } from '../users/userMinimalDto';

export type TeamDto = {
  id: string;
  name: string;
  elo: number;

  user: UserMinimalDto;
};
