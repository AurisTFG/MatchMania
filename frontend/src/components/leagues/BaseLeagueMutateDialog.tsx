import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Grid,
} from '@mui/material';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { CreateLeagueDto } from 'types/dtos/requests/leagues/createLeagueDto';
import { leagueDtoValidator } from 'validators/leagues/leagueDtoValidator';

type BaseLeagueMutateDialogProps = {
  title: string;
  buttonText: string;
  league: CreateLeagueDto;
  submitAsync: (payload: CreateLeagueDto) => Promise<void>;
  open: boolean;
  onClose: () => void;
};

export default function BaseLeagueMutateDialog({
  title,
  buttonText,
  league,
  submitAsync,
  open,
  onClose,
}: BaseLeagueMutateDialogProps) {
  const { user } = useAuth();

  const form = useAppForm({
    defaultValues: league,
    validators: {
      onSubmit: leagueDtoValidator,
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

  const tracksOptions =
    user?.tracks.map((track) => ({
      key: track.id,
      value: track.name,
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
            {(field) => <field.Text label="League Name" />}
          </form.AppField>

          <form.AppField name="trackIds">
            {(field) => (
              <field.MultiSelect
                label="Select Tracks"
                options={tracksOptions}
                renderOptionText={(trackOption) => (
                  <span
                    key={trackOption.key}
                    dangerouslySetInnerHTML={{
                      __html: MPStyle.Parser.toHTML(trackOption.value),
                    }}
                  />
                )}
              />
            )}
          </form.AppField>

          <Grid
            container
            spacing={2}
            sx={{ mt: 2 }}
          >
            <Grid size={6}>
              <form.AppField name="startDate">
                {(field) => <field.DatePicker label="Start Date" />}
              </form.AppField>
            </Grid>
            <Grid size={6}>
              <form.AppField name="endDate">
                {(field) => <field.DatePicker label="End Date" />}
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
