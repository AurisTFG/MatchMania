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

    const accessToken = response.data.accessToken;
    saveAccessToken(accessToken);

    return response.data;
  } catch (error) {
    throw new Error(String(error));
  }
}

export async function login(email: string, password: string) {
  try {
    const response = await api.post("/auth/login", { email, password });

    const accessToken = response.data.accessToken;
    saveAccessToken(accessToken);

    return response.data;
  } catch (error) {
    throw new Error(String(error));
  }
}

export async function logout() {
  try {
    await api.post("/auth/logout", null);

    removeAccessToken();
  } catch (error) {
    console.error("Failed to logout:", error);
  }
}

export async function refreshToken() {
  try {
    const response = await api.post("/auth/refresh-token", null);
    // const response = await api.post("/auth/refresh-token", null, {
    //   withCredentials: true,
    // });

    const newAccessToken = response.data.accessToken;

    saveAccessToken(newAccessToken);
  } catch (error) {
    console.error("Failed to refresh token:", error);
  }
}

export function saveAccessToken(accessToken: string) {
  localStorage.setItem("accessToken", accessToken);
}

export function getAccessToken() {
  return localStorage.getItem("accessToken");
}

export function removeAccessToken() {
  localStorage.removeItem("accessToken");
}
