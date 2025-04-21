import { z } from 'zod';
import { CreateSeasonDto } from '../../types/dtos/requests/seasons/createSeasonDto';

export const createSeasonDtoValidator = z
  .object({
    name: z
      .string()
      .min(3, { message: 'Name must be at least 3 characters long' })
      .max(100, { message: 'Name must not exceed 100 characters' }),
    startDate: z
      .date()
      .refine(
        (date) => {
          const now = new Date();
          const minDate = new Date(now.setDate(now.getDate() - 30));
          return date >= minDate;
        },
        { message: 'Start date must be within the last 30 days' },
      )
      .refine(
        (date) => {
          const now = new Date();
          const maxDate = new Date(now.setDate(now.getDate() + 3650));
          return date <= maxDate;
        },
        { message: 'Start date must be within the next 10 years' },
      ),
    endDate: z.date().refine(
      (date) => {
        const now = new Date();
        const maxDate = new Date(now.setDate(now.getDate() + 3650));
        return date <= maxDate;
      },
      { message: 'End date must be within the next 10 years' },
    ),
  })
  .refine(
    (data) => {
      const startDate = new Date(data.startDate);
      const endDate = new Date(data.endDate);
      const diffMs = endDate.getTime() - startDate.getTime();
      const oneWeekMs = 1000 * 60 * 60 * 24 * 7; // 1 week in milliseconds

      return diffMs >= oneWeekMs;
    },
    {
      message:
        'The difference between the Start Date and End Date must be at least 1 week.',
      path: ['endDate'],
    },
  ) satisfies z.ZodType<CreateSeasonDto>;
