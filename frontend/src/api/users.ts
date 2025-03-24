import api from "./api";
import { handleApiError } from "./utils";
import { User, UsersResponse } from "../types/users";

export async function getAllUsers(): Promise<User[]> {
  try {
    const { data } = await api.get<UsersResponse>("/users", {
      withCredentials: true,
    });

    return data.users || [];
  } catch (error) {
    handleApiError(error);
    return [];
  }
}
