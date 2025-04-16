import { z } from 'zod';
import { LoginDto } from '../../types/dtos/requests/auth/loginDto';

export const loginDtoValidator = z.object({
  email: z
    .string()
    .min(1, { message: 'Email is required.' })
    .email({ message: 'Email must be a valid email address.' }),

  password: z
    .string()
    .min(1, { message: 'Password is required.' })
    .min(5, { message: 'Password must be at least 5 characters long.' })
    .max(100, { message: 'Password must be at most 100 characters long.' }),
}) satisfies z.ZodType<LoginDto>;
