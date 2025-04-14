import { AxiosError } from 'axios';
import { ErrorDto } from '../types';

export const getErrorMessageFromAxiosError = (error: AxiosError): string => {
  const errorDto = error.response?.data as ErrorDto;

  if (errorDto.error) {
    return errorDto.error;
  }

  if (error.response?.status) {
    return `Unhandled error. Status code: ${String(error.response.status)}`;
  }

  return `Unexpected error format: ${JSON.stringify(error.response)}`;
};
