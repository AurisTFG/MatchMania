import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { toast } from 'sonner';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import { CreateTeamDto } from '../../types/dtos/requests/teams/createTeamDto';
import { UpdateTeamDto } from '../../types/dtos/requests/teams/updateTeamDto';
import { TeamDto } from '../../types/dtos/responses/teams/teamDto';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchTeams = (seasonID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.ALL(seasonID),
    queryFn: () =>
      getRequest<TeamDto[]>({ url: ENDPOINTS.TEAMS.ROOT(seasonID) }),
    enabled: !!seasonID,
  });

export const useFetchTeam = (seasonID: number, teamID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.BY_ID(seasonID, teamID),
    queryFn: () =>
      getRequest<TeamDto>({ url: ENDPOINTS.TEAMS.BY_ID(seasonID, teamID) }),
    enabled: !!seasonID && !!teamID,
  });

export const useCreateTeam = (seasonID: number) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (payload: CreateTeamDto) =>
      postRequest<TeamDto>({
        url: ENDPOINTS.TEAMS.ROOT(seasonID),
        body: payload,
      }),
    onSuccess: async () => {
      toast.success('Team created successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(seasonID),
      });
    },
  });
};

export const useUpdateTeam = (seasonID: number) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({
      teamID,
      payload,
    }: {
      teamID: number;
      payload: UpdateTeamDto;
    }) =>
      patchRequest<TeamDto>({
        url: ENDPOINTS.TEAMS.BY_ID(seasonID, teamID),
        body: payload,
      }),
    onSuccess: async (_, { teamID }) => {
      toast.success('Team updated successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.BY_ID(seasonID, teamID),
      });
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(seasonID),
      });
    },
  });
};

export const useDeleteTeam = (seasonID: number) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (teamID: number) =>
      deleteRequest({ url: ENDPOINTS.TEAMS.BY_ID(seasonID, teamID) }),
    onSuccess: async () => {
      toast.success('Team deleted successfully');

      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(seasonID),
      });
    },
  });
};
