import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import { CreateResultDto } from '../../types/dtos/requests/results/createResultDto';
import { UpdateResultDto } from '../../types/dtos/requests/results/updateResultDto';
import { ResultDto } from '../../types/dtos/responses/results/resultDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchResults = (seasonID: number, teamID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.RESULTS.ALL(seasonID, teamID),
    queryFn: () =>
      getRequest<ResultDto[]>({
        url: ENDPOINTS.RESULTS.ROOT(seasonID, teamID),
      }),
    enabled: !!seasonID && !!teamID,
  });

export const useFetchResult = (
  seasonID: number,
  teamID: number,
  resultID: number,
) =>
  useQuery({
    queryKey: QUERY_KEYS.RESULTS.BY_ID(seasonID, teamID, resultID),
    queryFn: () =>
      getRequest<ResultDto>({
        url: ENDPOINTS.RESULTS.BY_ID(seasonID, teamID, resultID),
      }),
    enabled: !!seasonID && !!teamID && !!resultID,
  });

export const useCreateResult = (seasonID: number, teamID: number) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateResultDto) =>
      postRequest<ResultDto>({
        url: ENDPOINTS.RESULTS.ROOT(seasonID, teamID),
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Result created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(seasonID, teamID),
      });
    },
  });
};

export const useUpdateResult = (seasonID: number, teamID: number) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      resultID,
      payload,
    }: {
      resultID: number;
      payload: UpdateResultDto;
    }) =>
      patchRequest<ResultDto>({
        url: ENDPOINTS.RESULTS.BY_ID(seasonID, teamID, resultID),
        body: payload,
      }),
    onSuccess: async (_, { resultID }) => {
      toast.success('Result updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.BY_ID(seasonID, teamID, resultID),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(seasonID, teamID),
      });
    },
  });
};

export const useDeleteResult = (seasonID: number, teamID: number) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (resultID: number) =>
      deleteRequest({
        url: ENDPOINTS.RESULTS.BY_ID(seasonID, teamID, resultID),
      }),
    onSuccess: async () => {
      toast.success('Result deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.RESULTS.ALL(seasonID, teamID),
      });
    },
  });
};
