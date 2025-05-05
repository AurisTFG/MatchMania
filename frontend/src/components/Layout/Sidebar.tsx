import {
  Box,
  Divider,
  Drawer,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Stack,
  Toolbar,
  Typography,
  useTheme,
} from '@mui/material';
import { Link, useLocation } from 'react-router-dom';
import UserAvatar from 'components/UserAvatar';
import { ROUTES } from 'constants/routes';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAuth } from 'providers/AuthProvider';
import { Permission } from 'types/enums/permission';
import OptionsMenu from './OptionsMenu';

function Sidebar() {
  const theme = useTheme();
  const { user, isLoggedIn } = useAuth();
  const location = useLocation();

  const sidebarRoutes = ROUTES.filter(
    (route) =>
      route.icon &&
      (!route.permission ||
        (user?.permissions ?? []).includes(route.permission) ||
        (isLoggedIn && route.permission === Permission.LoggedIn)),
  );

  return (
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
                color:
                  location.pathname === path
                    ? theme.palette.primary.main
                    : theme.palette.text.primary,
                backgroundColor:
                  location.pathname === path
                    ? theme.palette.mode === 'dark'
                      ? 'rgba(99,102,241,0.15)'
                      : 'rgba(79,70,229,0.08)'
                    : 'transparent',
                fontWeight: location.pathname === path ? 600 : 400,
                '&:hover': {
                  backgroundColor:
                    theme.palette.mode === 'dark'
                      ? 'rgba(99,102,241,0.1)'
                      : 'rgba(79,70,229,0.05)',
                },
              }}
            >
              <ListItemIcon
                sx={{
                  color:
                    location.pathname === path
                      ? theme.palette.primary.main
                      : theme.palette.text.secondary,
                }}
              >
                {icon}
              </ListItemIcon>
              <ListItemText
                primary={label}
                slotProps={{
                  primary: {
                    sx: {
                      fontWeight: location.pathname === path ? 600 : 500,
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
          imageUrl={user?.profilePhotoUrl}
          name={user?.username}
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
  );
}

export default withErrorBoundary(Sidebar);
