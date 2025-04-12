import { QueryClient } from '@tanstack/react-query';

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
});
