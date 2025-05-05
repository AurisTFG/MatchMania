import Brightness4Icon from '@mui/icons-material/Brightness4';
import Brightness7Icon from '@mui/icons-material/Brightness7';
import DirectionsCarIcon from '@mui/icons-material/DirectionsCar';
import {
  AppBar,
  Box,
  IconButton,
  Toolbar,
  Typography,
  alpha,
  useTheme,
} from '@mui/material';
import { useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { ROUTES } from 'constants/routes';
import withErrorBoundary from 'hocs/withErrorBoundary';

type HeaderProps = {
  toggleTheme: () => void;
  mode: 'light' | 'dark';
};

function Header({ toggleTheme, mode }: HeaderProps) {
  const theme = useTheme();
  const location = useLocation();

  const currentRoute = ROUTES.find((route) => route.path === location.pathname);
  const pageTitle = currentRoute?.label ?? '';

  useEffect(() => {
    document.title = pageTitle ? `${pageTitle} - MatchMania` : 'MatchMania';
  }, [pageTitle]);

  return (
    <AppBar
      position="fixed"
      elevation={0}
      sx={{
        background: alpha(theme.palette.background.paper, 0.9),
        backdropFilter: 'blur(10px)',
        color: theme.palette.text.primary,
        zIndex: (theme) => theme.zIndex.drawer + 1,
      }}
    >
      <Toolbar sx={{ position: 'relative' }}>
        <Typography
          variant="h6"
          noWrap
          component="div"
          sx={{
            fontWeight: 600,
            display: 'flex',
            alignItems: 'center',
            gap: 1,
          }}
        >
          <DirectionsCarIcon sx={{ fontSize: '1.5rem' }} />
          MatchMania
        </Typography>

        <Typography
          component="div"
          sx={{
            position: 'absolute',
            left: '50%',
            transform: 'translateX(-50%)',
            fontWeight: 700,
            fontSize: '1.25rem',
            letterSpacing: '0.5px',
            textTransform: 'uppercase',
            color: theme.palette.text.primary,
          }}
        >
          {pageTitle}
        </Typography>

        <Box sx={{ marginLeft: 'auto' }}>
          <IconButton
            color="inherit"
            onClick={toggleTheme}
          >
            {mode === 'dark' ? <Brightness7Icon /> : <Brightness4Icon />}
          </IconButton>
        </Box>
      </Toolbar>
    </AppBar>
  );
}

export default withErrorBoundary(Header);
