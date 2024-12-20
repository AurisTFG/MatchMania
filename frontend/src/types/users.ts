export interface User {
  id: string;
  username: string;
  email: string;
  role: string;
}

export interface UsersResponse {
  users?: User[];
}

export interface UserResponse {
  user?: User;
}
