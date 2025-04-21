import { z } from 'zod';
import { CreateResultDto } from '../../types/dtos/requests/results/createResultDto';

export const resultDtoValidator = z
  .object({
    startDate: z
      .date()
      .refine(
        (date) => date >= new Date(Date.now() - 30 * 24 * 60 * 60 * 1000),
        {
          message: 'Start date must not be earlier than 30 days ago',
        },
      )
      .refine(
        (date) => date <= new Date(Date.now() + 3650 * 24 * 60 * 60 * 1000),
        { message: 'Start date must not be more than 10 years in the future' },
      ),

    endDate: z
      .date()
      .refine(
        (date) => date <= new Date(Date.now() + 3650 * 24 * 60 * 60 * 1000),
        { message: 'End date must not be more than 10 years in the future' },
      ),

    score: z.string().refine((value) => /^\d+$/.test(value), {
      message: 'Score must be a non-negative integer',
    }),

    opponentScore: z.string().refine((value) => /^\d+$/.test(value), {
      message: 'Opponent score must be a non-negative integer',
    }),

    opponentTeamId: z.string().uuid({
      message: 'Opponent team ID must be a valid UUID',
    }),
  })
  .refine((data) => data.endDate >= data.startDate, {
    message: 'End date must be after or equal to the start date',
    path: ['endDate'],
  })
  .refine(
    (data) => {
      const diffInDays =
        (data.endDate.getTime() - data.startDate.getTime()) /
        (1000 * 60 * 60 * 24);
      return diffInDays <= 1;
    },
    {
      message: 'End date must be at most 1 day after the start date',
      path: ['endDate'],
    },
  ) satisfies z.ZodType<CreateResultDto>;
