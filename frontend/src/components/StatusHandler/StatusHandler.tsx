import { Typography } from 'antd';
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
        <Typography.Title level={4}>Loading...</Typography.Title>
      </div>
    );
  }

  if (error) {
    return (
      <div style={{ padding: 20 }}>
        <Typography.Title level={4}>
          {`${errorMessage}: ${error.message}`}
        </Typography.Title>
      </div>
    );
  }

  if (isEmpty) {
    return (
      <div style={{ padding: 20 }}>
        <Typography.Title level={4}>{emptyMessage}</Typography.Title>
      </div>
    );
  }

  return <>{children}</>;
}
