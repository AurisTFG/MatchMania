import { ReactNode, createContext, useContext, useMemo } from 'react';
import { useFetchMe } from 'api/hooks/authHooks';
import { UserDto } from 'types/dtos/responses/users/userDto';

type AuthContextType = {
  user: UserDto | null;
  isLoggedIn: boolean;
  isAuthLoading: boolean;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function useAuth() {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error('useAuth must be used within the AuthProvider component');
  }

  return context;
}

export function AuthProvider({ children }: { children: ReactNode }) {
  const { data: fetchMeQuery, isLoading: isAuthLoading } = useFetchMe();

  const user = fetchMeQuery ?? null;
  const isLoggedIn = !!user?.id;

  const authContextValue = useMemo(
    () => ({ user, isLoggedIn, isAuthLoading }),
    [user, isLoggedIn, isAuthLoading],
  );

  return (
    <AuthContext.Provider value={authContextValue}>
      {children}
    </AuthContext.Provider>
  );
}
