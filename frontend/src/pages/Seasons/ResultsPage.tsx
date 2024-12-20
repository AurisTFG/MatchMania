import React, { useEffect, useState } from "react";
import {
  getAllResults,
  createResult,
  updateResult,
  deleteResult,
} from "../../api/results.ts";
import { getSeason } from "../../api/seasons.ts";
import { getTeam } from "../../api/teams.ts";
import { Result, Team, Season } from "../../types/index.ts";
import {
  Modal,
  Button,
  Input,
  List,
  Space,
  Typography,
  DatePicker,
  message,
} from "antd";
import { EditOutlined, DeleteOutlined, PlusOutlined } from "@ant-design/icons";
import { useParams } from "react-router-dom";
import moment from "moment";

const isValidResult = (
  seasonId: string | undefined,
  teamId: string | undefined
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

const ResultsPage: React.FC = () => {
  const { seasonId, teamId } = useParams<{
    seasonId: string;
    teamId: string;
  }>();
  const [season, setSeason] = useState<Partial<Season>>({});
  const [team, setTeam] = useState<Partial<Team>>({});

  const [results, setResults] = useState<Result[]>([]);
  const [loading, setLoading] = useState(false);
  const [resultsNotFound, setResultsNotFound] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingResult, setEditingResult] = useState<Partial<Result>>({});
  const [formData, setFormData] = useState({
    matchStartDate: moment(),
    matchEndDate: moment(),
    score: "",
    opponentScore: "",
    opponentTeamId: 0,
  });

  const fetchResults = async () => {
    if (!seasonId || !teamId || resultsNotFound) return;

    try {
      const data = await getAllResults(parseInt(seasonId), parseInt(teamId));
      setResults(data);
    } catch (error) {
      message.error("Failed to fetch results.");
      console.error(error);
      setResultsNotFound(true);
    }
  };

  const fetchSeason = async () => {
    if (!seasonId) return;

    try {
      const data = await getSeason(parseInt(seasonId));

      setSeason(data);
    } catch (error) {
      console.error(error);
      setResultsNotFound(true);
    }
  };

  const fetchTeam = async () => {
    if (!teamId || !seasonId) return;

    try {
      const data = await getTeam(parseInt(seasonId), parseInt(teamId));

      setTeam(data);
    } catch (error) {
      console.error(error);
      setResultsNotFound(true);
    }
  };

  useEffect(() => {
    if (!isValidResult(seasonId, teamId)) {
      setResultsNotFound(true);
      return;
    }

    setLoading(true);
    fetchSeason();
    fetchTeam();
    fetchResults();
    setLoading(false);
  }, [seasonId, teamId, resultsNotFound]);

  if (resultsNotFound) {
    return (
      <div style={{ padding: 20 }}>
        <Typography.Title level={4}>Results not found</Typography.Title>
      </div>
    );
  }

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

  const openCreateModal = () => {
    setIsEditing(false);
    setEditingResult({});
    setFormData({
      matchStartDate: moment(),
      matchEndDate: moment(),
      score: "",
      opponentScore: "",
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
      score: "",
      opponentScore: "",
      opponentTeamId: 0,
    });
  };

  const handleCreateOrEdit = async () => {
    try {
      if (isEditing && editingResult.id) {
        await updateResult(
          parseInt(seasonId!),
          parseInt(teamId!),
          editingResult.id,
          {
            matchStartDate: formData.matchStartDate.toDate(),
            matchEndDate: formData.matchEndDate.toDate(),
            score: formData.score,
            opponentScore: formData.opponentScore,
          }
        );
        message.success("Result updated successfully.");
      } else {
        await createResult(parseInt(seasonId!), parseInt(teamId!), {
          matchStartDate: formData.matchStartDate.toDate(),
          matchEndDate: formData.matchEndDate.toDate(),
          score: formData.score,
          opponentScore: formData.opponentScore,
          opponentTeamId: formData.opponentTeamId,
        });
        message.success("Result created successfully.");
      }
      setIsModalOpen(false);
      fetchResults();
    } catch (error) {
      message.error("Failed to save result.");
      console.error(error);
    }
  };

  const handleDelete = async (resultID: number) => {
    try {
      await deleteResult(parseInt(seasonId!), parseInt(teamId!), resultID);
      message.success("Result deleted successfully.");
      fetchResults();
    } catch (error) {
      message.error("Failed to delete result.");
      console.error(error);
    }
  };

  return (
    <div style={{ padding: 20 }}>
      <Space
        style={{
          marginBottom: 16,
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <Typography.Title level={4}>
          Results for Team "{team.name}" in Season "{season.name}"
        </Typography.Title>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={openCreateModal}
        >
          Create Result
        </Button>
      </Space>

      <List
        loading={loading}
        bordered
        dataSource={results}
        renderItem={(result) => (
          <List.Item
            actions={[
              <EditOutlined key="edit" onClick={() => openEditModal(result)} />,
              <DeleteOutlined
                key="delete"
                onClick={() => handleDelete(result.id)}
                style={{ color: "red" }}
              />,
            ]}
          >
            <List.Item.Meta
              title={`Score: ${result.score} - Opponent Score: ${result.opponentScore}`}
              description={`Date: ${new Date(result.matchStartDate).toLocaleDateString()}`}
            />
          </List.Item>
        )}
      />

      <Modal
        title={isEditing ? "Edit Result" : "Create Result"}
        open={isModalOpen}
        onCancel={closeModal}
        onOk={handleCreateOrEdit}
      >
        <DatePicker
          placeholder="Match Start Date"
          value={formData.matchStartDate}
          onChange={(date) =>
            setFormData({ ...formData, matchStartDate: date })
          }
          style={{ marginBottom: 8 }}
        />
        <DatePicker
          placeholder="Match End Date"
          value={formData.matchEndDate}
          onChange={(date) => setFormData({ ...formData, matchEndDate: date })}
          style={{ marginBottom: 8 }}
        />
        <Input
          placeholder="Score"
          value={formData.score}
          onChange={(e) => setFormData({ ...formData, score: e.target.value })}
          style={{ marginBottom: 8 }}
        />
        <Input
          placeholder="Opponent Score"
          value={formData.opponentScore}
          onChange={(e) =>
            setFormData({ ...formData, opponentScore: e.target.value })
          }
          style={{ marginBottom: 8 }}
        />
        <Input
          placeholder="Opponent Team ID"
          value={formData.opponentTeamId}
          onChange={(e) =>
            setFormData({
              ...formData,
              opponentTeamId: parseInt(e.target.value),
            })
          }
        />
      </Modal>
    </div>
  );
};

export default ResultsPage;
