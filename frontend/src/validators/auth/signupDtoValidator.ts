import { z } from 'zod';
import { SignupDto } from '../../types/dtos/requests/auth/signupDto';

export const signupDtoValidator = z
  .object({
    username: z
      .string()
      .min(3, { message: 'Username must be at least 3 characters long.' })
      .max(100, { message: 'Username must be at most 100 characters long.' }),

    email: z
      .string()
      .email({ message: 'Email must be a valid email address.' }),

    password: z
      .string()
      .min(8, { message: 'Password must be at least 8 characters long.' })
      .max(100, { message: 'Password must be at most 100 characters long.' }),

    confirmPassword: z.string(),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: 'Passwords do not match.',
    path: ['confirmPassword'],
  }) satisfies z.ZodType<SignupDto>;
