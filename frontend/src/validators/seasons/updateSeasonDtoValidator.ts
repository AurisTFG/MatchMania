import { z } from 'zod';
import { UpdateSeasonDto } from '../../types/dtos/requests/seasons/updateSeasonDto';
import { createSeasonDtoValidator } from './createSeasonDtoValidator';

export const updateSeasonDtoValidator =
  createSeasonDtoValidator satisfies z.ZodType<UpdateSeasonDto>;
