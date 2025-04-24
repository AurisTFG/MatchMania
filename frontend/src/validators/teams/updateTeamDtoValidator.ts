import { z } from 'zod';
import { UpdateTeamDto } from 'types/dtos/requests/teams/updateTeamDto';
import { createTeamDtoValidator } from './createTeamDtoValidator';

export const updateTeamDtoValidator =
  createTeamDtoValidator satisfies z.ZodType<UpdateTeamDto>;
