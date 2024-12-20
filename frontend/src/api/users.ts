import api from "./api";
import { handleApiError } from "./utils";
import { getAccessToken } from "./auth";
import { jwtDecode } from "jwt-decode";
import { User, UserResponse, UsersResponse } from "../types/users";

export async function getAllUsers(): Promise<User[]> {
  try {
    const { data } = await api.get<UsersResponse>("/users");

    return data.users || [];
  } catch (error) {
    handleApiError(error);
  }
}

export async function getUser(userId: string): Promise<User> {
  try {
    const { data } = await api.get<UserResponse>(`/users/${userId}`);

    return data.user || ({} as User);
  } catch (error) {
    handleApiError(error);
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
