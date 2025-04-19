import { DeleteOutlined, EditOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Input, List, Modal, Space, Typography } from 'antd';
import { useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { useFetchSeason } from '../../api/hooks/seasonsHooks';
import {
  useCreateTeam,
  useDeleteTeam,
  useFetchTeams,
  useUpdateTeam,
} from '../../api/hooks/teamsHooks';
import { useFetchUsers } from '../../api/hooks/usersHooks';
import { useAuth } from '../../providers/AuthProvider/AuthProvider';
import { TeamDto } from '../../types/dtos/responses/teams/teamDto';
import { UserMinimalDto } from '../../types/dtos/responses/users/userMinimalDto';

export default function TeamsPage() {
  const { seasonId = '' } = useParams<{ seasonId: string }>();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingTeam, setEditingTeam] = useState<Partial<TeamDto>>({});
  const [formData, setFormData] = useState({
    name: '',
  });
  const { user } = useAuth();

  const { data: season, isLoading: isSeasonsLoading } =
    useFetchSeason(seasonId);
  const { data: teams, isLoading: isTeamsLoading } = useFetchTeams(seasonId);
  const { data: users, isLoading: isUsersLoading } = useFetchUsers();
  const { mutateAsync: createTeam } = useCreateTeam(seasonId);
  const { mutateAsync: updateTeam } = useUpdateTeam(seasonId);
  const { mutateAsync: deleteTeam } = useDeleteTeam(seasonId);

  const getUserById = (userId: string): UserMinimalDto => {
    return (
      users?.find((user) => user.id === userId) ??
      ({
        id: '',
        username: 'Unknown User',
      } as UserMinimalDto)
    );
  };

  if (isSeasonsLoading || isTeamsLoading || isUsersLoading) {
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
      setFormData({ name: team.name });
    } else {
      setIsEditing(false);
      setFormData({ name: '' });
    }
    setIsModalOpen(true);
  };

  const handleCreateOrEdit = async () => {
    if (isEditing && editingTeam.id) {
      await updateTeam({ teamId: editingTeam.id, payload: formData });
    } else {
      await createTeam(formData);
    }
    setIsModalOpen(false);
    setFormData({ name: '' });
  };

  const handleDelete = async (teamId: string) => {
    await deleteTeam(teamId);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setFormData({ name: '' });
    setEditingTeam({});
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
                team.userId === user.id) ? (
                <EditOutlined
                  key="edit"
                  onClick={() => {
                    openModal(team);
                  }}
                />
              ) : null,

              user && (user.role === 'admin' || team.userId === user.id) ? (
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
                    {`By: ${getUserById(team.userId).username}`}
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
        // eslint-disable-next-line @typescript-eslint/no-misused-promises
        onOk={handleCreateOrEdit}
      >
        <Input
          placeholder="Team Name"
          value={formData.name}
          onChange={(e) => {
            setFormData({ ...formData, name: e.target.value });
          }}
          style={{ marginBottom: 8 }}
        />
      </Modal>
    </div>
  );
}
