import {
  Box,
  Button,
  CircularProgress,
  Container,
  Typography,
} from '@mui/material';
import { useEffect } from 'react';
import {
  useCheckMatchStatus,
  useGetQueueTeamsCount,
  useJoinQueue,
  useLeaveQueue,
} from 'api/hooks/matchmakingHooks';
import { useFetchSeasons } from 'api/hooks/seasonsHooks';
import { useFetchTeams } from 'api/hooks/teamsHooks';
import { useAppForm } from 'hooks/form/useAppForm';
import { queueDtoValidator } from 'validators/matchmaking/queueDtoValidator';

export default function MatchmakingQueuePage() {
  const form = useAppForm({
    defaultValues: {
      seasonId: '',
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

  const seasonId = form.getFieldValue('seasonId');
  const teamId = form.getFieldValue('teamId');

  const { data: seasons, isLoading: seasonsLoading } = useFetchSeasons();
  const {
    data: teams,
    isLoading: isTeamsLoading,
    refetch: refetchTeams,
  } = useFetchTeams(seasonId);

  useEffect(() => {
    console.log('seasonId', seasonId);
    const fetchTeams = async () => {
      if (seasonId) {
        await refetchTeams();
      }
    };
    if (seasonId) {
      void fetchTeams();
    }
  }, [form, seasonId, refetchTeams]);

  const { data: teamsCount, isLoading: isLoadingTeamsCount } =
    useGetQueueTeamsCount(seasonId);
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

  if (seasonsLoading || isTeamsLoading) {
    return <CircularProgress />;
  }

  return (
    <Container maxWidth="md">
      <Box sx={{ mt: 4 }}>
        <Typography
          variant="h4"
          gutterBottom
        >
          Matchmaking Queue
        </Typography>
        <form
          onSubmit={(e) => {
            e.preventDefault();
            void form.handleSubmit();
          }}
        >
          <form.AppField name="seasonId">
            {(field) => (
              <field.Select
                label="Select Season"
                options={(seasons ?? []).map((season) => ({
                  key: season.id,
                  value: season.name,
                }))}
              />
            )}
          </form.AppField>

          <form.AppField name="teamId">
            {(field) => (
              <field.Select
                label="Select Team"
                options={(teams ?? []).map((team) => ({
                  key: team.id,
                  value: team.name,
                }))}
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

        <Box sx={{ mt: 4 }}>
          <Typography variant="h6">Queue Status</Typography>
          {isLoadingTeamsCount ? (
            <CircularProgress />
          ) : (
            <Typography>
              Teams in Queue:{' '}
              {typeof teamsCount === 'number'
                ? teamsCount
                : (teamsCount?.teamsCount ?? 'N/A')}
            </Typography>
          )}
        </Box>

        <Box sx={{ mt: 4 }}>
          <Typography variant="h6">Match Status</Typography>
          {isLoadingMatchStatus ? (
            <CircularProgress />
          ) : (
            <Typography>
              {isInMatch?.isInMatch
                ? 'Your team is currently in a match!'
                : 'Your team is not in a match.'}
            </Typography>
          )}
        </Box>
      </Box>
    </Container>
  );
}
