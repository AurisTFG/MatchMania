import { MutationCache, QueryCache, QueryClient } from '@tanstack/react-query';
import { onMutationError, onQueryError } from 'utils/queryClientUtils';

export const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 1, // 1 minute
      gcTime: 1000 * 60 * 10, // 10 minutes
      retry: false,
      refetchOnReconnect: 'always',
    },
    mutations: {
      retry: false,
    },
  },
  queryCache: new QueryCache({
    onError: (error, query) => {
      onQueryError(error, query).catch(console.error);
    },
  }),
  mutationCache: new MutationCache({
    onError: onMutationError,
  }),
});
