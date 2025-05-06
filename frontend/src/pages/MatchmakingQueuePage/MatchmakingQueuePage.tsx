import {
  Box,
  Button,
  CircularProgress,
  Container,
  Paper,
  Typography,
} from '@mui/material';
import {
  useCheckMatchStatus,
  useGetQueueTeamsCount,
  useJoinQueue,
  useLeaveQueue,
} from 'api/hooks/matchmakingHooks';
import { useFetchTeams } from 'api/hooks/teamsHooks';
import withAuth from 'hocs/withAuth';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { Permission } from 'types/enums/permission';
import { queueDtoValidator } from 'validators/matchmaking/queueDtoValidator';

function MatchmakingQueuePage() {
  const form = useAppForm({
    defaultValues: {
      leagueId: '',
      teamId: '',
    },
    validators: {
      onSubmit: queueDtoValidator,
    },
    onSubmitMeta: {} as { isJoining: boolean },
    onSubmit: async ({ value, meta }) => {
      if (meta.isJoining) {
        await joinQueueAsync(value);
      } else {
        await leaveQueueAsync(value);
      }
    },
  });

  const leagueId = form.getFieldValue('leagueId');
  const teamId = form.getFieldValue('teamId');

  const { user } = useAuth();

  const { data: teams, isLoading: isTeamsLoading } = useFetchTeams();

  const myTeams =
    teams?.filter((team) =>
      team.players.some((player) => player.id === user?.id),
    ) ?? [];
  const myTeamOptions = myTeams.map((team) => ({
    key: team.id,
    value: team.name,
  }));

  const myTeamLeagues = myTeams.flatMap((team) => team.leagues);
  const myTeamLeaguesOptions = myTeamLeagues.map((league) => ({
    key: league.id,
    value: league.name,
  }));

  const { data: teamsCount, isLoading: isLoadingTeamsCount } =
    useGetQueueTeamsCount(leagueId);
  const { data: isInMatch, isLoading: isLoadingMatchStatus } =
    useCheckMatchStatus(teamId);

  const { mutateAsync: joinQueueAsync, isPending: isJoinPending } =
    useJoinQueue();
  const { mutateAsync: leaveQueueAsync, isPending: isLeavePending } =
    useLeaveQueue();

  const handleJoinQueue = async () => {
    await form.handleSubmit({ isJoining: true });
  };

  const handleLeaveQueue = async () => {
    await form.handleSubmit({ isJoining: false });
  };

  if (isTeamsLoading) {
    return (
      <Box sx={{ display: 'flex', justifyContent: 'center', mt: 6 }}>
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Container maxWidth="md">
      <Box sx={{ mt: 4 }}>
        <Typography
          variant="h4"
          fontWeight={700}
          gutterBottom
        >
          Matchmaking
        </Typography>

        <Paper
          elevation={2}
          sx={{
            p: 3,
            borderRadius: 2,
            backgroundColor: 'background.paper',
            boxShadow: 2,
          }}
        >
          <form
            onSubmit={(e) => {
              e.preventDefault();
              void form.handleSubmit();
            }}
          >
            <form.AppField name="teamId">
              {(field) => (
                <field.Select
                  label="Select Team"
                  options={myTeamOptions}
                />
              )}
            </form.AppField>

            <form.AppField name="leagueId">
              {(field) => (
                <field.Select
                  label="Select League"
                  options={myTeamLeaguesOptions}
                />
              )}
            </form.AppField>
          </form>

          <Box sx={{ mt: 3, display: 'flex', gap: 2 }}>
            <Button
              variant="contained"
              color="primary"
              onClick={() => void handleJoinQueue()}
              disabled={isJoinPending}
            >
              {isJoinPending ? <CircularProgress size={24} /> : 'Join Queue'}
            </Button>
            <Button
              variant="outlined"
              color="secondary"
              onClick={() => void handleLeaveQueue()}
              disabled={isLeavePending}
            >
              {isLeavePending ? <CircularProgress size={24} /> : 'Leave Queue'}
            </Button>
          </Box>
        </Paper>

        <Paper
          elevation={2}
          sx={{
            mt: 4,
            p: 3,
            borderRadius: 2,
            backgroundColor: 'background.paper',
            boxShadow: 2,
          }}
        >
          <Typography
            variant="h6"
            fontWeight={600}
          >
            Queue Status
          </Typography>
          {isLoadingTeamsCount ? (
            <CircularProgress />
          ) : (
            <Typography sx={{ mt: 1 }}>
              Teams in Queue:{' '}
              {typeof teamsCount === 'number'
                ? teamsCount
                : (teamsCount?.teamsCount ?? 'N/A')}
            </Typography>
          )}
        </Paper>

        <Paper
          elevation={2}
          sx={{
            mt: 4,
            p: 3,
            borderRadius: 2,
            backgroundColor: 'background.paper',
            boxShadow: 2,
          }}
        >
          <Typography
            variant="h6"
            fontWeight={600}
          >
            Match Status
          </Typography>
          {isLoadingMatchStatus ? (
            <CircularProgress />
          ) : (
            <Typography sx={{ mt: 1 }}>
              {isInMatch?.isInMatch
                ? '✅ Your team is currently in a match!'
                : '❌ Your team is not in a match.'}
            </Typography>
          )}
        </Paper>
      </Box>
    </Container>
  );
}

export default withAuth(MatchmakingQueuePage, {
  permission: Permission.ManageQueue,
});
