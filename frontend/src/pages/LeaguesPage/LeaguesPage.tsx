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
  useCreateLeague,
  useDeleteLeague,
  useFetchLeagues,
  useUpdateLeague,
} from 'api/hooks/leaguesHooks';
import { StatusHandler } from 'components/StatusHandler';
import { getTeamsLink } from 'constants/route_paths';
import withAuth from 'hocs/withAuth';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { LeagueDto } from 'types/dtos/responses/leagues/leagueDto';
import { getStartOfDay } from 'utils/dateUtils';
import { leagueDtoValidator } from 'validators/leagues/leagueDtoValidator';

function LeaguesPage() {
  const { user } = useAuth();
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingLeagueId, setEditingLeagueId] = useState<string | null>(null);

  const {
    data: leagues,
    isLoading: leaguesLoading,
    error: leaguesError,
  } = useFetchLeagues();
  const { mutateAsync: createLeagueMutation } = useCreateLeague();
  const { mutateAsync: updateLeagueMutation } = useUpdateLeague();
  const { mutateAsync: deleteLeagueMutation } = useDeleteLeague();

  const form = useAppForm({
    defaultValues: {
      name: '',
      startDate: getStartOfDay(),
      endDate: getStartOfDay(7),
    },
    validators: {
      onSubmit: leagueDtoValidator,
    },
    onSubmit: async ({ value }) => {
      if (isEditing && editingLeagueId) {
        await updateLeagueMutation({
          leagueId: editingLeagueId,
          payload: value,
        });
      } else {
        await createLeagueMutation(value);
      }
      closeDialog();
    },
  });

  const openEditDialog = (league: LeagueDto) => {
    setIsEditing(true);
    setEditingLeagueId(league.id);
    form.setFieldValue('name', league.name);
    form.setFieldValue('startDate', new Date(league.startDate));
    form.setFieldValue('endDate', new Date(league.endDate));
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
    setEditingLeagueId(null);
  };

  const handleDelete = async (leagueId: string) => {
    await deleteLeagueMutation(leagueId);
  };

  return (
    <Box sx={{ p: 4, maxWidth: 800, mx: 'auto' }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
        <Typography
          variant="h4"
          fontWeight={700}
        >
          Leagues
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
          Create League
        </Button>
      </Box>

      <StatusHandler
        isLoading={leaguesLoading}
        error={leaguesError}
        isEmpty={!leagues || leagues.length === 0}
      >
        <List sx={{ borderRadius: 2, boxShadow: 2, overflow: 'hidden' }}>
          {leagues?.map((league, index) => (
            <>
              <ListItem
                key={league.id}
                secondaryAction={
                  user && (
                    <>
                      {league.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => {
                            openEditDialog(league);
                          }}
                        >
                          <Edit />
                        </IconButton>
                      )}
                      {league.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => void handleDelete(league.id)}
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
                      to={getTeamsLink(league.id)}
                      style={{ textDecoration: 'none', color: 'inherit' }}
                    >
                      <Typography
                        variant="subtitle1"
                        fontWeight={600}
                      >
                        {league.name}
                      </Typography>
                    </Link>
                  }
                  secondary={
                    <>
                      <Typography variant="body2">
                        {`${dayjs(league.startDate).format('YYYY-MM-DD')} - ${dayjs(league.endDate).format('YYYY-MM-DD')}`}
                      </Typography>
                      <Typography
                        variant="caption"
                        color="text.secondary"
                      >
                        {`Created By: ${league.user.username}`}
                      </Typography>
                    </>
                  }
                />
              </ListItem>

              {index < leagues.length - 1 && <Divider />}
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
        <DialogTitle>{isEditing ? 'Edit League' : 'Create League'}</DialogTitle>
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

export default withAuth(withErrorBoundary(LeaguesPage), {
  isLoggedInOnly: true,
});
