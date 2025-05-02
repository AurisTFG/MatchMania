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
import { useSearchParams } from 'react-router-dom';
import { useFetchLeague } from 'api/hooks/leaguesHooks';
import {
  useCreateResult,
  useDeleteResult,
  useFetchResults,
  useUpdateResult,
} from 'api/hooks/resultsHooks';
import { useFetchTeam, useFetchTeams } from 'api/hooks/teamsHooks';
import { StatusHandler } from 'components/StatusHandler';
import { PARAMS } from 'constants/route_paths';
import { SELECT_OPTIONS } from 'constants/selectOptions';
import withAuth from 'hocs/withAuth';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { ResultDto } from 'types/dtos/responses/results/resultDto';
import { getStartOfDay } from 'utils/dateUtils';
import { resultDtoValidator } from 'validators/results/resultDtoValidator';

function ResultsPage() {
  const [searchParams] = useSearchParams();
  const leagueId = searchParams.get(PARAMS.SEASON_ID) ?? '';
  const teamId = searchParams.get(PARAMS.TEAM_ID) ?? '';

  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingResultId, setEditingResultId] = useState<string | null>(null);

  const { user } = useAuth();

  const {
    data: results,
    isLoading,
    isError,
  } = useFetchResults(leagueId, teamId);
  const { data: league } = useFetchLeague(leagueId);
  const { data: teams } = useFetchTeams(leagueId);
  const { data: team } = useFetchTeam(leagueId, teamId);

  const { mutateAsync: createResult } = useCreateResult(leagueId, teamId);
  const { mutateAsync: updateResult } = useUpdateResult(leagueId, teamId);
  const { mutateAsync: deleteResult } = useDeleteResult(leagueId, teamId);

  const form = useAppForm({
    defaultValues: {
      startDate: getStartOfDay(),
      endDate: getStartOfDay(),
      score: '',
      opponentScore: '',
      opponentTeamId: SELECT_OPTIONS.NOT_SELECTED.key,
    },
    validators: {
      onSubmit: resultDtoValidator,
    },
    onSubmit: async ({ value }) => {
      if (isEditing && editingResultId) {
        await updateResult({ resultId: editingResultId, payload: value });
      } else {
        await createResult(value);
      }
      closeDialog();
    },
  });

  const openCreateDialog = () => {
    setIsEditing(false);
    form.reset();
    setIsDialogOpen(true);
  };

  const openEditDialog = (result: ResultDto) => {
    setIsEditing(true);
    setEditingResultId(result.id);
    form.setFieldValue('startDate', new Date(result.startDate));
    form.setFieldValue('endDate', new Date(result.endDate));
    form.setFieldValue('score', result.score);
    form.setFieldValue('opponentScore', result.opponentScore);
    form.setFieldValue('opponentTeamId', result.opponentTeam.id);
    setIsDialogOpen(true);
  };

  const closeDialog = () => {
    setIsDialogOpen(false);
    form.reset();
    setEditingResultId(null);
  };

  const handleDelete = async (resultId: string) => {
    await deleteResult(resultId);
  };

  if (isError) {
    return (
      <Box sx={{ p: 4 }}>
        <Typography
          variant="h4"
          fontWeight={700}
        >
          Results not found
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
          Results for Team &quot;{team?.name}&quot; in League &quot;
          {league?.name}&quot;
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
          Create Result
        </Button>
      </Box>

      <StatusHandler
        isLoading={isLoading}
        error={undefined}
        isEmpty={!results || results.length === 0}
      >
        <List sx={{ borderRadius: 2, boxShadow: 2, overflow: 'hidden' }}>
          {results?.map((result, index) => (
            <Box key={result.id}>
              <ListItem
                secondaryAction={
                  user && (
                    <>
                      {result.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => {
                            openEditDialog(result);
                          }}
                        >
                          <Edit />
                        </IconButton>
                      )}
                      {result.user.id === user.id && (
                        <IconButton
                          edge="end"
                          onClick={() => void handleDelete(result.id)}
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
                    <Typography
                      variant="subtitle1"
                      fontWeight={600}
                    >
                      {`${result.team.name} vs ${result.opponentTeam.name}`}
                    </Typography>
                  }
                  secondary={
                    <>
                      <Typography variant="body2">
                        {`${result.score} - ${result.opponentScore}`}
                      </Typography>
                      <Typography
                        variant="caption"
                        color="text.secondary"
                      >
                        {dayjs(result.startDate).format('YYYY-MM-DD')}
                      </Typography>
                      <br />
                      <Typography
                        variant="caption"
                        color="text.secondary"
                      >
                        {`By: ${result.user.username}`}
                      </Typography>
                    </>
                  }
                />
              </ListItem>

              {index < results.length - 1 && <Divider />}
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
        <DialogTitle>{isEditing ? 'Edit Result' : 'Create Result'}</DialogTitle>
        <DialogContent dividers>
          <Box
            component="form"
            noValidate
            autoComplete="off"
            sx={{ mt: 2 }}
          >
            <Grid
              container
              spacing={2}
            >
              <Grid size={6}>
                <form.AppField name="startDate">
                  {(field) => <field.DatePicker label="Match Start Date" />}
                </form.AppField>
              </Grid>
              <Grid size={6}>
                <form.AppField name="endDate">
                  {(field) => <field.DatePicker label="Match End Date" />}
                </form.AppField>
              </Grid>
            </Grid>

            <Box sx={{ mt: 2 }}>
              <form.AppField name="score">
                {(field) => <field.Text label="Score" />}
              </form.AppField>
            </Box>

            <Box sx={{ mt: 2 }}>
              <form.AppField name="opponentScore">
                {(field) => <field.Text label="Opponent Score" />}
              </form.AppField>
            </Box>

            <Box sx={{ mt: 2 }}>
              <form.AppField name="opponentTeamId">
                {(field) => (
                  <field.Select
                    label="Select Opponent Team"
                    options={(teams ?? [])
                      .filter((opponentTeam) => opponentTeam.id !== teamId)
                      .map((opponentTeam) => ({
                        key: opponentTeam.id,
                        value: opponentTeam.name,
                      }))}
                  />
                )}
              </form.AppField>
            </Box>
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

export default withAuth(withErrorBoundary(ResultsPage), {
  isLoggedInOnly: true,
});
