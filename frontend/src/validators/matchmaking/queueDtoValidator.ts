import { z } from 'zod';
import { JoinQueueDto } from 'types/dtos/requests/queues/joinQueueDto';
import { isSelectOptionValid } from 'utils/selectOptionUtils';

export const queueDtoValidator = z.object({
  leagueId: z.string().refine((value) => isSelectOptionValid(value), {
    message: 'League is required',
  }),

  teamId: z.string().refine((value) => isSelectOptionValid(value), {
    message: 'Team is required',
  }),
}) satisfies z.ZodType<JoinQueueDto>;
