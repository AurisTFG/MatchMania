import axios, { AxiosError } from 'axios';
import { getErrorFromAxiosError } from '../utils/axiosUtils';

const baseURL = import.meta.env.MATCHMANIA_API_BASE_URL as string;
if (!baseURL) {
  throw new Error('Env var MATCHMANIA_API_BASE_URL is not defined');
}

const axiosClient = axios.create({
  baseURL: baseURL,
  withCredentials: true,
});

const axiosClientWithInterceptors = axios.create({
  baseURL: baseURL,
  withCredentials: true,
});

const onSuccess = <T>(response: { data: T }) => {
  return response.data;
};

const onError = async (axiosError: AxiosError) => {
  if (axiosError.response?.status === 401) {
    try {
      await axiosClient.post('/auth/refresh', null);

      if (axiosError.config) {
        // Retry the original request
        return await axiosClient.request(axiosError.config);
      }

      return await Promise.reject(
        new Error('Request configuration is undefined'),
      );
    } catch (newAxiosError) {
      return Promise.reject(
        getErrorFromAxiosError(newAxiosError as AxiosError),
      );
    }
  }

  return Promise.reject(getErrorFromAxiosError(axiosError));
};

axiosClientWithInterceptors.interceptors.response.use(onSuccess, onError);

export default axiosClientWithInterceptors;
