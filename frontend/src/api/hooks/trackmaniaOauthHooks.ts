import { useMutation } from '@tanstack/react-query';
import ENDPOINTS from 'constants/endpoints';
import { TrackmaniaOAuthUrlDto } from 'types/dtos/responses/trackmaniaOAuth/trackmaniaOAuthUrlDto';
import { getRequest } from '../httpRequests';

export const useGetTrackmaniaOAuthUrl = () =>
  useMutation({
    mutationFn: () =>
      getRequest<TrackmaniaOAuthUrlDto>({
        url: ENDPOINTS.TRACKMANIA.OAUTH.URL,
      }),
  });
