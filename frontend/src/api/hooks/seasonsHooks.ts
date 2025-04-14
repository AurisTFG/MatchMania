import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import { Season } from '../../types';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchSeasons = () =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.ALL,
    queryFn: () => getRequest<Season[]>({ url: ENDPOINTS.SEASONS.ROOT }),
  });

export const useFetchSeason = (seasonID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.BY_ID(seasonID),
    queryFn: () =>
      getRequest<Season>({ url: ENDPOINTS.SEASONS.BY_ID(seasonID) }),
    enabled: !!seasonID,
  });

export const useCreateSeason = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (season: { name: string; startDate: Date; endDate: Date }) =>
      postRequest<Season>({ url: ENDPOINTS.SEASONS.ROOT, body: season }),
    onSuccess: async () => {
      toast.success('Season created successfully');

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
      patchRequest<Season>({
        url: ENDPOINTS.SEASONS.BY_ID(seasonID),
        body: season,
      }),
    onSuccess: async (_, { seasonID }) => {
      toast.success('Season updated successfully');

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
      deleteRequest({ url: ENDPOINTS.SEASONS.BY_ID(seasonID) }),
    onSuccess: async () => {
      toast.success('Season deleted successfully');

      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};
