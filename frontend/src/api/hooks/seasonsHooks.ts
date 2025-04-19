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

export const useFetchSeason = (seasonId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.BY_ID(seasonId),
    queryFn: () =>
      getRequest<SeasonDto>({ url: ENDPOINTS.SEASONS.BY_ID(seasonId) }),
    enabled: !!seasonId,
  });

export const useCreateSeason = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateSeasonDto) =>
      postRequest({ url: ENDPOINTS.SEASONS.ROOT, body: payload }),
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
      seasonId,
      payload,
    }: {
      seasonId: string;
      payload: UpdateSeasonDto;
    }) =>
      patchRequest({
        url: ENDPOINTS.SEASONS.BY_ID(seasonId),
        body: payload,
      }),
    onSuccess: async (_, { seasonId }) => {
      toast.success('Season updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.SEASONS.BY_ID(seasonId),
      });
      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};

export const useDeleteSeason = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (seasonId: string) =>
      deleteRequest({ url: ENDPOINTS.SEASONS.BY_ID(seasonId) }),
    onSuccess: async () => {
      toast.success('Season deleted successfully');

      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};
