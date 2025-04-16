import { z } from 'zod';
import { CreateSeasonDto } from '../../types/dtos/requests/seasons/createSeasonDto';

const isValidStartDate = (date: Date) => {
  const today = new Date();
  const yearAgo = new Date(today);
  yearAgo.setFullYear(today.getFullYear() - 1);

  const tenYearsFromNow = new Date(today);
  tenYearsFromNow.setFullYear(today.getFullYear() + 10);

  return date > yearAgo && date < tenYearsFromNow;
};

const isValidEndDate = (date: Date) => {
  const today = new Date();
  const tenYearsFromNow = new Date(today);
  tenYearsFromNow.setFullYear(today.getFullYear() + 10);

  return date > today && date < tenYearsFromNow;
};

export const createSeasonDtoValidator = z
  .object({
    name: z
      .string()
      .min(1, { message: 'Name is required.' })
      .min(3, { message: 'Name must be at least 3 characters long.' })
      .max(100, { message: 'Name can be up to 100 characters long.' }),

    startDate: z.date().refine((date) => isValidStartDate(date), {
      message:
        'Start Date must be a valid date between 1 year ago and 10 years from now.',
    }),

    endDate: z.date().refine((date) => isValidEndDate(date), {
      message:
        'End Date must be a valid date between today and 10 years from now.',
    }),
  })
  .refine(
    (data) => {
      return new Date(data.endDate) > new Date(data.startDate);
    },
    {
      message: 'End Date must be later than the Start Date.',
      path: ['endDate'],
    },
  )
  .refine(
    (data) => {
      const startDate = new Date(data.startDate);
      const endDate = new Date(data.endDate);
      const diffMs = endDate.getTime() - startDate.getTime();
      const oneMonthMs = 1000 * 60 * 60 * 24 * 30; // 1 month in milliseconds

      return diffMs >= oneMonthMs;
    },
    {
      message:
        'The difference between the Start Date and End Date must be at least 1 month.',
      path: ['endDate'],
    },
  ) satisfies z.ZodType<CreateSeasonDto>;
