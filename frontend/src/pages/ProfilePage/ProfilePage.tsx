import {
  AccountCircle as AccountCircleIcon,
  Edit as EditIcon,
  Mail as MailIcon,
  SportsEsports as SportsEsportsIcon,
} from '@mui/icons-material';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CircularProgress,
  Divider,
  Stack,
  Typography,
} from '@mui/material';
import { useState } from 'react';
import { useFetchMe } from '../../api/hooks/authHooks';
import { useGetTrackmaniaOAuthUrl } from '../../api/hooks/trackmaniaOauthHooks';

export default function ProfilePage() {
  const { data: user, isLoading, error, refetch } = useFetchMe();
  const [connecting, setConnecting] = useState(false);
  const { data: urlDto, isLoading: isUrlLoading } = useGetTrackmaniaOAuthUrl();

  const handleConnectTrackmania = () => {
    setConnecting(true);
    window.location.href = urlDto?.url ?? '';
    setConnecting(false);
  };

  if (isLoading || isUrlLoading) {
    return (
      <Box
        display="flex"
        justifyContent="center"
        alignItems="center"
        minHeight="60vh"
      >
        <CircularProgress size={48} />
      </Box>
    );
  }

  if (error) {
    return (
      <Card
        sx={{ maxWidth: 400, mx: 'auto', mt: 8, textAlign: 'center', p: 3 }}
      >
        <Typography
          color="error"
          variant="h6"
          gutterBottom
        >
          Failed to load profile.
        </Typography>
        <Button
          onClick={() => void refetch()}
          variant="contained"
          sx={{ mt: 2 }}
        >
          Retry
        </Button>
      </Card>
    );
  }

  return (
    <Box
      display="flex"
      justifyContent="center"
      alignItems="center"
      minHeight="80vh"
      bgcolor="#f5f7fa"
    >
      <Card
        sx={{
          width: 420,
          borderRadius: 4,
          boxShadow: 6,
          p: 4,
          bgcolor: 'background.paper',
        }}
      >
        <Stack
          direction="column"
          alignItems="center"
          spacing={2}
        >
          <Avatar sx={{ width: 80, height: 80, bgcolor: 'primary.main' }}>
            <AccountCircleIcon sx={{ fontSize: 60 }} />
          </Avatar>
          <Typography
            variant="h5"
            fontWeight={600}
          >
            {user?.username}
          </Typography>
          {user?.role && (
            <Typography
              variant="body2"
              sx={{
                color: 'primary.main',
                bgcolor: 'primary.50',
                px: 2,
                py: 0.5,
                borderRadius: 2,
                fontWeight: 500,
                letterSpacing: 0.8,
                textTransform: 'uppercase',
                mt: -0.5,
                mb: 1,
              }}
            >
              {user.role}
            </Typography>
          )}
          <Stack
            direction="row"
            alignItems="center"
            spacing={1}
          >
            <MailIcon color="action" />
            <Typography
              variant="body2"
              color="text.secondary"
            >
              {user?.email}
            </Typography>
          </Stack>
          {user?.trackmaniaName && (
            <Typography
              variant="body2"
              color="text.secondary"
            >
              <strong>Trackmania Name:</strong> {user.trackmaniaName}
            </Typography>
          )}
          {user?.trackmaniaId && (
            <Typography
              variant="body2"
              color="text.secondary"
            >
              <strong>Trackmania ID:</strong> {user.trackmaniaId}
            </Typography>
          )}
        </Stack>

        <Divider sx={{ my: 3 }} />

        <CardActions
          sx={{ justifyContent: 'center', gap: 2, flexWrap: 'wrap' }}
        >
          <Button
            variant="contained"
            startIcon={<EditIcon />}
            size="large"
            sx={{ borderRadius: 2, px: 4, minWidth: 180 }}
          >
            Edit Profile
          </Button>
          <Button
            variant="outlined"
            startIcon={<SportsEsportsIcon />}
            size="large"
            sx={{ borderRadius: 2, px: 3, minWidth: 220 }}
            onClick={handleConnectTrackmania}
            disabled={connecting}
          >
            Sync Trackmania Account
          </Button>
        </CardActions>
      </Card>
    </Box>
  );
}
