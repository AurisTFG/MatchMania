import axios from "axios";
import { getAccessToken, refreshToken } from "./auth";

const api = axios.create({
  baseURL: import.meta.env.MATCHMANIA_API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

api.interceptors.request.use((config) => {
  const token = getAccessToken();

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      console.log("Refreshing token...");

      await refreshToken();

      error.config.headers.Authorization = `Bearer ${getAccessToken()}`;

      return axios.request(error.config); // Retry the original request
    }

    return Promise.reject(error);
  }
);

export default api;
