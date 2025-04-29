import { Mutation, Query } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { toast } from 'sonner';
import axiosClient from 'configs/axiosClient';
import { ENDPOINTS } from 'constants/endpoints';
import { ROUTE_PATHS } from 'constants/route_paths';
import { getErrorMessageFromAxiosError } from './errorUtils';

async function handleError(
  error: unknown,
  retryCallback: () => Promise<void>,
  skipNavigation = false,
) {
  const axiosError = error as AxiosError;

  if (!axiosError.isAxiosError) {
    toast.error('An unknown error occurred.');
    return;
  }

  if (axiosError.response?.status === 401) {
    try {
      await axiosClient.post(ENDPOINTS.AUTH.REFRESH, null);
      await retryCallback();
      return;
    } catch {
      if (
        !window.location.pathname.startsWith(ROUTE_PATHS.LOGIN) &&
        !skipNavigation
      ) {
        const errorMessage = encodeURIComponent(
          'Session expired. Please log in again.',
        );
        window.location.href = `${ROUTE_PATHS.LOGIN}?error=${errorMessage}`;
      }
      return;
    }
  }

  const message = getErrorMessageFromAxiosError(axiosError);
  toast.error(`An error occurred: ${message}`);
}

export const onQueryError = async (
  error: unknown,
  query: Query<unknown, unknown, unknown>,
) => {
  const skipNavigation = (error as AxiosError).config?.url?.includes(
    ENDPOINTS.AUTH.ME,
  );

  await handleError(
    error,
    async () => {
      await query.fetch();
    },
    skipNavigation,
  );
};

export const onMutationError = async (
  error: unknown,
  variables: unknown,
  _context: unknown,
  mutation: Mutation<unknown, unknown>,
) => {
  await handleError(error, async () => {
    await mutation.execute(variables);
  });
};
