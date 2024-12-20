import React, { useEffect, useState } from "react";
import {
  getAllSeasons,
  createSeason,
  updateSeason,
  deleteSeason,
} from "../../api/seasons.ts";
import { Season } from "../../types/index.ts";
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
import moment from "moment";
import { Link } from "react-router-dom";

const SeasonsPage: React.FC = () => {
  const [seasons, setSeasons] = useState<Season[]>([]);
  const [loading, setLoading] = useState(false);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingSeason, setEditingSeason] = useState<Partial<Season>>({});
  const [formData, setFormData] = useState({
    name: "",
    startDate: moment(),
    endDate: moment(),
  });

  const fetchSeasons = async () => {
    setLoading(true);
    try {
      const data = await getAllSeasons();
      setSeasons(data);
    } catch (error) {
      message.error("Failed to fetch seasons.");
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSeasons();
  }, []);

  const handleCreateOrEdit = async () => {
    try {
      if (isEditing && editingSeason.id) {
        await updateSeason(editingSeason.id, {
          ...formData,
          startDate: formData.startDate.toDate(),
          endDate: formData.endDate.toDate(),
        });

        message.success("Season updated successfully.");
      } else {
        await createSeason({
          ...formData,
          startDate: formData.startDate.toDate(),
          endDate: formData.endDate.toDate(),
        });

        message.success("Season created successfully.");
      }
      setIsModalOpen(false);
      fetchSeasons();
    } catch (error) {
      message.error("Failed to save season.");
      console.error(error);
    }
  };

  const handleDelete = async (seasonID: number) => {
    try {
      await deleteSeason(seasonID);
      message.success("Season deleted successfully.");
      fetchSeasons();
    } catch (error) {
      message.error("Failed to delete season.");
      console.error(error);
    }
  };

  const openEditModal = (season: Season) => {
    setIsEditing(true);
    setEditingSeason(season);
    setFormData({
      name: season.name,
      startDate: moment(season.startDate),
      endDate: moment(season.endDate),
    });
    setIsModalOpen(true);
  };

  const openCreateModal = () => {
    setIsEditing(false);
    setEditingSeason({});
    setFormData({
      name: "",
      startDate: moment(),
      endDate: moment(),
    });
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setEditingSeason({});
    setFormData({
      name: "",
      startDate: moment(),
      endDate: moment(),
    });
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
        <Typography.Title level={4}>Seasons</Typography.Title>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={openCreateModal}
        >
          Create Season
        </Button>
      </Space>

      <List
        loading={loading}
        bordered
        dataSource={seasons}
        renderItem={(season) => (
          <List.Item
            actions={[
              <EditOutlined key="edit" onClick={() => openEditModal(season)} />,
              <DeleteOutlined
                key="delete"
                onClick={() => handleDelete(season.id)}
                style={{ color: "red" }}
              />,
            ]}
          >
            <List.Item.Meta
              title={
                <Link to={`/seasons/${season.id}/teams`}>{season.name}</Link>
              }
              description={`${season.startDate.toLocaleString().split("T")[0]} - ${season.endDate.toLocaleString().split("T")[0]}`}
            />
          </List.Item>
        )}
      />

      <Modal
        title={isEditing ? "Edit Season" : "Create Season"}
        open={isModalOpen}
        onCancel={closeModal}
        onOk={handleCreateOrEdit}
      >
        <Input
          placeholder="Name"
          value={formData.name}
          onChange={(e) => setFormData({ ...formData, name: e.target.value })}
          style={{ marginBottom: 8 }}
        />
        <DatePicker
          placeholder="Start Date"
          value={formData.startDate}
          onChange={(date) => setFormData({ ...formData, startDate: date })}
          style={{ marginBottom: 8 }}
        />
        <DatePicker
          placeholder="End Date"
          value={formData.endDate}
          onChange={(date) => setFormData({ ...formData, endDate: date })}
        />
      </Modal>
    </div>
  );
};

export default SeasonsPage;
