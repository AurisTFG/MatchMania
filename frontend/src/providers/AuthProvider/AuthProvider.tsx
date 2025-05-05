import { ReactNode, createContext, useContext, useMemo, useState } from 'react';
import { useFetchMe } from 'api/hooks/authHooks';
import { UserDto } from 'types/dtos/responses/users/userDto';

type AuthContextType = {
  user: UserDto | null;
  isLoggedIn: boolean;
  isAuthLoading: boolean;
  isRefreshingToken: boolean;
  setIsRefreshingToken: (isRefreshing: boolean) => void;
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
  const [isRefreshingToken, setIsRefreshingToken] = useState(false);

  const {
    data: fetchMeQuery,
    isLoading: isAuthLoading,
    error: fetchMeError,
  } = useFetchMe();

  const user = fetchMeQuery ?? null;
  const isLoggedIn = !!user?.id && !fetchMeError;

  const authContextValue = useMemo(
    () => ({
      user,
      isLoggedIn,
      isAuthLoading,
      isRefreshingToken,
      setIsRefreshingToken,
    }),
    [user, isLoggedIn, isAuthLoading, isRefreshingToken, setIsRefreshingToken],
  );

  return (
    <AuthContext.Provider value={authContextValue}>
      {children}
    </AuthContext.Provider>
  );
}
