import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import ENDPOINTS from 'constants/endpoints';
import QUERY_KEYS from 'constants/queryKeys';
import { CreateResultDto } from 'types/dtos/requests/results/createResultDto';
import { UpdateResultDto } from 'types/dtos/requests/results/updateResultDto';
import { ResultDto } from 'types/dtos/responses/results/resultDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export function useFetchResults() {
  return useQuery({
    queryKey: QUERY_KEYS.RESULTS.ALL,
    queryFn: () =>
      getRequest<ResultDto[]>({
        url: ENDPOINTS.RESULTS.ROOT,
      }),
  });
}

export function useFetchResult(resultId: string) {
  return useQuery({
    queryKey: QUERY_KEYS.RESULTS.BY_ID(resultId),
    queryFn: () =>
      getRequest<ResultDto>({
        url: ENDPOINTS.RESULTS.BY_ID(resultId),
      }),
    enabled: !!resultId,
  });
}
export function useCreateResult() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateResultDto) =>
      postRequest({
        url: ENDPOINTS.RESULTS.ROOT,
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Result created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL,
      });
    },
  });
}

export function useUpdateResult() {
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
        url: ENDPOINTS.RESULTS.BY_ID(resultId),
        body: payload,
      }),
    onSuccess: async (_, { resultId }) => {
      toast.success('Result updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.BY_ID(resultId),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL,
      });
    },
  });
}

export function useDeleteResult() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (resultId: string) =>
      deleteRequest({
        url: ENDPOINTS.RESULTS.BY_ID(resultId),
      }),
    onSuccess: async () => {
      toast.success('Result deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL,
      });
    },
  });
}
