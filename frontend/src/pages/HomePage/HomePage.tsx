import { Box, Button, Typography } from '@mui/material';
import { Link } from 'react-router-dom';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAuth } from 'providers/AuthProvider';

function HomePage() {
  const { user } = useAuth();

  return (
    <Box
      sx={{
        textAlign: 'center',
        height: '83.5vh',
        backgroundImage: 'url("https://i.imgur.com/2k81qri.jpeg")',
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        backgroundRepeat: 'no-repeat',
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
      }}
    >
      <Typography
        variant="h3"
        gutterBottom
        sx={{
          wordWrap: 'break-word',
          overflowWrap: 'break-word',
          width: '100%',
          color: 'rgba(94, 140, 192, 0.9)',
          fontSize: '6rem',
          textShadow: '5px 5px 5px rgba(0, 0, 0, 0.8)',
        }}
      >
        Welcome to MatchMania,
        <br />
        {user ? user.username : 'Guest'}!
      </Typography>

      {!user && (
        <Box sx={{ mt: 3 }}>
          <Button
            component={Link}
            to="/login"
            variant="contained"
            color="primary"
            sx={{
              mx: 1,
              padding: '12px 24px',
              borderRadius: '8px',
              fontSize: '16px',
              boxShadow: 3,
              '&:hover': {
                boxShadow: 6,
              },
            }}
          >
            Login
          </Button>

          <Button
            component={Link}
            to="/signup"
            variant="contained"
            color="secondary"
            sx={{
              mx: 1,
              padding: '12px 24px',
              borderRadius: '8px',
              fontSize: '16px',
              boxShadow: 3,
              '&:hover': {
                boxShadow: 6,
              },
            }}
          >
            Sign Up
          </Button>
        </Box>
      )}
    </Box>
  );
}

export default withErrorBoundary(HomePage);
