import { AxiosError } from 'axios';
import { ErrorDto } from 'types/dtos/responses/errors/errorDto';

export const getErrorMessageFromAxiosError = (error: AxiosError): string => {
  const errorDto = error.response?.data as ErrorDto | undefined;

  if (errorDto?.message) {
    return errorDto.message;
  }

  if (error.response?.status) {
    return `Unhandled error. Status code: ${String(error.response.status)}`;
  }

  return `Unexpected error: ${JSON.stringify(error.message)}`;
};
