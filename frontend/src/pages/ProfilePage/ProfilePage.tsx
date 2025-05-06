import {
  Edit as EditIcon,
  Mail as MailIcon,
  ShieldOutlined,
  SportsEsports as SportsEsportsIcon,
} from '@mui/icons-material';
import {
  Box,
  Button,
  Card,
  CardActions,
  Chip,
  CircularProgress,
  Divider,
  Grid,
  Stack,
  Tooltip,
  Typography,
} from '@mui/material';
import { useState } from 'react';
import { useFetchMe } from 'api/hooks/authHooks';
import { useGetTrackmaniaOAuthUrl } from 'api/hooks/trackmaniaOauthHooks';
import UserAvatar from 'components/UserAvatar';
import withAuth from 'hocs/withAuth';
import { Permission } from 'types/enums/permission';
import { getCountryIconUrl, getCountryName } from 'utils/countryUtils';
import { isStringEmpty, isUuidEmpty } from 'utils/stringUtils';

function ProfilePage() {
  const { data: user, isLoading, error, refetch } = useFetchMe();
  const [connecting, setConnecting] = useState(false);
  const { mutateAsync: getTrackmaniaOAuthUrlAsync } =
    useGetTrackmaniaOAuthUrl();

  const handleConnectTrackmania = async () => {
    setConnecting(true);
    const urlDto = await getTrackmaniaOAuthUrlAsync();
    window.location.href = urlDto.url;
    setConnecting(false);
  };

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
        sx={(theme) => ({
          maxWidth: 400,
          mx: 'auto',
          mt: theme.spacing(8),
          textAlign: 'center',
          p: theme.spacing(3),
          borderRadius: theme.shape.borderRadius,
          boxShadow: theme.shadows[2],
        })}
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
      alignItems="flex-start"
      py={(theme) => theme.spacing(8)}
      px={(theme) => theme.spacing(2)}
      bgcolor={(theme) => theme.palette.background.default}
    >
      <Card
        sx={(theme) => ({
          width: '100%',
          maxWidth: 1000,
          borderRadius: 4,
          boxShadow: theme.shadows[2],
          p: theme.spacing(4),
          bgcolor: theme.palette.background.paper,
          backdropFilter: theme.palette.mode === 'dark' ? 'blur(6px)' : 'none',
        })}
      >
        <Stack
          direction={{ xs: 'column', md: 'row' }}
          spacing={4}
        >
          <Box flex={2}>
            <Stack
              direction="column"
              alignItems="center"
              spacing={2}
            >
              <UserAvatar
                imageUrl={user?.profilePhotoUrl}
                name={user?.username}
                size={80}
              />

              <Typography
                variant="h6"
                fontWeight={700}
              >
                {user?.username}
              </Typography>

              {user?.roles && user.roles.length > 0 && (
                <Stack
                  direction="row"
                  spacing={1}
                  flexWrap="wrap"
                  justifyContent="center"
                  sx={{ mt: 2 }}
                >
                  {user.roles.map((role) => (
                    <Tooltip
                      key={role.name}
                      title="User Role"
                    >
                      <Chip
                        icon={<ShieldOutlined fontSize="small" />}
                        label={role.name}
                        color="primary"
                        variant="filled"
                      />
                    </Tooltip>
                  ))}
                </Stack>
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

              <Chip
                icon={
                  <Box sx={{ pl: 1, display: 'flex', alignItems: 'center' }}>
                    <Box
                      component="img"
                      src={getCountryIconUrl(user?.country)}
                      alt={user?.country ?? 'N/A'}
                      sx={{
                        width: 20,
                        height: 15,
                        borderRadius: '2px',
                      }}
                    />
                  </Box>
                }
                label={getCountryName(user?.country)}
                variant="outlined"
                sx={{ mt: 1 }}
              />

              <Typography
                variant="body2"
                color="text.secondary"
              >
                <strong>Trackmania Name: </strong>
                {isStringEmpty(user?.trackmaniaName)
                  ? 'N/A'
                  : user?.trackmaniaName}
              </Typography>

              <Typography
                variant="body2"
                color="text.secondary"
              >
                <strong>Trackmania ID: </strong>
                {isUuidEmpty(user?.trackmaniaId) ? 'N/A' : user?.trackmaniaId}
              </Typography>
            </Stack>

            <Divider sx={{ my: (theme) => theme.spacing(3) }} />

            <CardActions
              sx={{ justifyContent: 'center', gap: 2, flexWrap: 'wrap' }}
            >
              <Button
                variant="contained"
                startIcon={<EditIcon />}
                size="large"
                sx={(theme) => ({
                  borderRadius: 2,
                  px: theme.spacing(4),
                  minWidth: 180,
                })}
              >
                Edit Profile
              </Button>
              <Button
                variant="outlined"
                startIcon={<SportsEsportsIcon />}
                size="large"
                onClick={() => void handleConnectTrackmania()}
                disabled={connecting}
                sx={(theme) => ({
                  borderRadius: 2,
                  px: theme.spacing(4),
                  minWidth: 240,
                })}
              >
                Sync Trackmania Account
              </Button>
            </CardActions>
          </Box>

          <Box flex={2.5}>
            <Typography
              variant="h6"
              gutterBottom
              fontWeight={700}
            >
              Your Favorite Tracks
            </Typography>
            <Grid
              container
              spacing={2}
              sx={{ maxHeight: 500, overflowY: 'auto', pr: 1 }}
            >
              {user?.tracks.length ? (
                user.tracks.map((track) => (
                  <Grid
                    key={track.id}
                    size={{ xs: 12, sm: 6, md: 4 }}
                  >
                    <Card
                      sx={(theme) => ({
                        p: theme.spacing(1),
                        borderRadius: 2,
                        boxShadow: theme.shadows[2],
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                      })}
                    >
                      <Box
                        component="img"
                        src={track.thumbnailUrl}
                        alt={track.name}
                        sx={{
                          width: '100%',
                          height: 100,
                          objectFit: 'cover',
                          borderRadius: '10px',
                        }}
                      />
                      <Typography
                        variant="body2"
                        sx={{
                          mt: 1,
                          textAlign: 'center',
                          wordBreak: 'break-word',
                        }}
                        dangerouslySetInnerHTML={{
                          __html: MPStyle.Parser.toHTML(track.name),
                        }}
                      />
                    </Card>
                  </Grid>
                ))
              ) : (
                <Grid size={12}>
                  <Typography
                    variant="body2"
                    color="text.secondary"
                    align="center"
                  >
                    No favorite tracks yet.
                  </Typography>
                </Grid>
              )}
            </Grid>
          </Box>
        </Stack>
      </Card>
    </Box>
  );
}

export default withAuth(ProfilePage, {
  permission: Permission.LoggedIn,
});
