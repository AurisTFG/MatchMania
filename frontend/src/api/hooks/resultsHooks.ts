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

export const useFetchResults = (seasonId: string, teamId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.RESULTS.ALL(seasonId, teamId),
    queryFn: () =>
      getRequest<ResultDto[]>({
        url: ENDPOINTS.RESULTS.ROOT(seasonId, teamId),
      }),
    enabled: !!seasonId && !!teamId,
  });

export const useFetchResult = (
  seasonId: string,
  teamId: string,
  resultId: string,
) =>
  useQuery({
    queryKey: QUERY_KEYS.RESULTS.BY_ID(seasonId, teamId, resultId),
    queryFn: () =>
      getRequest<ResultDto>({
        url: ENDPOINTS.RESULTS.BY_ID(seasonId, teamId, resultId),
      }),
    enabled: !!seasonId && !!teamId && !!resultId,
  });

export const useCreateResult = (seasonId: string, teamId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateResultDto) =>
      postRequest({
        url: ENDPOINTS.RESULTS.ROOT(seasonId, teamId),
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Result created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(seasonId, teamId),
      });
    },
  });
};

export const useUpdateResult = (seasonId: string, teamId: string) => {
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
        url: ENDPOINTS.RESULTS.BY_ID(seasonId, teamId, resultId),
        body: payload,
      }),
    onSuccess: async (_, { resultId }) => {
      toast.success('Result updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.BY_ID(seasonId, teamId, resultId),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(seasonId, teamId),
      });
    },
  });
};

export const useDeleteResult = (seasonId: string, teamId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (resultId: string) =>
      deleteRequest({
        url: ENDPOINTS.RESULTS.BY_ID(seasonId, teamId, resultId),
      }),
    onSuccess: async () => {
      toast.success('Result deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(seasonId, teamId),
      });
    },
  });
};
