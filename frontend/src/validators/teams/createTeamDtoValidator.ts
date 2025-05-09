import { z } from 'zod';
import { CreateTeamDto } from 'types/dtos/requests/teams/createTeamDto';

export const createTeamDtoValidator = z.object({
  name: z
    .string()
    .min(3, { message: 'Name must be at least 3 characters long.' })
    .max(100, { message: 'Name can be up to 100 characters long.' }),

  logoUrl: z
    .string()
    .url({ message: 'Logo URL must be a valid URL.' })
    .max(255, { message: 'Logo URL can be up to 200 characters long.' })
    .nullable(),

  leagueIds: z
    .array(z.string())
    .nonempty({ message: 'At least one league is required.' })
    .max(10, { message: 'You can select up to 10 leagues.' }),

  playerIds: z
    .array(z.string())
    .nonempty({ message: 'At least one player is required.' })
    .max(10, { message: 'You can select up to 10 players.' }),
}) satisfies z.ZodType<CreateTeamDto>;
