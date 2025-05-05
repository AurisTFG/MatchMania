import { Box, CircularProgress } from '@mui/material';
import { Navigate } from 'react-router-dom';
import ROUTE_PATHS from 'constants/route_paths';
import { useAuth } from 'providers/AuthProvider/AuthProvider';
import { Permission } from 'types/enums/permission';
import withErrorBoundary from './withErrorBoundary';

type WithAuthProps = {
  permission: Permission;
  dataOwnerUserId?: (props: unknown) => string;
  redirect?: boolean;
};

export default function withAuth<P extends object>(
  WrappedComponent: React.ComponentType<P>,
  { permission, dataOwnerUserId, redirect = true }: WithAuthProps,
) {
  const ComponentWithErrorBoundary = withErrorBoundary(WrappedComponent);

  return function AuthWrapper(props: P) {
    const { user, isLoggedIn, isAuthLoading, isRefreshingToken } = useAuth();
    const ownerUserId = dataOwnerUserId ? dataOwnerUserId(props) : null;

    if (isAuthLoading || isRefreshingToken) {
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

    if (permission === Permission.LoggedOut && isLoggedIn) {
      if (redirect) {
        return <Navigate to={ROUTE_PATHS.HOME} />;
      }

      return null;
    }

    if (!isLoggedIn) {
      if (redirect) {
        return <Navigate to={ROUTE_PATHS.UNAUTHORIZED} />;
      }

      return null;
    }

    if (permission === Permission.LoggedIn) {
      return <ComponentWithErrorBoundary {...props} />;
    }

    if (user?.permissions.includes(Permission.Admin)) {
      return <ComponentWithErrorBoundary {...props} />;
    }

    if (
      !user?.permissions.includes(permission) ||
      (ownerUserId && user.id !== ownerUserId)
    ) {
      if (redirect) {
        return <Navigate to={ROUTE_PATHS.FORBIDDEN} />;
      }

      return null;
    }

    return <ComponentWithErrorBoundary {...props} />;
  };
}
