import axios from "axios";

const axiosClient = axios.create({
  baseURL: import.meta.env.MATCHMANIA_API_BASE_URL as string,
  withCredentials: true,
});

axiosClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    if (error.response?.status === 401) {
      console.log("Refreshing token...");

      await axiosClient.post("/auth/refresh", null);

      return axios.request(error.config);
    }

    return Promise.reject(
      error instanceof Error ? error : new Error(String(error)),
    );
  },
);

export default axiosClient;
