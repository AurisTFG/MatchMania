import Brightness4Icon from '@mui/icons-material/Brightness4';
import Brightness7Icon from '@mui/icons-material/Brightness7';
import DirectionsCarIcon from '@mui/icons-material/DirectionsCar';
import {
  AppBar,
  Box,
  Divider,
  Drawer,
  IconButton,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Stack,
  Toolbar,
  Typography,
  alpha,
  useTheme,
} from '@mui/material';
import { Link } from 'react-router-dom';
import { UserAvatar } from 'components/UserAvatar';
import { ROUTES } from 'constants/routes';
import { useAuth } from 'providers/AuthProvider';
import OptionsMenu from './OptionsMenu';

export default function Layout({
  children,
  toggleTheme,
  mode,
}: {
  children: React.ReactNode;
  toggleTheme: () => void;
  mode: 'light' | 'dark';
}) {
  const theme = useTheme();
  const { user } = useAuth();

  const sidebarRoutes = ROUTES.filter((route) => route.icon);

  return (
    <Box sx={{ display: 'flex', background: theme.palette.background.default }}>
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
        <Toolbar>
          <Typography
            variant="h6"
            noWrap
            component="div"
            sx={{
              flexGrow: 1,
              fontWeight: 600,
              display: 'flex',
              alignItems: 'center',
              gap: 1,
            }}
          >
            <DirectionsCarIcon sx={{ fontSize: '1.5rem' }} />
            MatchMania
          </Typography>
          <IconButton
            color="inherit"
            onClick={toggleTheme}
          >
            {mode === 'dark' ? <Brightness7Icon /> : <Brightness4Icon />}
          </IconButton>
        </Toolbar>
      </AppBar>

      <Drawer
        variant="permanent"
        sx={{
          width: 260,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: 260,
            boxSizing: 'border-box',
            backgroundColor: theme.palette.background.paper,
            borderRight: 'none',
            overflowX: 'hidden',
          },
        }}
      >
        <Box sx={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
          <Toolbar />

          <Divider sx={{ mb: 2 }} />

          <List sx={{ px: 1.5 }}>
            {sidebarRoutes.map(({ label, path, icon }, idx) => (
              <ListItem
                key={idx}
                component={Link}
                to={path}
                sx={{
                  borderRadius: 2,
                  mb: 0.5,
                  '&:hover': {
                    backgroundColor:
                      theme.palette.mode === 'dark'
                        ? 'rgba(99,102,241,0.1)'
                        : 'rgba(79,70,229,0.05)',
                  },
                }}
              >
                <ListItemIcon sx={{ color: theme.palette.text.secondary }}>
                  {icon}
                </ListItemIcon>
                <ListItemText
                  primary={label}
                  slotProps={{
                    primary: {
                      sx: {
                        fontWeight: 500,
                      },
                    },
                  }}
                />
              </ListItem>
            ))}
          </List>
        </Box>
        <Stack
          direction="row"
          sx={{
            p: 2,
            gap: 1,
            alignItems: 'center',
            borderTop: '1px solid',
            borderColor: 'divider',
          }}
        >
          <UserAvatar
            profilePhotoUrl={user?.profilePhotoUrl}
            username={user?.username}
            size={36}
          />
          <Box sx={{ mr: 'auto' }}>
            <Typography
              variant="body2"
              sx={{ fontWeight: 500, lineHeight: '16px' }}
            >
              {user?.username ?? 'Guest'}
            </Typography>
            <Typography
              variant="caption"
              sx={{ color: 'text.secondary' }}
            >
              {user?.email ?? 'Not logged in'}
            </Typography>
          </Box>

          <OptionsMenu />
        </Stack>
      </Drawer>

      <Box
        component="main"
        sx={{ flexGrow: 1, minHeight: '100vh', p: 3, pt: 11 }}
      >
        {children}
      </Box>
    </Box>
  );
}
