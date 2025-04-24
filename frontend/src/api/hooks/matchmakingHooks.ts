import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import { JoinQueueDto } from '../../types/dtos/requests/matchmaking/joinQueueDto';
import { LeaveQueueDto } from '../../types/dtos/requests/matchmaking/leaveQueueDto';
import { MatchStatusDto } from '../../types/dtos/responses/matchmaking/matchStatusDto';
import { QueueTeamsCountDto } from '../../types/dtos/responses/matchmaking/queueTeamsCountDto';
import { getRequest, postRequest } from '../httpRequests';

export const useJoinQueue = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: JoinQueueDto) =>
      postRequest({
        url: ENDPOINTS.MATCHMAKING.JOIN_QUEUE,
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
        url: ENDPOINTS.MATCHMAKING.LEAVE_QUEUE,
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

export const useGetQueueTeamsCount = (seasonId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.MATCHMAKING.QUEUE_TEAMS_COUNT(seasonId),
    queryFn: () =>
      getRequest<QueueTeamsCountDto>({
        url: ENDPOINTS.MATCHMAKING.GET_QUEUED_TEAMS_COUNT(seasonId),
      }),
    enabled: !!seasonId,
  });

export const useCheckMatchStatus = (teamId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.MATCHMAKING.CHECK_MATCH_STATUS(teamId),
    queryFn: () =>
      getRequest<MatchStatusDto>({
        url: ENDPOINTS.MATCHMAKING.GET_QUEUE_STATUS(teamId),
      }),
    enabled: !!teamId,
  });
