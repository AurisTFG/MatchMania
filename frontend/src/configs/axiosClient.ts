import axios, { AxiosError } from 'axios';

const axiosClient = axios.create({
  baseURL: import.meta.env.MATCHMANIA_API_BASE_URL as string,
  withCredentials: true,
});

const onFulfilled = <T>(response: { data: T }) => {
  return response.data;
};

const onRejected = async (error: AxiosError) => {
  if (error.response?.status === 401) {
    console.log('Refreshing token...');

    await axiosClient.post('/auth/refresh', null);

    if (error.config) {
      return axios.request(error.config);
    }
  }

  return Promise.reject(
    error instanceof Error ? error : new Error(String(error)),
  );
};

axiosClient.interceptors.response.use(onFulfilled, onRejected);

export default axiosClient;
