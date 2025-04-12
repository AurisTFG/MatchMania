import axios, { AxiosError } from 'axios';
import { ErrorDto } from '../types';

const axiosClient = axios.create({
  baseURL: import.meta.env.MATCHMANIA_API_BASE_URL as string,
  withCredentials: true,
});

const onSuccess = <T>(response: { data: T }) => {
  return response.data;
};

const onError = async (error: AxiosError) => {
  if (error.response?.status === 401) {
    console.log('Refreshing token...');

    await axiosClient.post('/auth/refresh', null);

    if (error.config) {
      return axios.request(error.config);
    }
  }

  const errorDto = error?.response?.data as ErrorDto;

  if (errorDto && errorDto.error) {
    return Promise.reject(new Error(errorDto.error));
  } else {
    return Promise.reject(new Error('An unknown error occurred.'));
  }
};

axiosClient.interceptors.response.use(onSuccess, onError);

export default axiosClient;
