import { useQuery } from '@tanstack/react-query';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import { UserDto } from '../../types/dtos/responses/users/userDto';
import { getRequest } from '../httpRequests';

export const useFetchUsers = () =>
  useQuery({
    queryKey: QUERY_KEYS.USERS.ALL,
    queryFn: () => getRequest<UserDto[]>({ url: ENDPOINTS.USERS.ROOT }),
  });
