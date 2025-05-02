import { z } from 'zod';
import { CreateLeagueDto } from 'types/dtos/requests/leagues/createLeagueDto';

export const leagueDtoValidator = z
  .object({
    name: z
      .string()
      .min(3, { message: 'Name must be at least 3 characters long' })
      .max(100, { message: 'Name must be at most 100 characters long' }),

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
  })
  .refine((data) => data.endDate > data.startDate, {
    message: 'End date must be after the start date',
    path: ['endDate'],
  })
  .refine(
    (data) => {
      const diffInDays =
        (data.endDate.getTime() - data.startDate.getTime()) /
        (1000 * 60 * 60 * 24);
      return diffInDays >= 7;
    },
    {
      message: 'End date must be at least 7 days after the start date',
      path: ['endDate'],
    },
  ) satisfies z.ZodType<CreateLeagueDto>;
