import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { getRequest, postRequest } from 'api/httpRequests';
import ENDPOINTS from 'constants/endpoints';
import QUERY_KEYS from 'constants/queryKeys';
import { EndMatchDto } from 'types/dtos/requests/matches/endMatchDto';
import { MatchDto } from 'types/dtos/responses/matches/matchDto';

export function useFetchMatches() {
  return useQuery({
    queryKey: QUERY_KEYS.MATCHMAKING.MATCHES.ALL,
    queryFn: () =>
      getRequest<MatchDto[]>({
        url: ENDPOINTS.MATCHMAKING.MATCHES.ROOT,
      }),
    staleTime: 0,
    gcTime: 0,
    refetchInterval: 5000, // 5 seconds
  });
}

export const useEndMatch = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: EndMatchDto) =>
      postRequest({
        url: ENDPOINTS.MATCHMAKING.MATCHES.END,
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Successfully ended the match');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.MATCHMAKING.ROOT,
      });
    },
  });
};
