import api from "./api";

export async function signup(
  username: string,
  email: string,
  password: string
) {
  try {
    const response = await api.post("/auth/signup", {
      username,
      email,
      password,
    });

    return response.data;
  } catch (error) {
    throw new Error(String(error));
  }
}

export async function login(email: string, password: string) {
  try {
    await api.post("/auth/login", { email, password });
  } catch (error) {
    throw new Error(String(error));
  }
}

export async function logout() {
  try {
    await api.post("/auth/logout", null);
  } catch (error) {
    console.error("Failed to logout:", error);
  }
}

export async function refreshToken() {
  try {
    await api.post("/auth/refresh", null);
  } catch (error) {
    console.error("Failed to refresh token:", error);
  }
}

export async function getMe() {
  try {
    const response = await api.get("/auth/me");
    return response.data.user;
  } catch (error) {
    console.error("Failed to get user data:", error);
    throw error;
  }
}
