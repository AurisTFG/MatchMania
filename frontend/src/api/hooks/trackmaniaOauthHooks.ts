import { useQuery } from '@tanstack/react-query';
import { ENDPOINTS } from 'constants/endpoints';
import { QUERY_KEYS } from 'constants/queryKeys';
import { TrackmaniaOAuthUrlDto } from 'types/dtos/responses/trackmaniaOAuth/trackmaniaOAuthUrlDto';
import { getRequest } from '../httpRequests';

export const useGetTrackmaniaOAuthUrl = () =>
  useQuery({
    queryKey: QUERY_KEYS.TRACKMANIA.OAUTH.URL,
    queryFn: () =>
      getRequest<TrackmaniaOAuthUrlDto>({
        url: ENDPOINTS.TRACKMANIA.OAUTH.URL,
      }),
  });
