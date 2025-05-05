import { ArrowDownward, ArrowUpward } from '@mui/icons-material';
import { Box, Chip, List, ListItem, Stack, Typography } from '@mui/material';
import dayjs from 'dayjs';
import { useSearchParams } from 'react-router-dom';
import { useFetchResults } from 'api/hooks/resultsHooks';
import StatusHandler from 'components/StatusHandler';
import UserAvatar from 'components/UserAvatar';
import LeaguesFilter from 'components/leagues/LeaguesFilter';
import CreateResultButton from 'components/results/CreateResultButton';
import DeleteResultButton from 'components/results/DeleteResultButton';
import UpdateResultButton from 'components/results/UpdateResultButton';
import TeamsFilter from 'components/teams/TeamsFilter';
import PARAMS from 'constants/params';
import withAuth from 'hocs/withAuth';
import { Permission } from 'types/enums/permission';
import { isSelectOptionValid } from 'utils/selectOptionUtils';
import { isUuidEmpty } from 'utils/stringUtils';

function ResultsPage() {
  const [searchParams, setSearchParams] = useSearchParams();
  const leagueId = searchParams.get(PARAMS.LEAGUE_ID) ?? '';
  const teamId = searchParams.get(PARAMS.TEAM_ID) ?? '';

  const {
    data: results,
    isLoading: isResultsLoading,
    error: resultsError,
  } = useFetchResults();

  const filteredResults = results?.filter((result) => {
    const matchesLeague = leagueId ? result.league.id === leagueId : true;
    const matchesTeam = teamId
      ? result.team.id === teamId || result.opponentTeam.id === teamId
      : true;
    return matchesLeague && matchesTeam;
  });

  return (
    <Stack
      sx={{ p: 4, maxWidth: 800, mx: 'auto' }}
      spacing={1}
    >
      <Box>
        <CreateResultButton />
      </Box>

      <LeaguesFilter
        leagueId={leagueId}
        onLeagueChange={(leagueId) => {
          setSearchParams((prevParams) => {
            const newParams = new URLSearchParams(prevParams);
            if (isSelectOptionValid(leagueId)) {
              newParams.set(PARAMS.LEAGUE_ID, leagueId);
            } else {
              newParams.delete(PARAMS.LEAGUE_ID);
            }
            return newParams;
          });
        }}
      />

      <TeamsFilter
        teamId={teamId}
        onTeamChange={(teamId) => {
          setSearchParams((prevParams) => {
            const newParams = new URLSearchParams(prevParams);
            if (isSelectOptionValid(teamId)) {
              newParams.set(PARAMS.TEAM_ID, teamId);
            } else {
              newParams.delete(PARAMS.TEAM_ID);
            }
            return newParams;
          });
        }}
      />

      <StatusHandler
        isLoading={isResultsLoading}
        error={resultsError}
        isEmpty={!filteredResults || filteredResults.length === 0}
        emptyMessage={
          leagueId && teamId
            ? 'No results found for this league and team.'
            : leagueId
              ? 'No results found in this league.'
              : teamId
                ? 'No results found for this team.'
                : 'No results available.'
        }
      >
        <List sx={{ borderRadius: 2, boxShadow: 2 }}>
          {filteredResults?.map((result) => (
            <ListItem
              key={result.id}
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
                  <UpdateResultButton result={result} />
                  <DeleteResultButton result={result} />
                </Stack>
              }
            >
              <Stack
                direction="row"
                justifyContent="space-between"
                alignItems="center"
                width="100%"
              >
                <Typography
                  variant="h6"
                  fontWeight={700}
                >
                  {`${result.team.name} vs ${result.opponentTeam.name}`}
                </Typography>

                <Stack
                  direction="row"
                  spacing={4}
                  alignItems="flex-start"
                >
                  <Stack alignItems="center">
                    <Typography
                      variant="h5"
                      fontWeight={800}
                      color="primary"
                    >
                      {result.score}
                    </Typography>
                    <Stack
                      direction="row"
                      alignItems="center"
                      spacing={0.5}
                    >
                      {result.eloDiff > 0 ? (
                        <ArrowUpward
                          fontSize="small"
                          sx={{ color: 'success.main' }}
                        />
                      ) : result.eloDiff < 0 ? (
                        <ArrowDownward
                          fontSize="small"
                          sx={{ color: 'error.main' }}
                        />
                      ) : null}
                      <Typography
                        variant="body2"
                        fontWeight={800}
                        sx={{
                          color:
                            result.eloDiff > 0
                              ? 'success.main'
                              : result.eloDiff < 0
                                ? 'error.main'
                                : 'text.primary',
                        }}
                      >
                        {Math.abs(result.eloDiff)}
                      </Typography>
                    </Stack>
                  </Stack>

                  <Typography
                    variant="h5"
                    fontWeight={800}
                    color="text.secondary"
                  >
                    -
                  </Typography>

                  <Stack alignItems="center">
                    <Typography
                      variant="h5"
                      fontWeight={800}
                      color="primary"
                    >
                      {result.opponentScore}
                    </Typography>
                    <Stack
                      direction="row"
                      alignItems="center"
                      spacing={0.5}
                    >
                      {result.opponentEloDiff > 0 ? (
                        <ArrowUpward
                          fontSize="small"
                          sx={{ color: 'success.main' }}
                        />
                      ) : result.opponentEloDiff < 0 ? (
                        <ArrowDownward
                          fontSize="small"
                          sx={{ color: 'error.main' }}
                        />
                      ) : null}
                      <Typography
                        variant="body2"
                        fontWeight={800}
                        sx={{
                          color:
                            result.opponentEloDiff > 0
                              ? 'success.main'
                              : result.opponentEloDiff < 0
                                ? 'error.main'
                                : 'text.primary',
                        }}
                      >
                        {Math.abs(result.opponentEloDiff)}
                      </Typography>
                    </Stack>
                  </Stack>
                </Stack>
              </Stack>

              <Typography
                variant="caption"
                color="text.secondary"
                mt={0.5}
              >
                {dayjs(result.startDate).format('YYYY-MM-DD HH:mm')} -{' '}
                {dayjs(result.endDate).format('YYYY-MM-DD HH:mm')}
              </Typography>

              <Typography
                variant="caption"
                color="text.secondary"
                mt={0.5}
              >
                {isUuidEmpty(result.user.id)
                  ? 'Matchmaking Result'
                  : `Created by: ${result.user.username}`}
              </Typography>

              <Stack
                direction="row"
                flexWrap="wrap"
                gap={1}
                mt={2}
              >
                <Chip
                  avatar={
                    <UserAvatar
                      imageUrl={result.league.logoUrl}
                      name={result.league.name}
                      size={18}
                      marginLeft={5}
                    />
                  }
                  label={result.league.name}
                  size="small"
                />
              </Stack>
            </ListItem>
          ))}
        </List>
      </StatusHandler>
    </Stack>
  );
}

export default withAuth(ResultsPage, {
  permission: Permission.LoggedIn,
});
