import { Add, Delete, Edit } from '@mui/icons-material';
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  Grid,
  IconButton,
  List,
  ListItem,
  ListItemText,
  Typography,
} from '@mui/material';
import dayjs from 'dayjs';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import {
  useCreateSeason,
  useDeleteSeason,
  useFetchSeasons,
  useUpdateSeason,
} from 'api/hooks/seasonsHooks';
import { StatusHandler } from 'components/StatusHandler';
import { getTeamsLink } from 'constants/route_paths';
import withAuth from 'hocs/withAuth';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { SeasonDto } from 'types/dtos/responses/seasons/seasonDto';
import { getStartOfDay } from 'utils/dateUtils';
import { seasonDtoValidator } from 'validators/seasons/seasonDtoValidator';

function SeasonsPage() {
  const { user } = useAuth();
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingSeasonId, setEditingSeasonId] = useState<string | null>(null);

  const {
    data: seasons,
    isLoading: seasonsLoading,
    error: seasonsError,
  } = useFetchSeasons();
  const { mutateAsync: createSeasonMutation } = useCreateSeason();
  const { mutateAsync: updateSeasonMutation } = useUpdateSeason();
  const { mutateAsync: deleteSeasonMutation } = useDeleteSeason();

  const form = useAppForm({
    defaultValues: {
      name: '',
      startDate: getStartOfDay(),
      endDate: getStartOfDay(7),
    },
    validators: {
      onSubmit: seasonDtoValidator,
    },
    onSubmit: async ({ value }) => {
      if (isEditing && editingSeasonId) {
        await updateSeasonMutation({
          seasonId: editingSeasonId,
          payload: value,
        });
      } else {
        await createSeasonMutation(value);
      }
      closeDialog();
    },
  });

  const openEditDialog = (season: SeasonDto) => {
    setIsEditing(true);
    setEditingSeasonId(season.id);
    form.setFieldValue('name', season.name);
    form.setFieldValue('startDate', new Date(season.startDate));
    form.setFieldValue('endDate', new Date(season.endDate));
    setIsDialogOpen(true);
  };

  const openCreateDialog = () => {
    setIsEditing(false);
    form.reset();
    setIsDialogOpen(true);
  };

  const closeDialog = () => {
    setIsDialogOpen(false);
    form.reset();
    setEditingSeasonId(null);
  };

  const handleDelete = async (seasonId: string) => {
    await deleteSeasonMutation(seasonId);
  };

  return (
    <Box sx={{ p: 4, maxWidth: 800, mx: 'auto' }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
        <Typography
          variant="h4"
          fontWeight={700}
        >
          Seasons
        </Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={openCreateDialog}
          disabled={!user}
          sx={{
            filter: !user ? 'blur(1px)' : 'none',
            cursor: !user ? 'not-allowed' : 'pointer',
          }}
        >
          Create Season
        </Button>
      </Box>

      <StatusHandler
        isLoading={seasonsLoading}
        error={seasonsError}
        isEmpty={!seasons || seasons.length === 0}
      >
        <List sx={{ borderRadius: 2, boxShadow: 2, overflow: 'hidden' }}>
          {seasons?.map((season, index) => (
            <>
              <ListItem
                key={season.id}
                secondaryAction={
                  user && (
                    <>
                      {season.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => {
                            openEditDialog(season);
                          }}
                        >
                          <Edit />
                        </IconButton>
                      )}
                      {season.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => void handleDelete(season.id)}
                          sx={{ color: 'error.main' }}
                        >
                          <Delete />
                        </IconButton>
                      )}
                    </>
                  )
                }
                sx={{
                  borderRadius: 2,
                  px: 2,
                  py: 1,
                  my: 1,
                  backgroundColor: 'background.paper',
                  transition: 'background-color 0.2s ease',
                  '&:hover': {
                    backgroundColor: 'action.hover',
                  },
                }}
              >
                <ListItemText
                  primary={
                    <Link
                      to={getTeamsLink(season.id)}
                      style={{ textDecoration: 'none', color: 'inherit' }}
                    >
                      <Typography
                        variant="subtitle1"
                        fontWeight={600}
                      >
                        {season.name}
                      </Typography>
                    </Link>
                  }
                  secondary={
                    <>
                      <Typography variant="body2">
                        {`${dayjs(season.startDate).format('YYYY-MM-DD')} - ${dayjs(season.endDate).format('YYYY-MM-DD')}`}
                      </Typography>
                      <Typography
                        variant="caption"
                        color="text.secondary"
                      >
                        {`Created By: ${season.user.username}`}
                      </Typography>
                    </>
                  }
                />
              </ListItem>

              {index < seasons.length - 1 && <Divider />}
            </>
          ))}
        </List>
      </StatusHandler>

      <Dialog
        open={isDialogOpen}
        onClose={closeDialog}
        fullWidth
        maxWidth="sm"
      >
        <DialogTitle>{isEditing ? 'Edit Season' : 'Create Season'}</DialogTitle>
        <DialogContent dividers>
          <Box
            component="form"
            noValidate
            autoComplete="off"
            sx={{ mt: 2 }}
          >
            <form.AppField name="name">
              {(field) => <field.Text label="Season Name" />}
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
          <Button onClick={closeDialog}>Cancel</Button>
          <Button
            onClick={() => void form.handleSubmit()}
            variant="contained"
          >
            {isEditing ? 'Update' : 'Create'}
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}

export default withAuth(withErrorBoundary(SeasonsPage), {
  isLoggedInOnly: true,
});
