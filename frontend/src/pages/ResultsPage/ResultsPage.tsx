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
} from 'antd';
import moment from 'moment';
import { useState } from 'react';
import { useParams } from 'react-router-dom';
import {
  useCreateResult,
  useDeleteResult,
  useFetchResults,
  useUpdateResult,
} from '../../api/hooks/resultsHooks';
import { useFetchSeason } from '../../api/hooks/seasonsHooks';
import { useFetchTeam, useFetchTeams } from '../../api/hooks/teamsHooks';
import { useAuth } from '../../providers/AuthProvider';
import { ResultDto } from '../../types/dtos/responses/results/resultDto';

const { Option } = Select;

export default function ResultsPage() {
  const { seasonId = '', teamId = '' } = useParams<{
    seasonId: string;
    teamId: string;
  }>();

  const { user } = useAuth();

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingResult, setEditingResult] = useState<Partial<ResultDto>>({});
  const [formData, setFormData] = useState({
    startDate: moment(),
    endDate: moment(),
    score: '',
    opponentScore: '',
    opponentTeamId: '',
  });

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

  const openCreateModal = () => {
    setIsEditing(false);
    setEditingResult({});
    setFormData({
      startDate: moment(),
      endDate: moment(),
      score: '',
      opponentScore: '',
      opponentTeamId: '',
    });
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setEditingResult({});
    setFormData({
      startDate: moment(),
      endDate: moment(),
      score: '',
      opponentScore: '',
      opponentTeamId: '',
    });
  };

  const openEditModal = (result: ResultDto) => {
    setIsEditing(true);
    setEditingResult(result);
    setFormData({
      startDate: moment(result.startDate),
      endDate: moment(result.endDate),
      score: result.score,
      opponentScore: result.opponentScore,
      opponentTeamId: result.opponentTeam.id,
    });
    setIsModalOpen(true);
  };

  const handleCreateOrEdit = async () => {
    if (isEditing && editingResult.id) {
      await updateResult({
        resultId: editingResult.id,
        payload: {
          startDate: formData.startDate.toDate(),
          endDate: formData.endDate.toDate(),
          score: formData.score,
          opponentScore: formData.opponentScore,
        },
      });
    } else {
      await createResult({
        startDate: formData.startDate.toDate(),
        endDate: formData.endDate.toDate(),
        score: formData.score,
        opponentScore: formData.opponentScore,
        opponentTeamId: formData.opponentTeamId,
      });
    }
    closeModal();
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
        // eslint-disable-next-line @typescript-eslint/no-misused-promises
        onOk={handleCreateOrEdit}
      >
        <DatePicker
          placeholder="Match Start Date"
          value={formData.startDate}
          onChange={(date) => {
            setFormData({ ...formData, startDate: date });
          }}
          style={{ marginBottom: 8 }}
        />
        <DatePicker
          placeholder="Match End Date"
          value={formData.endDate}
          onChange={(date) => {
            setFormData({ ...formData, endDate: date });
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
            .filter((opponentTeam) => opponentTeam.id !== teamId)
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
