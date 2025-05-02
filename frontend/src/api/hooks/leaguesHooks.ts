import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from 'constants/endpoints';
import { QUERY_KEYS } from 'constants/queryKeys';
import { CreateLeagueDto } from 'types/dtos/requests/leagues/createLeagueDto';
import { UpdateLeagueDto } from 'types/dtos/requests/leagues/updateLeagueDto';
import { LeagueDto } from 'types/dtos/responses/leagues/leagueDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchLeagues = () =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.ALL,
    queryFn: () => getRequest<LeagueDto[]>({ url: ENDPOINTS.SEASONS.ROOT }),
  });

export const useFetchLeague = (leagueId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.SEASONS.BY_ID(leagueId),
    queryFn: () =>
      getRequest<LeagueDto>({ url: ENDPOINTS.SEASONS.BY_ID(leagueId) }),
    enabled: !!leagueId,
  });

export const useCreateLeague = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateLeagueDto) =>
      postRequest({ url: ENDPOINTS.SEASONS.ROOT, body: payload }),
    onSuccess: async () => {
      toast.success('League created successfully');

      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};

export const useUpdateLeague = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      leagueId,
      payload,
    }: {
      leagueId: string;
      payload: UpdateLeagueDto;
    }) =>
      patchRequest({
        url: ENDPOINTS.SEASONS.BY_ID(leagueId),
        body: payload,
      }),
    onSuccess: async (_, { leagueId }) => {
      toast.success('League updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.SEASONS.BY_ID(leagueId),
      });
      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};

export const useDeleteLeague = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (leagueId: string) =>
      deleteRequest({ url: ENDPOINTS.SEASONS.BY_ID(leagueId) }),
    onSuccess: async () => {
      toast.success('League deleted successfully');

      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.SEASONS.ALL });
    },
  });
};
