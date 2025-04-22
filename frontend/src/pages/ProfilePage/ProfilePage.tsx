import {
  Edit as EditIcon,
  Mail as MailIcon,
  Person as PersonIcon,
} from '@mui/icons-material';
import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CircularProgress,
  Stack,
  Typography,
} from '@mui/material';
import { useFetchMe } from '../../api/hooks/authHooks';

export default function ProfilePage() {
  const { data: user, isLoading, error, refetch } = useFetchMe();

  if (isLoading) {
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
          width: 400,
          borderRadius: 4,
          boxShadow: 6,
          p: 3,
          bgcolor: 'background.paper',
        }}
      >
        <Stack
          spacing={3}
          alignItems="center"
        >
          <Avatar
            sx={{
              width: 96,
              height: 96,
              bgcolor: 'primary.main',
              fontSize: 48,
            }}
          >
            <PersonIcon fontSize="inherit" />
          </Avatar>
          <Typography
            variant="h5"
            fontWeight={600}
          >
            {user?.username}
          </Typography>
          <Stack
            direction="row"
            alignItems="center"
            spacing={1}
          >
            <MailIcon color="action" />
            <Typography
              variant="body1"
              color="text.secondary"
            >
              {user?.email}
            </Typography>
          </Stack>
          {user?.role && (
            <Typography
              variant="body2"
              color="primary"
              sx={{
                bgcolor: 'primary.50',
                px: 2,
                py: 0.5,
                borderRadius: 2,
                fontWeight: 500,
                mt: 1,
                letterSpacing: 1,
                textTransform: 'uppercase',
              }}
            >
              {user.role}
            </Typography>
          )}
        </Stack>
        <CardActions sx={{ justifyContent: 'center', mt: 3 }}>
          <Button
            variant="contained"
            startIcon={<EditIcon />}
            size="large"
            sx={{ borderRadius: 2, px: 4 }}
          >
            Edit Profile
          </Button>
        </CardActions>
      </Card>
    </Box>
  );
}
