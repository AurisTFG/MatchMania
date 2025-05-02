import { Add, Delete, Edit } from '@mui/icons-material';
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  IconButton,
  List,
  ListItem,
  ListItemText,
  Typography,
} from '@mui/material';
import { useState } from 'react';
import { Link, useSearchParams } from 'react-router-dom';
import { useFetchLeague } from 'api/hooks/leaguesHooks';
import {
  useCreateTeam,
  useDeleteTeam,
  useFetchTeams,
  useUpdateTeam,
} from 'api/hooks/teamsHooks';
import { StatusHandler } from 'components/StatusHandler';
import { PARAMS, getResultsLink } from 'constants/route_paths';
import withAuth from 'hocs/withAuth';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { TeamDto } from 'types/dtos/responses/teams/teamDto';
import { createTeamDtoValidator } from 'validators/teams/createTeamDtoValidator';

function TeamsPage() {
  const [searchParams] = useSearchParams();
  const leagueId = searchParams.get(PARAMS.SEASON_ID) ?? '';

  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);

  const [editingTeamId, setEditingTeamId] = useState<string | null>(null);
  const { user } = useAuth();

  const { data: league, isLoading: isLeagueLoading } = useFetchLeague(leagueId);
  const { data: teams, isLoading: isTeamsLoading } = useFetchTeams(leagueId);
  const { mutateAsync: createTeam } = useCreateTeam(leagueId);
  const { mutateAsync: updateTeam } = useUpdateTeam(leagueId);
  const { mutateAsync: deleteTeam } = useDeleteTeam(leagueId);

  const form = useAppForm({
    defaultValues: {
      name: '',
    },
    validators: {
      onSubmit: createTeamDtoValidator,
    },
    onSubmit: async ({ value }) => {
      if (isEditing && editingTeamId) {
        await updateTeam({ teamId: editingTeamId, payload: value });
      } else {
        await createTeam(value);
      }
      closeDialog();
    },
  });

  const openCreateDialog = () => {
    setIsEditing(false);
    form.reset();
    setIsDialogOpen(true);
  };

  const openEditDialog = (team: TeamDto) => {
    setIsEditing(true);
    setEditingTeamId(team.id);
    form.setFieldValue('name', team.name);
    setIsDialogOpen(true);
  };

  const closeDialog = () => {
    setIsDialogOpen(false);
    form.reset();
    setEditingTeamId(null);
  };

  const handleDelete = async (teamId: string) => {
    await deleteTeam(teamId);
  };

  if (isLeagueLoading || isTeamsLoading) {
    return (
      <Box sx={{ p: 4 }}>
        <Typography
          variant="h4"
          fontWeight={700}
        >
          Loading...
        </Typography>
      </Box>
    );
  }

  return (
    <Box sx={{ p: 4, maxWidth: 800, mx: 'auto' }}>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 3 }}>
        <Typography
          variant="h4"
          fontWeight={700}
        >
          Teams for League &quot;{league?.name}&quot;
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
          Create Team
        </Button>
      </Box>

      <StatusHandler
        isLoading={isTeamsLoading}
        error={undefined}
        isEmpty={!teams || teams.length === 0}
      >
        <List sx={{ borderRadius: 2, boxShadow: 2, overflow: 'hidden' }}>
          {teams?.map((team, index) => (
            <Box key={team.id}>
              <ListItem
                secondaryAction={
                  user && (
                    <>
                      {team.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => {
                            openEditDialog(team);
                          }}
                        >
                          <Edit />
                        </IconButton>
                      )}
                      {team.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => void handleDelete(team.id)}
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
                      to={getResultsLink(leagueId, team.id)}
                      style={{ textDecoration: 'none', color: 'inherit' }}
                    >
                      <Typography
                        variant="subtitle1"
                        fontWeight={600}
                      >
                        {team.name}
                      </Typography>
                    </Link>
                  }
                  secondary={
                    <>
                      <Typography variant="body2">
                        {`Elo: ${String(team.elo)}`}
                      </Typography>
                      <Typography
                        variant="caption"
                        color="text.secondary"
                      >
                        {`Created By: ${team.user.username}`}
                      </Typography>
                    </>
                  }
                />
              </ListItem>

              {index < teams.length - 1 && <Divider />}
            </Box>
          ))}
        </List>
      </StatusHandler>

      <Dialog
        open={isDialogOpen}
        onClose={closeDialog}
        fullWidth
        maxWidth="sm"
      >
        <DialogTitle>{isEditing ? 'Edit Team' : 'Create Team'}</DialogTitle>
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

export default withAuth(withErrorBoundary(TeamsPage), {
  isLoggedInOnly: true,
});
