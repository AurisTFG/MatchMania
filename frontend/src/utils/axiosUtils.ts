import { AxiosError } from 'axios';
import { ErrorDto } from '../types';

export const getErrorFromAxiosError = (error: AxiosError): Error => {
  const errorDto = error.response?.data as ErrorDto;

  if (errorDto.error) {
    return new Error(errorDto.error);
  }

  if (error.response?.status) {
    return new Error(
      `Unhandled error. Status code: ${String(error.response.status)}`,
    );
  }

  return new Error(
    `Unexpected error format: ${JSON.stringify(error.response)}`,
  );
};
