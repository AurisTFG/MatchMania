import { Box, CircularProgress } from '@mui/material';
import { Navigate } from 'react-router-dom';
import { Permission } from 'constants/permissions';
import { ROUTE_PATHS } from 'constants/route_paths';
import { useAuth } from 'providers/AuthProvider/AuthProvider';

type WithAuthProps = {
  requiredPermission?: Permission;
  isLoggedInOnly?: boolean;
};

export default function withAuth<P extends object>(
  WrappedComponent: React.ComponentType<P>,
  { requiredPermission, isLoggedInOnly = false }: WithAuthProps,
) {
  return function AuthWrapper(props: P) {
    const { user, isLoggedIn, isAuthLoading } = useAuth();

    if (isAuthLoading) {
      return (
        <Box
          display="flex"
          justifyContent="center"
          alignItems="center"
          minHeight="100vh"
        >
          <CircularProgress />
        </Box>
      );
    }

    if (!isLoggedIn) {
      return <Navigate to={ROUTE_PATHS.UNAUTHORIZED} />;
    }

    if (isLoggedInOnly) {
      return <WrappedComponent {...props} />;
    }

    const hasPermission =
      !requiredPermission || user?.permissions.includes(requiredPermission);

    if (!hasPermission) {
      return <Navigate to={ROUTE_PATHS.FORBIDDEN} />;
    }

    return <WrappedComponent {...props} />;
  };
}
