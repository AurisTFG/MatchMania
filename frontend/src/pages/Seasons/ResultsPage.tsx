import { DeleteOutlined, EditOutlined, PlusOutlined } from '@ant-design/icons';
import {
  Button,
  DatePicker,
  Input,
  List,
  Modal,
  Select,
  Space,
  Typography,
  message,
} from 'antd';
import moment from 'moment';
import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import {
  useCreateResult,
  useDeleteResult,
  useFetchResults,
  useUpdateResult,
} from '../../api/hooks/resultsHooks.ts';
import { useFetchSeason } from '../../api/hooks/seasonsHooks.ts';
import { useFetchTeam, useFetchTeams } from '../../api/hooks/teamsHooks.ts';
import { useFetchUsers } from '../../api/hooks/usersHooks.ts';
import { useAuth } from '../../providers/AuthProvider.tsx';
import { Result, User } from '../../types';

const { Option } = Select;

const isValidResult = (
  seasonId: string | undefined,
  teamId: string | undefined,
) => {
  return (
    seasonId &&
    !isNaN(Number(seasonId)) &&
    Number(seasonId) > 0 &&
    teamId &&
    !isNaN(Number(teamId)) &&
    Number(teamId) > 0
  );
};

export default function ResultsPage() {
  const { seasonId, teamId } = useParams<{
    seasonId: string;
    teamId: string;
  }>();

  const seasonIdInteger = seasonId ? parseInt(seasonId) : 0;
  const teamIdInteger = teamId ? parseInt(teamId) : 0;

  const { user } = useAuth();

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingResult, setEditingResult] = useState<Partial<Result>>({});
  const [formData, setFormData] = useState({
    matchStartDate: moment(),
    matchEndDate: moment(),
    score: '',
    opponentScore: '',
    opponentTeamId: 0,
  });

  const {
    data: results,
    isLoading,
    isError,
  } = useFetchResults(seasonIdInteger, teamIdInteger);
  const { data: season } = useFetchSeason(seasonIdInteger);
  const { data: team } = useFetchTeam(seasonIdInteger, teamIdInteger);
  const { data: teams } = useFetchTeams(seasonIdInteger);
  const { data: users } = useFetchUsers();

  const { mutateAsync: createResult } = useCreateResult(
    seasonIdInteger,
    teamIdInteger,
  );
  const { mutateAsync: updateResult } = useUpdateResult(
    seasonIdInteger,
    teamIdInteger,
  );
  const { mutateAsync: deleteResult } = useDeleteResult(
    seasonIdInteger,
    teamIdInteger,
  );

  useEffect(() => {
    if (!isValidResult(seasonId, teamId)) {
      return;
    }
  }, [seasonId, teamId]);

  const getUserById = (userId: string): User => {
    return (
      users?.find((user) => user.id === userId) ??
      ({
        id: '',
        username: 'Unknown User',
      } as User)
    );
  };

  const openCreateModal = () => {
    setIsEditing(false);
    setEditingResult({});
    setFormData({
      matchStartDate: moment(),
      matchEndDate: moment(),
      score: '',
      opponentScore: '',
      opponentTeamId: 0,
    });
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setEditingResult({});
    setFormData({
      matchStartDate: moment(),
      matchEndDate: moment(),
      score: '',
      opponentScore: '',
      opponentTeamId: 0,
    });
  };

  const openEditModal = (result: Result) => {
    setIsEditing(true);
    setEditingResult(result);
    setFormData({
      matchStartDate: moment(result.matchStartDate),
      matchEndDate: moment(result.matchEndDate),
      score: result.score,
      opponentScore: result.opponentScore,
      opponentTeamId: result.opponentTeamId,
    });
    setIsModalOpen(true);
  };

  const handleCreateOrEdit = async () => {
    try {
      if (isEditing && editingResult.id) {
        await updateResult({
          resultID: editingResult.id,
          result: {
            matchStartDate: formData.matchStartDate.toDate(),
            matchEndDate: formData.matchEndDate.toDate(),
            score: formData.score,
            opponentScore: formData.opponentScore,
          },
        });
        message.success('Result updated successfully.');
      } else {
        await createResult({
          matchStartDate: formData.matchStartDate.toDate(),
          matchEndDate: formData.matchEndDate.toDate(),
          score: formData.score,
          opponentScore: formData.opponentScore,
          opponentTeamId: formData.opponentTeamId,
        });
        message.success('Result created successfully.');
      }
      closeModal();
    } catch (error) {
      message.error('Failed to save result.');
      console.error(error);
    }
  };

  const handleDelete = async (resultID: number) => {
    try {
      await deleteResult(resultID);
      message.success('Result deleted successfully.');
    } catch (error) {
      message.error('Failed to delete result.');
      console.error(error);
    }
  };

  const getTeamName = (id: number) => {
    const team = teams?.find((team) => team.id === id);
    return team ? team.name : 'Unknown Team';
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
                result.userUUID === user.id) ? (
                <EditOutlined
                  key="edit"
                  onClick={() => {
                    openEditModal(result);
                  }}
                />
              ) : null,

              user && (user.role === 'admin' || result.userUUID === user.id) ? (
                <DeleteOutlined
                  key="delete"
                  onClick={() => void handleDelete(result.id)}
                  style={{ color: 'red' }}
                />
              ) : null,
            ]}
          >
            <List.Item.Meta
              title={`${getTeamName(result.teamId)} vs ${getTeamName(
                result.opponentTeamId,
              )}`}
              description={
                <>
                  <Typography.Text>
                    {result.score} - {result.opponentScore}
                  </Typography.Text>
                  <br />
                  <Typography.Text type="secondary">
                    {new Date(result.matchStartDate).toLocaleDateString()}
                  </Typography.Text>
                  <br />
                  <Typography.Text type="secondary">
                    {`By: ${getUserById(result.userUUID).username}`}
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
        // eslint-disable-next-line @typescript-eslint/no-misused-promises
        onOk={handleCreateOrEdit}
      >
        <DatePicker
          placeholder="Match Start Date"
          value={formData.matchStartDate}
          onChange={(date) => {
            setFormData({ ...formData, matchStartDate: date });
          }}
          style={{ marginBottom: 8 }}
        />
        <DatePicker
          placeholder="Match End Date"
          value={formData.matchEndDate}
          onChange={(date) => {
            setFormData({ ...formData, matchEndDate: date });
          }}
          style={{ marginBottom: 8 }}
        />
        <Input
          placeholder="Score"
          value={formData.score}
          onChange={(e) => {
            setFormData({ ...formData, score: e.target.value });
          }}
          style={{ marginBottom: 8 }}
        />
        <Input
          placeholder="Opponent Score"
          value={formData.opponentScore}
          onChange={(e) => {
            setFormData({ ...formData, opponentScore: e.target.value });
          }}
          style={{ marginBottom: 8 }}
        />
        <Select
          placeholder="Select Opponent Team"
          value={formData.opponentTeamId}
          onChange={(value) => {
            setFormData({ ...formData, opponentTeamId: value });
          }}
          style={{ width: '100%', marginBottom: 8 }}
        >
          {(teams ?? [])
            .filter((opponentTeam) => opponentTeam.id !== teamIdInteger)
            .map((opponentTeam) => (
              <Option
                key={opponentTeam.id}
                value={opponentTeam.id}
              >
                {opponentTeam.name}
              </Option>
            ))}
        </Select>
      </Modal>
    </div>
  );
}
