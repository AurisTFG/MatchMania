import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
} from '@mui/material';
import { useFetchLeagues } from 'api/hooks/leaguesHooks';
import { useFetchPlayers } from 'api/hooks/playersHooks';
import { useAppForm } from 'hooks/form/useAppForm';
import { CreateTeamDto } from 'types/dtos/requests/teams/createTeamDto';
import { createTeamDtoValidator } from 'validators/teams/createTeamDtoValidator';

type BaseTeamMutateDialogProps = {
  title: string;
  buttonText: string;
  team: CreateTeamDto;
  submitAsync: (payload: CreateTeamDto) => Promise<void>;
  open: boolean;
  onClose: () => void;
};

export default function BaseTeamMutateDialog({
  title,
  buttonText,
  team,
  submitAsync,
  open,
  onClose,
}: BaseTeamMutateDialogProps) {
  const { data: leagues } = useFetchLeagues();
  const { data: players } = useFetchPlayers();

  const form = useAppForm({
    defaultValues: team,
    validators: {
      onSubmit: createTeamDtoValidator,
    },
    onSubmit: async ({ value }) => {
      await submitAsync(value);

      handleClose();
    },
  });

  const handleClose = () => {
    form.reset();
    onClose();
  };

  const leaguesOptions =
    leagues?.map((league) => ({
      key: league.id,
      value: league.name,
    })) ?? [];

  const playersOptions =
    players?.map((player) => ({
      key: player.id,
      value: player.trackmaniaName ?? 'Unknown Player',
    })) ?? [];

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      fullWidth
      maxWidth="sm"
    >
      <DialogTitle>{title}</DialogTitle>
      <DialogContent dividers>
        <Box
          component="form"
          noValidate
          autoComplete="off"
          sx={{ mt: 2 }}
        >
          <form.AppField name="name">
            {(field) => <field.Text label="Team Name" />}
          </form.AppField>

          <form.AppField name="logoUrl">
            {(field) => <field.Text label="Logo URL (Optional)" />}
          </form.AppField>

          <form.AppField name="playerIds">
            {(field) => (
              <field.MultiSelect
                label="Select Players"
                options={playersOptions}
              />
            )}
          </form.AppField>

          <form.AppField name="leagueIds">
            {(field) => (
              <field.MultiSelect
                label="Select Leagues"
                options={leaguesOptions}
              />
            )}
          </form.AppField>
        </Box>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Cancel</Button>
        <Button
          onClick={() => void form.handleSubmit()}
          variant="contained"
        >
          {buttonText}
        </Button>
      </DialogActions>
    </Dialog>
  );
}
