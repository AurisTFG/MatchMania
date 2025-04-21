import { z } from 'zod';
import { CreateResultDto } from '../../types/dtos/requests/results/createResultDto';

const isWithinTwoDays = (date: Date) => {
  const minDate = new Date();
  minDate.setDate(minDate.getDate() - 2);

  const maxDate = new Date();
  maxDate.setDate(maxDate.getDate() + 2);

  return date > minDate && date < maxDate;
};

export const createResultDtoValidator = z
  .object({
    startDate: z.date().refine((date) => isWithinTwoDays(date), {
      message: 'Match Start Date must be within 2 days from now.',
    }),

    endDate: z.date().refine((date) => isWithinTwoDays(date), {
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

    opponentTeamId: z.string().uuid({
      message: 'Opponent Team ID must be a valid UUID.',
    }),
  })
  .refine(
    (data) => {
      const startDate = new Date(data.startDate);
      const endDate = new Date(data.endDate);
      const diffInHours =
        (endDate.getTime() - startDate.getTime()) / (1000 * 60 * 60);
      return diffInHours >= 3;
    },
    {
      message:
        'Match End Date must be at least 3 hours later than the Match Start Date.',
      path: ['endDate'],
    },
  )
  .refine(
    (data) => {
      return new Date(data.endDate) > new Date(data.startDate);
    },
    {
      message: 'Match End Date must be later than the Match Start Date.',
      path: ['endDate'],
    },
  ) satisfies z.ZodType<CreateResultDto>;
