import { DeleteOutlined, EditOutlined, PlusOutlined } from '@ant-design/icons';
import { Grid } from '@mui/material';
import { Button, List, Modal, Space, Typography } from 'antd';
import { useState } from 'react';
import { useParams } from 'react-router-dom';
import {
  useCreateResult,
  useDeleteResult,
  useFetchResults,
  useUpdateResult,
} from 'api/hooks/resultsHooks';
import { useFetchSeason } from 'api/hooks/seasonsHooks';
import { useFetchTeam, useFetchTeams } from 'api/hooks/teamsHooks';
import { SELECT_OPTIONS } from 'constants/selectOptions';
import withAuth from 'hocs/withAuth';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { ResultDto } from 'types/dtos/responses/results/resultDto';
import { getStartOfDay } from 'utils/dateUtils';
import { resultDtoValidator } from 'validators/results/resultDtoValidator';

function ResultsPage() {
  const { seasonId = '', teamId = '' } = useParams<{
    seasonId: string;
    teamId: string;
  }>();

  const { user } = useAuth();

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingResult, setEditingResult] = useState<Partial<ResultDto>>({});

  const {
    data: results,
    isLoading,
    isError,
  } = useFetchResults(seasonId, teamId);
  const { data: season } = useFetchSeason(seasonId);
  const { data: teams } = useFetchTeams(seasonId);
  const { data: team } = useFetchTeam(seasonId, teamId);

  const { mutateAsync: createResult } = useCreateResult(seasonId, teamId);
  const { mutateAsync: updateResult } = useUpdateResult(seasonId, teamId);
  const { mutateAsync: deleteResult } = useDeleteResult(seasonId, teamId);

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
      if (isEditing && editingResult.id) {
        await updateResult({
          resultId: editingResult.id,
          payload: value,
        });
      } else {
        await createResult(value);
      }
      closeModal();
    },
  });

  const openCreateModal = () => {
    setIsEditing(false);
    form.reset();
    setIsModalOpen(true);
  };

  const openEditModal = (result: ResultDto) => {
    setIsEditing(true);
    setEditingResult(result);
    form.setFieldValue('startDate', result.startDate);
    form.setFieldValue('endDate', result.endDate);
    form.setFieldValue('score', result.score);
    form.setFieldValue('opponentScore', result.opponentScore);
    form.setFieldValue('opponentTeamId', result.opponentTeam.id);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    form.reset();
    setEditingResult({});
  };

  const handleDelete = async (resultId: string) => {
    await deleteResult(resultId);
  };

  if (isError) {
    return (
      <div style={{ padding: 20 }}>
        <Typography.Title level={4}>Results not found</Typography.Title>
      </div>
    );
  }

  return (
    <div style={{ padding: 20, width: '50%', margin: 'auto' }}>
      <Space
        style={{
          marginBottom: 16,
          display: 'flex',
          justifyContent: 'space-between',
        }}
      >
        <Typography.Title level={4}>
          Results for Team &quot;{team?.name}&quot; in Season &quot;
          {season?.name}&quot;
        </Typography.Title>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={openCreateModal}
          disabled={user === null}
          style={{
            filter: user === null ? 'blur(1px)' : 'none',
            cursor: user === null ? 'not-allowed' : 'pointer',
          }}
        >
          Create Result
        </Button>
      </Space>

      <List
        loading={isLoading}
        bordered
        dataSource={results}
        renderItem={(result) => (
          <List.Item
            actions={[
              user &&
              (user.role === 'moderator' ||
                user.role === 'admin' ||
                result.user.id === user.id) ? (
                <EditOutlined
                  key="edit"
                  onClick={() => {
                    openEditModal(result);
                  }}
                />
              ) : null,

              user && (user.role === 'admin' || result.user.id === user.id) ? (
                <DeleteOutlined
                  key="delete"
                  onClick={() => void handleDelete(result.id)}
                  style={{ color: 'red' }}
                />
              ) : null,
            ]}
          >
            <List.Item.Meta
              title={`${result.team.name} vs ${result.opponentTeam.name}`}
              description={
                <>
                  <Typography.Text>
                    {result.score} - {result.opponentScore}
                  </Typography.Text>
                  <br />
                  <Typography.Text type="secondary">
                    {new Date(result.startDate).toLocaleDateString()}
                  </Typography.Text>
                  <br />
                  <Typography.Text type="secondary">
                    {`By: ${result.user.username}`}
                  </Typography.Text>
                </>
              }
            />
          </List.Item>
        )}
      />

      <Modal
        title={isEditing ? 'Edit Result' : 'Create Result'}
        open={isModalOpen}
        onCancel={closeModal}
        onOk={() => void form.handleSubmit()}
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

        <form.AppField name="score">
          {(field) => <field.Text label="Score" />}
        </form.AppField>

        <form.AppField name="opponentScore">
          {(field) => <field.Text label="Opponent Score" />}
        </form.AppField>

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
      </Modal>
    </div>
  );
}

export default withAuth(withErrorBoundary(ResultsPage), {
  isLoggedInOnly: true,
});
