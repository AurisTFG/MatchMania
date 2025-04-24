import { z } from 'zod';
import { JoinQueueDto } from '../../types/dtos/requests/matchmaking/joinQueueDto';

export const queueDtoValidator = z.object({
  seasonId: z.string().min(1, 'Season is required'),

  teamId: z.string().min(1, 'Team is required'),
}) satisfies z.ZodType<JoinQueueDto>;
