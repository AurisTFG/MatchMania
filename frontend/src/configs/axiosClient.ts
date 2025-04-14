import axios from 'axios';

const baseURL = import.meta.env.MATCHMANIA_API_BASE_URL as string;
if (!baseURL) {
  throw new Error('Env var MATCHMANIA_API_BASE_URL is not defined');
}

const axiosClient = axios.create({
  baseURL: baseURL,
  withCredentials: true,
});

export default axiosClient;
