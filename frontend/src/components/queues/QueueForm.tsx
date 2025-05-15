import { Box, Button, CircularProgress, Paper } from '@mui/material';
import { useEffect, useState } from 'react';
import { useJoinQueue, useLeaveQueue } from 'api/hooks/queueHooks';
import { useFetchTeams } from 'api/hooks/teamsHooks';
import EndMatchDialog from 'components/matches/EndMatchDialog';
import withAuth from 'hocs/withAuth';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { Permission } from 'types/enums/permission';
import { queueDtoValidator } from 'validators/matchmaking/queueDtoValidator';

type QueueFormProps = {
  queueState: 'idle' | 'queued' | 'in-match';
  selectedTeamId?: string;
  selectedLeagueId?: string;
};

function QueueForm({
  queueState,
  selectedTeamId,
  selectedLeagueId,
}: QueueFormProps) {
  const [leagueOptions, setLeagueOptions] = useState<
    { key: string; value: string }[]
  >([]);
  const [openEndMatchDialog, setOpenEndMatchDialog] = useState(false);

  const { user } = useAuth();
  const { data: teams } = useFetchTeams();
  const { mutateAsync: joinQueueAsync, isPending: isJoinPending } =
    useJoinQueue();
  const { mutateAsync: leaveQueueAsync, isPending: isLeavePending } =
    useLeaveQueue();

  const myTeamOptions = (
    teams?.filter((team) =>
      team.players.some((player) => player.id === user?.id),
    ) ?? []
  ).map((team) => ({
    key: team.id,
    value: team.name,
  }));

  useEffect(() => {
    if (selectedTeamId) {
      const selectedTeam = teams?.find((team) => team.id === selectedTeamId);
      setLeagueOptions(
        selectedTeam?.leagues.map((league) => ({
          key: league.id,
          value: league.name,
        })) ?? [],
      );
    } else {
      setLeagueOptions([]);
    }
  }, [selectedTeamId, teams]);

  const form = useAppForm({
    defaultValues: {
      leagueId: selectedLeagueId ?? '',
      teamId: selectedTeamId ?? '',
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

  const handleTeamChange = (teamId: string) => {
    const selectedTeam = teams?.find((team) => team.id === teamId);
    if (selectedTeam) {
      setLeagueOptions(
        selectedTeam.leagues.map((league) => ({
          key: league.id,
          value: league.name,
        })),
      );
    } else {
      setLeagueOptions([]);
    }
  };

  const handleClick = () => {
    if (queueState === 'idle') {
      void form.handleSubmit({ isJoining: true });
    } else if (queueState === 'queued') {
      void form.handleSubmit({ isJoining: false });
      form.reset();
    } else {
      setOpenEndMatchDialog(true);
      form.reset();
    }
  };

  const getButtonProps = () => {
    switch (queueState) {
      case 'idle':
        return {
          label: 'Join Queue',
          variant: 'contained' as const,
          color: 'primary' as const,
          loading: isJoinPending,
          disabled: isJoinPending,
        };
      case 'queued':
        return {
          label: 'Leave Queue',
          variant: 'outlined' as const,
          color: 'secondary' as const,
          loading: isLeavePending,
          disabled: isLeavePending,
        };
      case 'in-match':
        return {
          label: 'End Match',
          variant: 'contained' as const,
          color: 'success' as const,
          loading: false,
          disabled: false,
        };
    }
  };

  const { label, variant, color, loading, disabled } = getButtonProps();

  return (
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
        <form.AppField
          name="teamId"
          listeners={{
            onChange: ({ value }) => {
              handleTeamChange(value);
            },
          }}
        >
          {(field) => (
            <field.Select
              label="Select Team"
              options={myTeamOptions}
              disable={queueState !== 'idle'}
            />
          )}
        </form.AppField>

        <form.AppField name="leagueId">
          {(field) => (
            <field.Select
              label="Select League"
              options={leagueOptions}
              disable={leagueOptions.length === 0 || queueState !== 'idle'}
            />
          )}
        </form.AppField>
      </form>

      <Box sx={{ mt: 3 }}>
        <Button
          variant={variant}
          color={color}
          onClick={handleClick}
          disabled={disabled}
        >
          {loading ? <CircularProgress size={24} /> : label}
        </Button>
      </Box>

      <EndMatchDialog
        open={openEndMatchDialog}
        onClose={() => {
          setOpenEndMatchDialog(false);
        }}
      />
    </Paper>
  );
}

export default withAuth(QueueForm, {
  permission: Permission.ManageQueue,
  redirect: false,
});
