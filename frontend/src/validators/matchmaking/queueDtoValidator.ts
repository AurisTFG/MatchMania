import { z } from 'zod';
import { JoinQueueDto } from 'types/dtos/requests/matchmaking/joinQueueDto';

export const queueDtoValidator = z.object({
  leagueId: z.string().min(1, 'League is required'),

  teamId: z.string().min(1, 'Team is required'),
}) satisfies z.ZodType<JoinQueueDto>;
