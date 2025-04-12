import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import {
  deleteRequest,
  getRequest,
  patchRequest,
  postRequest,
} from '../httpRequests';

export const useFetchTeams = (seasonID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.ALL(seasonID),
    queryFn: () => getRequest({ url: ENDPOINTS.TEAMS.ROOT(seasonID) }),
    enabled: !!seasonID,
  });

export const useFetchTeam = (seasonID: number, teamID: number) =>
  useQuery({
    queryKey: QUERY_KEYS.TEAMS.BY_ID(seasonID, teamID),
    queryFn: () => getRequest({ url: ENDPOINTS.TEAMS.BY_ID(seasonID, teamID) }),
    enabled: !!seasonID && !!teamID,
  });

export const useCreateTeam = (seasonID: number) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (team: { name: string }) =>
      postRequest({
        url: ENDPOINTS.TEAMS.ROOT(seasonID),
        body: team,
      }),
    onSuccess: async () => {
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
      team,
    }: {
      teamID: number;
      team: { name: string };
    }) =>
      patchRequest({
        url: ENDPOINTS.TEAMS.BY_ID(seasonID, teamID),
        body: team,
      }),
    onSuccess: async (_, { teamID }) => {
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
      deleteRequest({
        url: ENDPOINTS.TEAMS.BY_ID(seasonID, teamID),
      }),
    onSuccess: async () => {
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.TEAMS.ALL(seasonID),
      });
    },
  });
};
