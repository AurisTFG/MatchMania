import { Typography } from '@mui/material';
import { ReactNode } from 'react';

type StatusHandlerProps = {
  isLoading: boolean;
  error?: Error | null;
  errorMessage?: string;
  isEmpty: boolean;
  emptyMessage?: string;
  children: ReactNode;
};

export default function StatusHandler({
  isLoading,
  error = null,
  errorMessage = 'Error occurred',
  isEmpty,
  emptyMessage = 'No data found',
  children,
}: StatusHandlerProps) {
  if (isLoading) {
    return (
      <div style={{ padding: 20 }}>
        <Typography
          variant="h5"
          component="div"
        >
          Loading...
        </Typography>
      </div>
    );
  }

  if (error) {
    return (
      <div style={{ padding: 20 }}>
        <Typography
          variant="h5"
          component="div"
          color="error"
        >
          {`${errorMessage}: ${error.message}`}
        </Typography>
      </div>
    );
  }

  if (isEmpty) {
    return (
      <div style={{ padding: 20 }}>
        <Typography
          variant="h5"
          component="div"
        >
          {emptyMessage}
        </Typography>
      </div>
    );
  }

  return <>{children}</>;
}
