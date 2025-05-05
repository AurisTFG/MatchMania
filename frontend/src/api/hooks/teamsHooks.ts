import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import ENDPOINTS from 'constants/endpoints';
import QUERY_KEYS from 'constants/queryKeys';
import { CreateTeamDto } from 'types/dtos/requests/teams/createTeamDto';
import { UpdateTeamDto } from 'types/dtos/requests/teams/updateTeamDto';
import { TeamDto } from 'types/dtos/responses/teams/teamDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export function useFetchTeams() {
  return useQuery({
    queryKey: QUERY_KEYS.TEAMS.ALL,
    queryFn: () =>
      getRequest<TeamDto[]>({
        url: ENDPOINTS.TEAMS.ROOT,
      }),
  });
}

export function useFetchTeam(teamId: string) {
  return useQuery({
    queryKey: QUERY_KEYS.TEAMS.BY_ID(teamId),
    queryFn: () =>
      getRequest<TeamDto>({
        url: ENDPOINTS.TEAMS.BY_ID(teamId),
      }),
    enabled: !!teamId,
  });
}

export function useCreateTeam() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateTeamDto) =>
      postRequest({
        url: ENDPOINTS.TEAMS.ROOT,
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Team created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL,
      });
    },
  });
}

export function useUpdateTeam() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      teamId,
      payload,
    }: {
      teamId: string;
      payload: UpdateTeamDto;
    }) =>
      patchRequest({
        url: ENDPOINTS.TEAMS.BY_ID(teamId),
        body: payload,
      }),
    onSuccess: async (_, { teamId }) => {
      toast.success('Team updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.BY_ID(teamId),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL,
      });
    },
  });
}

export function useDeleteTeam() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (teamId: string) =>
      deleteRequest({ url: ENDPOINTS.TEAMS.BY_ID(teamId) }),
    onSuccess: async () => {
      toast.success('Team deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL,
      });
    },
  });
}
