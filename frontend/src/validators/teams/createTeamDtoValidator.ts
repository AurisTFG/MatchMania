import { z } from 'zod';
import { CreateTeamDto } from '../../types/dtos/requests/teams/createTeamDto';

export const createTeamDtoValidator = z.object({
  name: z
    .string()
    .min(3, { message: 'Name must be at least 3 characters long.' })
    .max(100, { message: 'Name can be up to 100 characters long.' }),
}) satisfies z.ZodType<CreateTeamDto>;
