import { useQuery } from '@tanstack/react-query';
import ENDPOINTS from 'constants/endpoints';
import QUERY_KEYS from 'constants/queryKeys';
import { PlayerMinimalDto } from 'types/dtos/responses/players/playerMinimalDto';
import { getRequest } from '../httpRequests';

export const useFetchPlayers = () =>
  useQuery({
    queryKey: QUERY_KEYS.PLAYERS.ALL,
    queryFn: () =>
      getRequest<PlayerMinimalDto[]>({ url: ENDPOINTS.PLAYERS.ROOT }),
  });
