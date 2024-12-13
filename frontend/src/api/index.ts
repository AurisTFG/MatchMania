import axios from "axios";

const api = axios.create({
  baseURL: "http://matchmania-api-ripky.ondigitalocean.app/api/v1",
  headers: {
    "Content-Type": "application/json",
  },
});

export default api;
