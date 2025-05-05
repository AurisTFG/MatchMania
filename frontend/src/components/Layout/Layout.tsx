import { Box, useTheme } from '@mui/material';
import { useEffect } from 'react';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAuth } from 'providers/AuthProvider';
import { initializeSetIsRefreshingToken } from 'utils/queryClientUtils';
import Header from './Header';
import Sidebar from './Sidebar';

type LayoutProps = {
  children: React.ReactNode;
  toggleTheme: () => void;
  mode: 'light' | 'dark';
};

function Layout({ children, toggleTheme, mode }: LayoutProps) {
  const { setIsRefreshingToken } = useAuth();
  const theme = useTheme();

  useEffect(() => {
    initializeSetIsRefreshingToken(setIsRefreshingToken);
  }, [setIsRefreshingToken]);

  return (
    <Box sx={{ display: 'flex', background: theme.palette.background.default }}>
      <Header
        toggleTheme={toggleTheme}
        mode={mode}
      />

      <Sidebar />
      <Box
        component="main"
        sx={{ flexGrow: 1, minHeight: '100vh', p: 3, pt: 11 }}
      >
        {children}
      </Box>
    </Box>
  );
}

export default withErrorBoundary(Layout);
