import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { ENDPOINTS } from "../../constants/endpoints";
import { QUERY_KEYS } from "../../constants/queryKeys";
import { getRequest, postRequest } from "../httpRequests";

export const useFetchMe = () =>
  useQuery({
    queryKey: QUERY_KEYS.AUTH.ME,
    queryFn: () =>
      getRequest({
        url: ENDPOINTS.AUTH.ME,
      }),
  });

export const useLogIn = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: { email: string; password: string }) =>
      postRequest({
        url: ENDPOINTS.AUTH.LOGIN,
        body: data,
      }),
    onSuccess: async () => {
      await queryClient.resetQueries({ queryKey: QUERY_KEYS.AUTH.ME });
    },
  });
};

export const useLogOut = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: () => postRequest({ url: ENDPOINTS.AUTH.LOGOUT }),
    onSuccess: async () => {
      await queryClient.resetQueries({ queryKey: QUERY_KEYS.AUTH.ME });
    },
  });
};

export const useSignUp = () => {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (data: { username: string; email: string; password: string }) =>
      postRequest({
        url: ENDPOINTS.AUTH.SIGNUP,
        body: data,
      }),
    onSuccess: async () => {
      await queryClient.resetQueries({ queryKey: QUERY_KEYS.AUTH.ME });
    },
  });
};

export const useRefreshToken = () =>
  useMutation({
    mutationFn: () => postRequest({ url: ENDPOINTS.AUTH.REFRESH }),
  });
