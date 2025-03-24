import axios from "axios";
import { refreshToken } from "./auth";

const api = axios.create({
  baseURL: import.meta.env.MATCHMANIA_API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

api.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      console.log("Refreshing token...");

      await refreshToken();

      return axios.request(error.config);
    }

    return Promise.reject(error);
  }
);

export default api;
