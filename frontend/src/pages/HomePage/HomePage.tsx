import { Box, Button, Stack, Typography, useTheme } from '@mui/material';
import { Link } from 'react-router-dom';
import backgroundImage from 'assets/background.webp';
import { ROUTE_PATHS } from 'constants/route_paths';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAuth } from 'providers/AuthProvider';

function HomePage() {
  const { isLoggedIn } = useAuth();
  const theme = useTheme();

  return (
    <Box
      sx={{
        position: 'relative',
        textAlign: 'center',
        height: 'calc(100vh - 112px)',
        overflow: 'hidden',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
      }}
    >
      <Box
        sx={{
          position: 'absolute',
          inset: 0,
          width: '100%',
          height: '100%',
          backgroundImage: `url(${backgroundImage})`,
          backgroundSize: 'cover',
          backgroundPosition: 'center',
          backgroundRepeat: 'no-repeat',
          zIndex: 0,
          '&::after': {
            content: '""',
            position: 'absolute',
            inset: 0,
            backgroundColor: 'rgba(0,0,0,0.7)',
          },
        }}
      />

      <Box
        sx={{
          position: 'relative',
          zIndex: 1,
          maxWidth: '90%',
          textAlign: 'center',
        }}
      >
        <Typography
          variant="h2"
          component="h1"
          sx={{
            color: theme.palette.primary.main,
            textShadow: '2px 2px 10px rgba(0,0,0,0.8)',
            fontWeight: 700,
            mb: 4,
            wordBreak: 'break-word',
            maxWidth: '600px',
            mx: 'auto',
          }}
        >
          Welcome to MatchMania!
        </Typography>

        {!isLoggedIn && (
          <Stack
            direction="row"
            spacing={2}
            justifyContent="center"
          >
            <Button
              component={Link}
              to={ROUTE_PATHS.LOGIN}
              variant="contained"
              size="large"
              sx={{
                px: 4,
                py: 1.5,
                borderRadius: theme.shape.borderRadius,
                textTransform: 'none',
                fontWeight: 600,
                fontSize: '1rem',
                boxShadow:
                  theme.palette.mode === 'dark'
                    ? '0 2px 8px rgba(0,0,0,0.5)'
                    : '0 2px 8px rgba(0,0,0,0.1)',
              }}
            >
              Log In
            </Button>
            <Button
              component={Link}
              to={ROUTE_PATHS.SIGNUP}
              variant="outlined"
              size="large"
              sx={{
                px: 4,
                py: 1.5,
                borderRadius: theme.shape.borderRadius,
                textTransform: 'none',
                fontWeight: 600,
                fontSize: '1rem',
                color: theme.palette.primary.main,
                borderColor: theme.palette.primary.main,
                '&:hover': {
                  borderColor: theme.palette.primary.dark,
                  backgroundColor:
                    theme.palette.mode === 'dark'
                      ? 'rgba(255,255,255,0.05)'
                      : theme.palette.action.hover,
                },
              }}
            >
              Sign Up
            </Button>
          </Stack>
        )}
      </Box>
    </Box>
  );
}

export default withErrorBoundary(HomePage);
