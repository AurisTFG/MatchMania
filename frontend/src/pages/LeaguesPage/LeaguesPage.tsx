import { Box, Chip, List, ListItem, Stack, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { Link } from 'react-router-dom';
import { useFetchLeagues } from 'api/hooks/leaguesHooks';
import StatusHandler from 'components/StatusHandler';
import UserAvatar from 'components/UserAvatar';
import CreateLeagueButton from 'components/leagues/CreateLeagueButton';
import DeleteLeagueButton from 'components/leagues/DeleteLeagueButton';
import UpdateLeagueButton from 'components/leagues/UpdateLeagueButton';
import withAuth from 'hocs/withAuth';
import { Permission } from 'types/enums/permission';
import { getTeamsPageLink } from 'utils/pageLinkUtils';

function LeaguesPage() {
  const {
    data: leagues,
    isLoading: leaguesLoading,
    error: leaguesError,
  } = useFetchLeagues();

  return (
    <Stack
      sx={{ p: 4, maxWidth: 800, mx: 'auto' }}
      spacing={1}
    >
      <Box>
        <CreateLeagueButton />
      </Box>

      <StatusHandler
        isLoading={leaguesLoading}
        error={leaguesError}
        isEmpty={!leagues || leagues.length === 0}
      >
        <List sx={{ borderRadius: 2, boxShadow: 2 }}>
          {leagues?.map((league) => (
            <ListItem
              key={league.id}
              sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'stretch',
                borderRadius: 2,
                px: 3,
                py: 2,
                my: 2,
                backgroundColor: 'background.paper',
                transition: 'background-color 0.2s ease',
                '&:hover': {
                  backgroundColor: 'action.hover',
                },
              }}
              secondaryAction={
                <Stack
                  direction="row"
                  spacing={1}
                >
                  <UpdateLeagueButton league={league} />
                  <DeleteLeagueButton league={league} />
                </Stack>
              }
            >
              <Stack
                direction="row"
                justifyContent="space-between"
                alignItems="center"
                width="100%"
              >
                <Stack
                  direction="row"
                  alignItems="center"
                  spacing={1}
                >
                  <UserAvatar
                    imageUrl={league.logoUrl}
                    name={league.name}
                    size={24}
                  />
                  <Link
                    to={getTeamsPageLink(league.id)}
                    style={{ textDecoration: 'none', color: 'inherit' }}
                  >
                    <Typography
                      variant="h6"
                      fontWeight={700}
                    >
                      {league.name}
                    </Typography>
                  </Link>
                </Stack>

                <Typography variant="body2">
                  {`${dayjs(league.startDate).format('YYYY-MM-DD')} - ${dayjs(league.endDate).format('YYYY-MM-DD')}`}
                </Typography>
              </Stack>

              <Typography
                variant="caption"
                color="text.secondary"
              >
                Created by: {league.user.username}
              </Typography>

              <Stack
                direction="column"
                spacing={1}
                mt={2}
              >
                <Typography
                  variant="body2"
                  fontWeight={600}
                >
                  Tracks
                </Typography>

                {league.tracks.length > 0 ? (
                  <Stack
                    direction="row"
                    flexWrap="wrap"
                    gap={1}
                  >
                    {league.tracks.map((track) => (
                      <Chip
                        key={track.id}
                        label={
                          <span
                            dangerouslySetInnerHTML={{
                              __html: MPStyle.Parser.toHTML(track.name),
                            }}
                          />
                        }
                        size="small"
                      />
                    ))}
                  </Stack>
                ) : (
                  <Typography
                    variant="body2"
                    color="text.secondary"
                  >
                    No tracks available.
                  </Typography>
                )}
              </Stack>
            </ListItem>
          ))}
        </List>
      </StatusHandler>
    </Stack>
  );
}

export default withAuth(LeaguesPage, {
  permission: Permission.LoggedIn,
});
