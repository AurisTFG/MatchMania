import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from 'constants/endpoints';
import { QUERY_KEYS } from 'constants/queryKeys';
import { CreateResultDto } from 'types/dtos/requests/results/createResultDto';
import { UpdateResultDto } from 'types/dtos/requests/results/updateResultDto';
import { ResultDto } from 'types/dtos/responses/results/resultDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchResults = (leagueId: string, teamId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.RESULTS.ALL(leagueId, teamId),
    queryFn: () =>
      getRequest<ResultDto[]>({
        url: ENDPOINTS.RESULTS.ROOT(leagueId, teamId),
      }),
    enabled: !!leagueId && !!teamId,
  });

export const useFetchResult = (
  leagueId: string,
  teamId: string,
  resultId: string,
) =>
  useQuery({
    queryKey: QUERY_KEYS.RESULTS.BY_ID(leagueId, teamId, resultId),
    queryFn: () =>
      getRequest<ResultDto>({
        url: ENDPOINTS.RESULTS.BY_ID(leagueId, teamId, resultId),
      }),
    enabled: !!leagueId && !!teamId && !!resultId,
  });

export const useCreateResult = (leagueId: string, teamId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateResultDto) =>
      postRequest({
        url: ENDPOINTS.RESULTS.ROOT(leagueId, teamId),
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Result created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(leagueId, teamId),
      });
    },
  });
};

export const useUpdateResult = (leagueId: string, teamId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      resultId,
      payload,
    }: {
      resultId: string;
      payload: UpdateResultDto;
    }) =>
      patchRequest({
        url: ENDPOINTS.RESULTS.BY_ID(leagueId, teamId, resultId),
        body: payload,
      }),
    onSuccess: async (_, { resultId }) => {
      toast.success('Result updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.BY_ID(leagueId, teamId, resultId),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(leagueId, teamId),
      });
    },
  });
};

export const useDeleteResult = (leagueId: string, teamId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (resultId: string) =>
      deleteRequest({
        url: ENDPOINTS.RESULTS.BY_ID(leagueId, teamId, resultId),
      }),
    onSuccess: async () => {
      toast.success('Result deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(leagueId, teamId),
      });
    },
  });
};
