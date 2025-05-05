import SportsEsportsIcon from '@mui/icons-material/SportsEsports';
import { Box, Chip, List, ListItem, Stack, Typography } from '@mui/material';
import { Link, useSearchParams } from 'react-router-dom';
import { useFetchTeams } from 'api/hooks/teamsHooks';
import StatusHandler from 'components/StatusHandler';
import UserAvatar from 'components/UserAvatar';
import LeaguesFilter from 'components/leagues/LeaguesFilter';
import CreateTeamButton from 'components/teams/CreateTeamButton';
import DeleteTeamButton from 'components/teams/DeleteTeamButton';
import UpdateTeamButton from 'components/teams/UpdateTeamButton';
import PARAMS from 'constants/params';
import withAuth from 'hocs/withAuth';
import { Permission } from 'types/enums/permission';
import { getResultsPageLink } from 'utils/pageLinkUtils';
import { isSelectOptionValid } from 'utils/selectOptionUtils';

function TeamsPage() {
  const [searchParams, setSearchParams] = useSearchParams();
  const leagueId = searchParams.get(PARAMS.LEAGUE_ID) ?? '';

  const {
    data: teams,
    isLoading: isTeamsLoading,
    error: teamsError,
  } = useFetchTeams();

  const filteredTeams = leagueId
    ? teams?.filter((team) =>
        team.leagues.some((league) => league.id === leagueId),
      )
    : teams;

  return (
    <Stack
      sx={{ p: 4, maxWidth: 800, mx: 'auto' }}
      spacing={1}
    >
      <Box>
        <CreateTeamButton />
      </Box>

      <LeaguesFilter
        leagueId={leagueId}
        onLeagueChange={(leagueId) => {
          setSearchParams(
            isSelectOptionValid(leagueId)
              ? { [PARAMS.LEAGUE_ID]: leagueId }
              : {},
          );
        }}
      />

      <StatusHandler
        isLoading={isTeamsLoading}
        error={teamsError}
        isEmpty={!filteredTeams || filteredTeams.length === 0}
        emptyMessage={
          leagueId ? 'No teams found in this league.' : 'No teams available.'
        }
      >
        <List sx={{ borderRadius: 2, boxShadow: 2 }}>
          {filteredTeams?.map((team) => (
            <ListItem
              key={team.id}
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
                  <UpdateTeamButton team={team} />
                  <DeleteTeamButton team={team} />
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
                    imageUrl={team.logoUrl}
                    name={team.name}
                    size={24}
                  />
                  <Link
                    to={getResultsPageLink(leagueId, team.id)}
                    style={{ textDecoration: 'none', color: 'inherit' }}
                  >
                    <Typography
                      variant="h6"
                      fontWeight={700}
                    >
                      {team.name}
                    </Typography>
                  </Link>
                </Stack>

                <Stack
                  direction="row"
                  alignItems="center"
                  spacing={1}
                >
                  <SportsEsportsIcon color="primary" />
                  <Typography
                    variant="h5"
                    fontWeight={800}
                    color="primary"
                  >
                    {team.elo}
                  </Typography>
                </Stack>
              </Stack>

              <Typography
                variant="caption"
                color="text.secondary"
                mt={0.5}
              >
                Created by: {team.user.username}
              </Typography>

              <Stack
                direction="row"
                flexWrap="wrap"
                gap={4}
                mt={2}
              >
                <Stack
                  spacing={1}
                  flex={1}
                  minWidth={200}
                >
                  <Typography
                    variant="body2"
                    fontWeight={600}
                  >
                    Players
                  </Typography>
                  {team.players.length > 0 ? (
                    <Stack
                      direction="row"
                      flexWrap="wrap"
                      gap={1}
                    >
                      {team.players.map((player) => (
                        <Chip
                          key={player.id}
                          avatar={
                            <UserAvatar
                              imageUrl={player.profilePhotoUrl}
                              name={player.trackmaniaName}
                              size={18}
                              marginLeft={5}
                            />
                          }
                          label={player.trackmaniaName}
                          size="small"
                        />
                      ))}
                    </Stack>
                  ) : (
                    <Typography
                      variant="body2"
                      color="text.secondary"
                    >
                      No players available.
                    </Typography>
                  )}
                </Stack>

                <Stack
                  spacing={1}
                  flex={1}
                  minWidth={200}
                >
                  <Typography
                    variant="body2"
                    fontWeight={600}
                  >
                    Leagues
                  </Typography>
                  {team.leagues.length > 0 ? (
                    <Stack
                      direction="row"
                      flexWrap="wrap"
                      gap={1}
                    >
                      {team.leagues.map((league) => (
                        <Chip
                          key={league.id}
                          avatar={
                            <UserAvatar
                              imageUrl={league.logoUrl}
                              name={league.name}
                              size={18}
                              marginLeft={5}
                            />
                          }
                          label={league.name}
                          size="small"
                        />
                      ))}
                    </Stack>
                  ) : (
                    <Typography
                      variant="body2"
                      color="text.secondary"
                    >
                      No leagues available.
                    </Typography>
                  )}
                </Stack>
              </Stack>
            </ListItem>
          ))}
        </List>
      </StatusHandler>
    </Stack>
  );
}

export default withAuth(TeamsPage, {
  permission: Permission.LoggedIn,
});
