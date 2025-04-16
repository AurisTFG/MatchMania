import { z } from 'zod';
import { UpdateResultDto } from '../../types/dtos/requests/results/updateResultDto';

const isWithinTwoDays = (date: Date) => {
  const minDate = new Date();
  minDate.setDate(minDate.getDate() - 2);

  const maxDate = new Date();
  maxDate.setDate(maxDate.getDate() + 2);

  return date > minDate && date < maxDate;
};

export const updateResultDtoValidator = z
  .object({
    matchStartDate: z.date().refine((date) => isWithinTwoDays(date), {
      message: 'Match Start Date must be within 2 days from now.',
    }),

    matchEndDate: z.date().refine((date) => isWithinTwoDays(date), {
      message: 'Match End Date must be within 2 days from now.',
    }),

    score: z.string().refine(
      (score) => {
        const scoreInt = parseInt(score);
        return !isNaN(scoreInt) && scoreInt >= 0 && scoreInt <= 100;
      },
      { message: 'Score must be between 0 and 100.' },
    ),

    opponentScore: z.string().refine(
      (score) => {
        const scoreInt = parseInt(score);
        return !isNaN(scoreInt) && scoreInt >= 0 && scoreInt <= 100;
      },
      { message: 'Opponent Score must be between 0 and 100.' },
    ),
  })
  .refine(
    (data) => {
      const startDate = new Date(data.matchStartDate);
      const endDate = new Date(data.matchEndDate);
      const diffInHours =
        (endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60);
      return diffInHours >= 3;
    },
    {
      message:
        'Match End Date must be at least 3 hours later than the Match Start Date.',
      path: ['matchEndDate'],
    },
  )
  .refine(
    (data) => {
      return new Date(data.matchEndDate) > new Date(data.matchStartDate);
    },
    {
      message: 'Match End Date must be later than the Match Start Date.',
      path: ['matchEndDate'],
    },
  ) satisfies z.ZodType<UpdateResultDto>;
