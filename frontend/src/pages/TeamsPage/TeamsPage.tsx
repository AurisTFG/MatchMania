import { DeleteOutlined, EditOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, List, Modal, Space, Typography } from 'antd';
import { useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { useFetchSeason } from 'api/hooks/seasonsHooks';
import {
  useCreateTeam,
  useDeleteTeam,
  useFetchTeams,
  useUpdateTeam,
} from 'api/hooks/teamsHooks';
import withAuth from 'hocs/withAuth';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider/AuthProvider';
import { TeamDto } from 'types/dtos/responses/teams/teamDto';
import { createTeamDtoValidator } from 'validators/teams/createTeamDtoValidator';

function TeamsPage() {
  const { seasonId = '' } = useParams<{ seasonId: string }>();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingTeam, setEditingTeam] = useState<Partial<TeamDto>>({});
  const { user } = useAuth();

  const { data: season, isLoading: isSeasonsLoading } =
    useFetchSeason(seasonId);
  const { data: teams, isLoading: isTeamsLoading } = useFetchTeams(seasonId);
  const { mutateAsync: createTeam } = useCreateTeam(seasonId);
  const { mutateAsync: updateTeam } = useUpdateTeam(seasonId);
  const { mutateAsync: deleteTeam } = useDeleteTeam(seasonId);

  const form = useAppForm({
    defaultValues: {
      name: '',
    },
    validators: {
      onSubmit: createTeamDtoValidator,
    },
    onSubmit: async ({ value }) => {
      if (isEditing && editingTeam.id) {
        await updateTeam({ teamId: editingTeam.id, payload: value });
      } else {
        await createTeam(value);
      }

      closeModal();
    },
  });

  if (isSeasonsLoading || isTeamsLoading) {
    return (
      <div style={{ padding: 20 }}>
        <Typography.Title level={4}>Loading...</Typography.Title>
      </div>
    );
  }

  const openModal = (team?: TeamDto) => {
    if (team) {
      setIsEditing(true);
      setEditingTeam(team);
      form.setFieldValue('name', team.name);
    } else {
      setIsEditing(false);
      form.reset();
    }
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    form.reset();
    setEditingTeam({});
  };

  const handleDelete = async (teamId: string) => {
    await deleteTeam(teamId);
  };

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
          Teams for Season &quot;{season?.name}&quot;
        </Typography.Title>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => {
            openModal();
          }}
          disabled={user === null}
          style={{
            filter: user === null ? 'blur(1px)' : 'none',
            cursor: user === null ? 'not-allowed' : 'pointer',
          }}
        >
          Create Team
        </Button>
      </Space>

      <List
        loading={isTeamsLoading}
        bordered
        dataSource={teams}
        renderItem={(team) => (
          <List.Item
            actions={[
              user &&
              (user.role === 'moderator' ||
                user.role === 'admin' ||
                team.user.id === user.id) ? (
                <EditOutlined
                  key="edit"
                  onClick={() => {
                    openModal(team);
                  }}
                />
              ) : null,

              user && (user.role === 'admin' || team.user.id === user.id) ? (
                <DeleteOutlined
                  key="delete"
                  onClick={() => void handleDelete(team.id)}
                  style={{ color: 'red' }}
                />
              ) : null,
            ]}
          >
            <List.Item.Meta
              title={
                <Link to={`/seasons/${seasonId}/teams/${team.id}/results`}>
                  {team.name}
                </Link>
              }
              description={
                <>
                  <Typography.Text type="secondary">
                    {`Elo: ${team.elo.toString()}`}
                  </Typography.Text>
                  <br />
                  <Typography.Text type="secondary">
                    {`By: ${team.user.username}`}
                  </Typography.Text>
                </>
              }
            />
          </List.Item>
        )}
      />

      <Modal
        title={isEditing ? 'Edit Team' : 'Create Team'}
        open={isModalOpen}
        onCancel={closeModal}
        onOk={() => void form.handleSubmit()}
      >
        <form.AppField name="name">
          {(field) => <field.Text label="Team Name" />}
        </form.AppField>
      </Modal>
    </div>
  );
}

export default withAuth(withErrorBoundary(TeamsPage), {
  isLoggedInOnly: true,
});
