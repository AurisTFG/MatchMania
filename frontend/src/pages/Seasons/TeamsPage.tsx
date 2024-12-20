import React, { useEffect, useState } from "react";
import {
  getAllTeams,
  createTeam,
  updateTeam,
  deleteTeam,
} from "../../api/teams.ts";
import { getSeason } from "../../api/seasons.ts";
import { Team, Season } from "../../types/index.ts";
import { Modal, Button, Input, List, Space, Typography, message } from "antd";
import { EditOutlined, DeleteOutlined, PlusOutlined } from "@ant-design/icons";
import { useParams, Link } from "react-router-dom";

const isValidTeam = (seasonId: string | undefined) => {
  return seasonId && !isNaN(Number(seasonId)) && Number(seasonId) > 0;
};

const TeamsPage: React.FC = () => {
  const { seasonId } = useParams<{ seasonId: string }>();
  const [season, setSeason] = useState<Partial<Season>>({});

  const [teams, setTeams] = useState<Team[]>([]);
  const [loading, setLoading] = useState(false);
  const [teamsNotFound, setTeamsNotFound] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingTeam, setEditingTeam] = useState<Partial<Team>>({});
  const [formData, setFormData] = useState({
    name: "",
  });

  const fetchTeams = async () => {
    if (!seasonId || teamsNotFound) return;

    setLoading(true);
    try {
      const data = await getAllTeams(parseInt(seasonId));

      setTeams(data);
    } catch (error) {
      message.error("Failed to fetch teams.");
      console.error(error);
      setTeamsNotFound(true);
    } finally {
      setLoading(false);
    }
  };

  const fetchSeason = async () => {
    if (!seasonId) return;

    try {
      const data = await getSeason(parseInt(seasonId));

      setSeason(data);
    } catch (error) {
      console.error(error);
      setTeamsNotFound(true);
    }
  };

  useEffect(() => {
    if (!isValidTeam(seasonId)) {
      setTeamsNotFound(true);
      return;
    }

    fetchSeason();
    fetchTeams();
  }, [seasonId, teamsNotFound]);

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
      setFormData({ name: "" });
    }
    setIsModalOpen(true);
  };

  const handleCreateOrEdit = async () => {
    try {
      if (isEditing && editingTeam.id) {
        await updateTeam(parseInt(seasonId!), editingTeam.id, formData);
        message.success("Team updated successfully.");
      } else {
        await createTeam(parseInt(seasonId!), formData);
        message.success("Team created successfully.");
      }
      setIsModalOpen(false);
      setFormData({ name: "" });
      fetchTeams();
    } catch (error) {
      message.error("Failed to save team.");
      console.error(error);
    }
  };

  const handleDelete = async (teamID: number) => {
    try {
      await deleteTeam(parseInt(seasonId!), teamID);
      message.success("Team deleted successfully.");
      fetchTeams();
    } catch (error) {
      message.error("Failed to delete team.");
      console.error(error);
    }
  };

  // Close the modal
  const closeModal = () => {
    setIsModalOpen(false);
    setFormData({ name: "" });
    setEditingTeam({});
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
          Teams for Season "{season.name}"
        </Typography.Title>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => openModal()}
        >
          Create Team
        </Button>
      </Space>

      <List
        loading={loading}
        bordered
        dataSource={teams}
        renderItem={(team) => (
          <List.Item
            actions={[
              <EditOutlined key="edit" onClick={() => openModal(team)} />,
              <DeleteOutlined
                key="delete"
                onClick={() => handleDelete(team.id)}
                style={{ color: "red" }}
              />,
            ]}
          >
            <List.Item.Meta
              title={
                <Link to={`/seasons/${seasonId}/teams/${team.id}/results`}>
                  {team.name}
                </Link>
              }
            />
          </List.Item>
        )}
      />

      <Modal
        title={isEditing ? "Edit Team" : "Create Team"}
        open={isModalOpen}
        onCancel={closeModal}
        onOk={handleCreateOrEdit}
      >
        <Input
          placeholder="Team Name"
          value={formData.name}
          onChange={(e) => setFormData({ ...formData, name: e.target.value })}
          style={{ marginBottom: 8 }}
        />
      </Modal>
    </div>
  );
};

export default TeamsPage;
