import { createTheme } from '@mui/material/styles';

export const getTheme = (mode: 'light' | 'dark') =>
  createTheme({
    palette: {
      mode,
      primary: { main: '#6366f1' },
      secondary: { main: '#4f46e5' },
      error: { main: '#ef4444' },
      warning: { main: '#f59e0b' },
      info: { main: '#3b82f6' },
      success: { main: '#22c55e' },
      background: {
        default: mode === 'dark' ? '#0e0e10' : '#f4f6fa',
        paper: mode === 'dark' ? '#1e1e22' : '#ffffff',
      },
      text: {
        primary: mode === 'dark' ? '#ffffffde' : '#1e1e2f',
        secondary: mode === 'dark' ? '#bbbbbb' : '#4f4f5a',
      },
    },
    shape: { borderRadius: 16 },
    typography: {
      fontFamily: 'Inter, Roboto, sans-serif',
      button: {
        fontWeight: 600,
        letterSpacing: '0.5px',
      },
      h6: {
        fontWeight: 700,
      },
      subtitle1: {
        fontWeight: 500,
        fontSize: '1rem',
      },
    },
    components: {
      MuiDrawer: {
        styleOverrides: {
          paper: {
            boxShadow: '2px 0 8px rgba(0,0,0,0.05)',
          },
        },
      },
      MuiPaper: {
        styleOverrides: {
          root: ({ theme }) => ({
            ...(theme.palette.mode === 'dark' && {
              backgroundColor: 'rgba(30,30,34,0.8)',
              backdropFilter: 'blur(6px)',
            }),
            backgroundImage: 'none',
            transition: 'box-shadow 0.3s ease, background-color 0.3s ease',
            boxShadow:
              theme.palette.mode === 'dark'
                ? '0 2px 8px rgba(0,0,0,0.3)'
                : '0 2px 8px rgba(0,0,0,0.05)',
          }),
        },
      },
      MuiMenuItem: {
        styleOverrides: {
          root: ({ theme }) => ({
            borderRadius: theme.shape.borderRadius,
            '&:hover': {
              backgroundColor:
                theme.palette.mode === 'dark'
                  ? 'rgba(99,102,241,0.1)'
                  : 'rgba(79,70,229,0.05)',
            },
            margin: '2px 4px',
          }),
        },
      },
      MuiButton: {
        styleOverrides: {
          root: ({ theme }) => ({
            borderRadius: theme.shape.borderRadius,
            textTransform: 'none',
            fontWeight: 600,
            transition: 'all 0.2s ease',
            boxShadow:
              theme.palette.mode === 'dark'
                ? '0 1px 3px rgba(0,0,0,0.4)'
                : '0 1px 3px rgba(0,0,0,0.1)',
            '&:hover': {
              boxShadow:
                theme.palette.mode === 'dark'
                  ? '0 4px 10px rgba(0,0,0,0.5)'
                  : '0 4px 10px rgba(0,0,0,0.15)',
            },
          }),
        },
      },
      MuiOutlinedInput: {
        styleOverrides: {
          root: ({ theme }) => ({
            borderRadius: theme.shape.borderRadius,
            '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
              borderColor: theme.palette.primary.main,
              boxShadow: `0 0 0 2px ${theme.palette.primary.main}33`,
            },
          }),
          notchedOutline: {
            borderColor: 'rgba(0,0,0,0.23)',
          },
          input: {
            '&:-webkit-autofill': {
              WebkitBoxShadow: `0 0 0 100px rgba(38, 103, 152, 0.25) inset`,
            },
          },

          inputAdornedStart: {
            paddingLeft: 0,
          },

          inputAdornedEnd: {
            paddingRight: 0,
          },

          adornedStart: {
            paddingLeft: 0,
          },
        },
      },
    },
  });
