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

export const useFetchTeams = (seasonId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.ALL(seasonId),
    queryFn: () =>
      getRequest<TeamDto[]>({ url: ENDPOINTS.TEAMS.ROOT(seasonId) }),
    enabled: !!seasonId,
  });

export const useFetchTeam = (seasonId: string, teamId: string) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.BY_ID(seasonId, teamId),
    queryFn: () =>
      getRequest<TeamDto>({ url: ENDPOINTS.TEAMS.BY_ID(seasonId, teamId) }),
    enabled: !!seasonId && !!teamId,
  });

export const useCreateTeam = (seasonId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateTeamDto) =>
      postRequest({
        url: ENDPOINTS.TEAMS.ROOT(seasonId),
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Team created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(seasonId),
      });
    },
  });
};

export const useUpdateTeam = (seasonId: string) => {
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
        url: ENDPOINTS.TEAMS.BY_ID(seasonId, teamId),
        body: payload,
      }),
    onSuccess: async (_, { teamId }) => {
      toast.success('Team updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.BY_ID(seasonId, teamId),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(seasonId),
      });
    },
  });
};

export const useDeleteTeam = (seasonId: string) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (teamId: string) =>
      deleteRequest({ url: ENDPOINTS.TEAMS.BY_ID(seasonId, teamId) }),
    onSuccess: async () => {
      toast.success('Team deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(seasonId),
      });
    },
  });
};
