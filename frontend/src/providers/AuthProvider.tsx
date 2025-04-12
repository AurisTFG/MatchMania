import { ReactNode, createContext, useContext } from 'react';
import { useFetchMe } from '../api/hooks/authHooks';
import { User } from '../types/userTypes';

type AuthContextType = {
  user: User | null;
  isLoggedIn: boolean;
  isLoading: boolean;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const { data: fetchMeQuery, isLoading } = useFetchMe();

  const user = fetchMeQuery as User | null;
  const isLoggedIn = !!user?.id;

  return (
    <AuthContext.Provider value={{ user, isLoggedIn, isLoading }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);

  if (!context) {
    throw new Error('useAuth must be used within the AuthProvider component');
  }

  return context;
};
