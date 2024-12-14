import api from "./api";
import { getAccessToken } from "./auth";
import { jwtDecode } from "jwt-decode";

export async function getAllUsers() {
  try {
    const response = await api.get("/users");

    return response.data;
  } catch (error) {
    throw new Error(String(error));
  }
}

export async function getUser(userId: string) {
  try {
    const response = await api.get(`/users/${userId}`);

    return response.data;
  } catch (error) {
    throw new Error(String(error));
  }
}

export async function getCurrentUser() {
  try {
    const accessToken = getAccessToken();
    if (!accessToken) {
      return null;
    }

    const { sub } = jwtDecode(accessToken);
    if (!sub) {
      return null;
    }

    const response = await api.get(`/users/${sub}`);

    return response.data;
  } catch (error) {
    throw new Error(String(error));
  }
}

export async function updateUser(
  username: string,
  email: string,
  password: string
) {
  try {
    const response = await api.put("/users", {
      username,
      email,
      password,
    });

    return response.data;
  } catch (error) {
    throw new Error(String(error));
  }
}
