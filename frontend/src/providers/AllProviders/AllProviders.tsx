import { CssBaseline } from '@mui/material';
import { ThemeProvider } from '@mui/material/styles';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { ReactNode, useState } from 'react';
import { BrowserRouter } from 'react-router-dom';
import { Toaster } from 'sonner';
import Layout from 'components/Layout';
import { queryClient } from 'configs/queryClient';
import { getTheme } from 'styles/theme';
import { AuthProvider } from '../AuthProvider';

export default function AllProviders({ children }: { children: ReactNode }) {
  const [mode, setMode] = useState<'light' | 'dark'>('dark');

  return (
    <ThemeProvider theme={getTheme(mode)}>
      <CssBaseline enableColorScheme />
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Toaster
              richColors
              visibleToasts={4}
              position="bottom-right"
            />
            <BrowserRouter>
              <Layout
                toggleTheme={() => {
                  setMode(mode === 'light' ? 'dark' : 'light');
                }}
                mode={mode}
              >
                {children}
              </Layout>
            </BrowserRouter>
          </LocalizationProvider>
        </AuthProvider>
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </ThemeProvider>
  );
}
