import { DeleteOutlined, EditOutlined, PlusOutlined } from "@ant-design/icons";
import {
  Button,
  DatePicker,
  Input,
  List,
  Modal,
  Space,
  Typography,
  message,
} from "antd";
import moment from "moment";
import { useState } from "react";
import { Link } from "react-router-dom";
import {
  useCreateSeason,
  useDeleteSeason,
  useFetchSeasons,
  useUpdateSeason,
} from "../../api/hooks/seasonsHooks.ts";
import { useFetchUsers } from "../../api/hooks/usersHooks.ts";
import { useAuth } from "../../providers/AuthProvider.tsx";
import { Season } from "../../types/index.ts";
import { User } from "../../types/users.ts";

export default function SeasonsPage() {
  const { user } = useAuth();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingSeason, setEditingSeason] = useState<Partial<Season>>({});
  const [formData, setFormData] = useState({
    name: "",
    startDate: moment(),
    endDate: moment(),
  });

  const { data: seasons = [], isLoading: seasonsLoading } = useFetchSeasons();
  const { data: users = [] } = useFetchUsers();
  const { mutateAsync: createSeasonMutation } = useCreateSeason();
  const { mutateAsync: updateSeasonMutation } = useUpdateSeason();
  const { mutateAsync: deleteSeasonMutation } = useDeleteSeason();

  const getUserById = (userId: string): User => {
    return (
      users.find((user) => user.id === userId) ??
      ({
        id: "",
        username: "Unknown User",
      } as User)
    );
  };

  const handleCreateOrEdit = async () => {
    try {
      if (isEditing && editingSeason.id) {
        await updateSeasonMutation({
          seasonID: editingSeason.id,
          season: {
            ...formData,
            startDate: formData.startDate.toDate(),
            endDate: formData.endDate.toDate(),
          },
        });

        message.success("Season updated successfully.");
      } else {
        await createSeasonMutation({
          ...formData,
          startDate: formData.startDate.toDate(),
          endDate: formData.endDate.toDate(),
        });

        message.success("Season created successfully.");
      }
      setIsModalOpen(false);
    } catch (error) {
      message.error("Failed to save season.");
      console.error(error);
    }
  };

  const handleDelete = async (seasonID: number) => {
    try {
      await deleteSeasonMutation(seasonID);
      message.success("Season deleted successfully.");
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
    <div style={{ padding: 20, width: "50%", margin: "auto" }}>
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
          disabled={user === null}
          style={{
            filter: user === null ? "blur(1px)" : "none",
            cursor: user === null ? "not-allowed" : "pointer",
          }}
        >
          Create Season
        </Button>
      </Space>

      <List
        loading={seasonsLoading}
        bordered
        dataSource={seasons}
        renderItem={(season) => (
          <List.Item
            actions={[
              user &&
              (user.role === "moderator" ||
                user.role === "admin" ||
                season.userUUID === user.id) ? (
                <EditOutlined
                  key="edit"
                  onClick={() => {
                    openEditModal(season);
                  }}
                />
              ) : null,

              user && (user.role === "admin" || season.userUUID === user.id) ? (
                <DeleteOutlined
                  key="delete"
                  onClick={() => handleDelete(season.id)}
                  style={{ color: "red" }}
                />
              ) : null,
            ]}
          >
            <List.Item.Meta
              title={
                <Link to={`/seasons/${season.id}/teams`}>{season.name}</Link>
              }
              description={
                <>
                  <Typography.Text>
                    {`${season.startDate.toLocaleString().split("T")[0]} - ${
                      season.endDate.toLocaleString().split("T")[0]
                    }`}
                  </Typography.Text>
                  <br />
                  <Typography.Text type="secondary">
                    {`By: ${getUserById(season.userUUID).username}`}
                  </Typography.Text>
                </>
              }
            />
          </List.Item>
        )}
      />

      <Modal
        title={isEditing ? "Edit Season" : "Create Season"}
        open={isModalOpen}
        onCancel={closeModal}
        // eslint-disable-next-line @typescript-eslint/no-misused-promises
        onOk={handleCreateOrEdit}
      >
        <Input
          placeholder="Name"
          value={formData.name}
          onChange={(e) => {
            setFormData({ ...formData, name: e.target.value });
          }}
          style={{ marginBottom: 8 }}
        />
        <DatePicker
          placeholder="Start Date"
          value={formData.startDate}
          onChange={(date) => {
            setFormData({ ...formData, startDate: date });
          }}
          style={{ marginBottom: 8 }}
        />
        <DatePicker
          placeholder="End Date"
          value={formData.endDate}
          onChange={(date) => {
            setFormData({ ...formData, endDate: date });
          }}
        />
      </Modal>
    </div>
  );
}
