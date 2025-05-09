import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import ENDPOINTS from 'constants/endpoints';
import QUERY_KEYS from 'constants/queryKeys';
import { JoinQueueDto } from 'types/dtos/requests/queues/joinQueueDto';
import { LeaveQueueDto } from 'types/dtos/requests/queues/leaveQueueDto';
import { QueueDto } from 'types/dtos/responses/queues/queueDto';
import { getRequest, postRequest } from '../httpRequests';

export function useFetchQueues() {
  return useQuery({
    queryKey: QUERY_KEYS.MATCHMAKING.QUEUES.ALL,
    queryFn: () =>
      getRequest<QueueDto[]>({
        url: ENDPOINTS.MATCHMAKING.QUEUES.ROOT,
      }),
    staleTime: 0,
    gcTime: 0,
    refetchInterval: 5000, // 5 seconds
  });
}

export const useJoinQueue = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: JoinQueueDto) =>
      postRequest({
        url: ENDPOINTS.MATCHMAKING.QUEUES.JOIN,
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Successfully joined the queue');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.MATCHMAKING.ROOT,
      });
    },
  });
};

export const useLeaveQueue = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: LeaveQueueDto) =>
      postRequest({
        url: ENDPOINTS.MATCHMAKING.QUEUES.LEAVE,
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Successfully left the queue');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.MATCHMAKING.ROOT,
      });
    },
  });
};
