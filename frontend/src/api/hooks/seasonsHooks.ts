import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import { CreateSeasonDto } from '../../types/dtos/requests/seasons/createSeasonDto';
import { UpdateSeasonDto } from '../../types/dtos/requests/seasons/updateSeasonDto';
import { SeasonDto } from '../../types/dtos/responses/seasons/seasonDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchSeasons = () =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.ALL,
    queryFn: () => getRequest<SeasonDto[]>({ url: ENDPOINTS.SEASONS.ROOT }),
  });

export const useFetchSeason = (seasonID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.BY_ID(seasonID),
    queryFn: () =>
      getRequest<SeasonDto>({ url: ENDPOINTS.SEASONS.BY_ID(seasonID) }),
    enabled: !!seasonID,
  });

export const useCreateSeason = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateSeasonDto) =>
      postRequest<SeasonDto>({ url: ENDPOINTS.SEASONS.ROOT, body: payload }),
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
      payload,
    }: {
      seasonID: number;
      payload: UpdateSeasonDto;
    }) =>
      patchRequest<SeasonDto>({
        url: ENDPOINTS.SEASONS.BY_ID(seasonID),
        body: payload,
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
