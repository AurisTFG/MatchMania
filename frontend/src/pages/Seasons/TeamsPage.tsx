import { DeleteOutlined, EditOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Input, List, Modal, Space, Typography, message } from 'antd';
import { useEffect, useState } from 'react';
import { Link, useParams } from 'react-router-dom';
import { useFetchSeason } from '../../api/hooks/seasonsHooks';
import {
  useCreateTeam,
  useDeleteTeam,
  useFetchTeams,
  useUpdateTeam,
} from '../../api/hooks/teamsHooks';
import { useFetchUsers } from '../../api/hooks/usersHooks';
import { useAuth } from '../../providers/AuthProvider';
import { Team, User } from '../../types';

const isValidTeam = (seasonId: string | undefined) => {
  return seasonId && !isNaN(Number(seasonId)) && Number(seasonId) > 0;
};

export default function TeamsPage() {
  const { seasonId } = useParams<{ seasonId: string }>();
  const [teamsNotFound, setTeamsNotFound] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingTeam, setEditingTeam] = useState<Partial<Team>>({});
  const [formData, setFormData] = useState({
    name: '',
  });
  const { user } = useAuth();

  const seasonIdNumber = seasonId ? parseInt(seasonId) : 0;

  const { data: season, isLoading: isSeasonsLoading } =
    useFetchSeason(seasonIdNumber);
  const { data: teams, isLoading: isTeamsLoading } =
    useFetchTeams(seasonIdNumber);
  const { data: users, isLoading: isUsersLoading } = useFetchUsers();
  const { mutateAsync: createTeam } = useCreateTeam(seasonIdNumber);
  const { mutateAsync: updateTeam } = useUpdateTeam(seasonIdNumber);
  const { mutateAsync: deleteTeam } = useDeleteTeam(seasonIdNumber);

  const getUserById = (userId: string): User => {
    return (
      users?.find((user) => user.id === userId) ??
      ({
        id: '',
        username: 'Unknown User',
      } as User)
    );
  };

  useEffect(() => {
    if (!isValidTeam(seasonId)) {
      setTeamsNotFound(true);
      return;
    }
  }, [seasonId]);

  if (isSeasonsLoading || isTeamsLoading || isUsersLoading) {
    return (
      <div style={{ padding: 20 }}>
        <Typography.Title level={4}>Loading...</Typography.Title>
      </div>
    );
  }

  if (teamsNotFound) {
    return (
      <div style={{ padding: 20 }}>
        <Typography.Title level={4}>Teams not found</Typography.Title>
      </div>
    );
  }

  const openModal = (team?: Team) => {
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
    try {
      if (isEditing && editingTeam.id) {
        await updateTeam({ teamID: editingTeam.id, team: formData });
        message.success('Team updated successfully.');
      } else {
        await createTeam(formData);
        message.success('Team created successfully.');
      }
      setIsModalOpen(false);
      setFormData({ name: '' });
    } catch (error) {
      message.error('Failed to save team.');
      console.error(error);
    }
  };

  const handleDelete = async (teamID: number) => {
    try {
      await deleteTeam(teamID);
      message.success('Team deleted successfully.');
    } catch (error) {
      message.error('Failed to delete team.');
      console.error(error);
    }
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
                team.userUUID === user.id) ? (
                <EditOutlined
                  key="edit"
                  onClick={() => {
                    openModal(team);
                  }}
                />
              ) : null,

              user && (user.role === 'admin' || team.userUUID === user.id) ? (
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
                <Link
                  to={`/seasons/${seasonId ?? ''}/teams/${String(team.id)}/results`}
                >
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
                    {`By: ${getUserById(team.userUUID).username}`}
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
