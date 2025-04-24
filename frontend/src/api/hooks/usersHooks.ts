import { useQuery } from '@tanstack/react-query';
import { ENDPOINTS } from 'constants/endpoints';
import { QUERY_KEYS } from 'constants/queryKeys';
import { UserMinimalDto } from 'types/dtos/responses/users/userMinimalDto';
import { getRequest } from '../httpRequests';

export const useFetchUsers = () =>
  useQuery({
    queryKey: QUERY_KEYS.USERS.ALL,
    queryFn: () => getRequest<UserMinimalDto[]>({ url: ENDPOINTS.USERS.ROOT }),
  });
