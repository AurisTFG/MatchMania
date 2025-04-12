import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchSeasons = () =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.ALL,
    queryFn: () => getRequest({ url: ENDPOINTS.SEASONS.ROOT }),
  });

export const useFetchSeason = (seasonID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.BY_ID(seasonID),
    queryFn: () => getRequest({ url: ENDPOINTS.SEASONS.BY_ID(seasonID) }),
    enabled: !!seasonID,
  });

export const useCreateSeason = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (season: { name: string; startDate: Date; endDate: Date }) =>
      postRequest({
        url: ENDPOINTS.SEASONS.ROOT,
        body: season,
      }),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};

export const useUpdateSeason = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      seasonID,
      season,
    }: {
      seasonID: number;
      season: { name: string; startDate: Date; endDate: Date };
    }) =>
      patchRequest({
        url: ENDPOINTS.SEASONS.BY_ID(seasonID),
        body: season,
      }),
    onSuccess: async (_, { seasonID }) => {
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.SEASONS.BY_ID(seasonID),
      });
      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};

export const useDeleteSeason = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (seasonID: number) =>
      deleteRequest({
        url: ENDPOINTS.SEASONS.BY_ID(seasonID),
      }),
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};
