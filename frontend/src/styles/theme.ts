import { createTheme } from '@mui/material/styles';

export const getTheme = (mode: 'light' | 'dark') =>
  createTheme({
    palette: {
      mode,
      primary: { main: '#00C37A' },
      secondary: { main: '#00A86B' },
      error: { main: '#ff4d4f' },
      warning: { main: '#ffc107' },
      info: { main: '#00b0ff' },
      success: { main: '#4caf50' },
      background: {
        default: mode === 'dark' ? '#0a0f0d' : '#e6f9f3',
        paper: mode === 'dark' ? '#101918' : '#ffffff',
      },
      text: {
        primary: mode === 'dark' ? '#ffffffde' : '#0a0f0d',
        secondary: mode === 'dark' ? '#bbbbbb' : '#4f4f5a',
      },
    },
    shape: { borderRadius: 12 },
    typography: {
      fontFamily: 'Inter, Roboto, sans-serif',
      button: {
        fontWeight: 700,
        letterSpacing: '0.75px',
      },
      h6: {
        fontWeight: 800,
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
            boxShadow: '2px 0 8px rgba(0,0,0,0.25)',
          },
        },
      },
      MuiPaper: {
        styleOverrides: {
          root: ({ theme }) => ({
            ...(theme.palette.mode === 'dark' && {
              backgroundColor: 'rgba(16,25,24,0.85)',
              backdropFilter: 'blur(8px)',
            }),
            backgroundImage: 'none',
            transition: 'box-shadow 0.3s ease, background-color 0.3s ease',
            boxShadow:
              theme.palette.mode === 'dark'
                ? '0 2px 8px rgba(0,0,0,0.5)'
                : '0 2px 8px rgba(0,0,0,0.1)',
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
                  ? 'rgba(0,195,122,0.1)'
                  : 'rgba(0,168,107,0.05)',
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
            fontWeight: 700,
            transition: 'all 0.2s ease',
            boxShadow:
              theme.palette.mode === 'dark'
                ? '0 1px 4px rgba(0,0,0,0.6)'
                : '0 1px 3px rgba(0,0,0,0.2)',
            '&:hover': {
              boxShadow:
                theme.palette.mode === 'dark'
                  ? '0 4px 12px rgba(0,0,0,0.7)'
                  : '0 4px 10px rgba(0,0,0,0.2)',
              backgroundColor: theme.palette.primary.main,
              color: '#ffffff',
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
            },
          }),
          notchedOutline: ({ theme }) => ({
            borderColor:
              theme.palette.mode === 'dark'
                ? 'rgba(255,255,255,0.5)'
                : 'rgba(0,0,0,0.23)',
          }),
          input: {
            '&:-webkit-autofill': {
              WebkitBoxShadow: `0 0 0 100px rgba(0, 195, 122, 0.1) inset`,
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
