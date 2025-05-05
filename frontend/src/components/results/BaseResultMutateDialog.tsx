import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Grid,
} from '@mui/material';
import { useFetchLeagues } from 'api/hooks/leaguesHooks';
import { useFetchTeams } from 'api/hooks/teamsHooks';
import { SELECT_OPTIONS } from 'constants/selectOptions';
import { useAppForm } from 'hooks/form/useAppForm';
import { CreateResultDto } from 'types/dtos/requests/results/createResultDto';
import { resultDtoValidator } from 'validators/results/resultDtoValidator';

type BaseResultMutateDialogProps = {
  title: string;
  buttonText: string;
  result: CreateResultDto;
  submitAsync: (payload: CreateResultDto) => Promise<void>;
  open: boolean;
  onClose: () => void;
};

export default function BaseResultMutateDialog({
  title,
  buttonText,
  result,
  submitAsync,
  open,
  onClose,
}: BaseResultMutateDialogProps) {
  const { data: teams } = useFetchTeams();
  const { data: leagues } = useFetchLeagues();

  const form = useAppForm({
    defaultValues: result,
    validators: {
      onSubmit: resultDtoValidator,
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

  const teamsOptions = (teams ?? []).map((team) => ({
    key: team.id,
    value: team.name,
  }));

  const leaguesOptions = (leagues ?? []).map((league) => ({
    key: league.id,
    value: league.name,
  }));

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
          <form.AppField name="leagueId">
            {(field) => (
              <field.Select
                label="Select League"
                options={leaguesOptions}
                notSelectedOption={SELECT_OPTIONS.NOT_SELECTED}
              />
            )}
          </form.AppField>
          <Grid
            container
            spacing={2}
          >
            <Grid size={6}>
              <form.AppField name="startDate">
                {(field) => <field.DateTimePicker label="Match Start Date" />}
              </form.AppField>
            </Grid>
            <Grid size={6}>
              <form.AppField name="endDate">
                {(field) => <field.DateTimePicker label="Match End Date" />}
              </form.AppField>
            </Grid>
          </Grid>

          <Grid
            container
            spacing={2}
          >
            <Grid size={6}>
              <form.AppField name="teamId">
                {(field) => (
                  <field.Select
                    label="Select Team"
                    options={teamsOptions}
                    notSelectedOption={SELECT_OPTIONS.NOT_SELECTED}
                  />
                )}
              </form.AppField>
            </Grid>
            <Grid size={6}>
              <form.AppField name="opponentTeamId">
                {(field) => (
                  <field.Select
                    label="Select Opponent Team"
                    options={teamsOptions}
                    notSelectedOption={SELECT_OPTIONS.NOT_SELECTED}
                  />
                )}
              </form.AppField>
            </Grid>
          </Grid>

          <Grid
            container
            spacing={2}
          >
            <Grid size={6}>
              <form.AppField name="score">
                {(field) => <field.Text label="Team Score" />}
              </form.AppField>
            </Grid>
            <Grid size={6}>
              <form.AppField name="opponentScore">
                {(field) => <field.Text label="Opponent Team Score" />}
              </form.AppField>
            </Grid>
          </Grid>
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
