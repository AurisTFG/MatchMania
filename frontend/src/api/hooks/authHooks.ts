import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { toast } from 'sonner';
import { ENDPOINTS } from '../../constants/endpoints';
import { QUERY_KEYS } from '../../constants/queryKeys';
import { ROUTES } from '../../constants/routes';
import { LoginDto } from '../../types/dtos/requests/auth/loginDto';
import { SignupDto } from '../../types/dtos/requests/auth/signupDto';
import { UserDto } from '../../types/dtos/responses/users/userDto';
import { getRequest, postRequest } from '../httpRequests';

export const useFetchMe = () =>
  useQuery({
    queryKey: QUERY_KEYS.AUTH.ME,
    queryFn: () => getRequest<UserDto>({ url: ENDPOINTS.AUTH.ME }),
    staleTime: 0,
    gcTime: 0,
  });

export const useLogIn = () => {
  const queryClient = useQueryClient();
  const navigation = useNavigate();

  return useMutation({
    mutationFn: (payload: LoginDto) =>
      postRequest({ url: ENDPOINTS.AUTH.LOGIN, body: payload }),
    onSuccess: async () => {
      toast.success('Successfully logged in');

      await queryClient.resetQueries();
      await navigation(ROUTES.HOME);
    },
  });
};

export const useLogOut = () => {
  const queryClient = useQueryClient();
  const navigation = useNavigate();

  return useMutation({
    mutationFn: () => postRequest({ url: ENDPOINTS.AUTH.LOGOUT }),
    onSuccess: async () => {
      toast.success('Successfully logged out');

      await queryClient.resetQueries();
      await navigation(ROUTES.HOME);
    },
  });
};

export const useSignUp = () => {
  const queryClient = useQueryClient();
  const navigation = useNavigate();

  return useMutation({
    mutationFn: (payload: SignupDto) =>
      postRequest({ url: ENDPOINTS.AUTH.SIGNUP, body: payload }),
    onSuccess: async () => {
      toast.success('Successfully signed up');

      await queryClient.resetQueries();
      await navigation(ROUTES.HOME);
    },
  });
};

export const useRefreshToken = () =>
  useMutation({
    mutationFn: () => postRequest({ url: ENDPOINTS.AUTH.REFRESH }),
  });
