import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from 'constants/endpoints';
import { QUERY_KEYS } from 'constants/queryKeys';
import { CreateTeamDto } from 'types/dtos/requests/teams/createTeamDto';
import { UpdateTeamDto } from 'types/dtos/requests/teams/updateTeamDto';
import { TeamDto } from 'types/dtos/responses/teams/teamDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchTeams = (leagueId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.ALL(leagueId),
    queryFn: () =>
      getRequest<TeamDto[]>({ url: ENDPOINTS.TEAMS.ROOT(leagueId) }),
    enabled: !!leagueId,
  });

export const useFetchTeam = (leagueId: string, teamId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.BY_ID(leagueId, teamId),
    queryFn: () =>
      getRequest<TeamDto>({ url: ENDPOINTS.TEAMS.BY_ID(leagueId, teamId) }),
    enabled: !!leagueId && !!teamId,
  });

export const useCreateTeam = (leagueId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateTeamDto) =>
      postRequest({
        url: ENDPOINTS.TEAMS.ROOT(leagueId),
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Team created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(leagueId),
      });
    },
  });
};

export const useUpdateTeam = (leagueId: string) => {
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
        url: ENDPOINTS.TEAMS.BY_ID(leagueId, teamId),
        body: payload,
      }),
    onSuccess: async (_, { teamId }) => {
      toast.success('Team updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.BY_ID(leagueId, teamId),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(leagueId),
      });
    },
  });
};

export const useDeleteTeam = (leagueId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (teamId: string) =>
      deleteRequest({ url: ENDPOINTS.TEAMS.BY_ID(leagueId, teamId) }),
    onSuccess: async () => {
      toast.success('Team deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(leagueId),
      });
    },
  });
};
