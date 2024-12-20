import React, { useEffect, useState } from "react";
import { getAllTeams } from "../../api/teams.ts";
import { Team } from "../../types/index.ts";
import { List, Button, Space, Typography, message } from "antd";
import { PlusOutlined } from "@ant-design/icons";
import { useParams, useNavigate } from "react-router-dom";

const TeamsPage: React.FC = () => {
  const { seasonId } = useParams<{ seasonId: string }>(); // Get the seasonId from URL
  const [teams, setTeams] = useState<Team[]>([]);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate(); // To navigate programmatically

  useEffect(() => {
    const fetchTeams = async () => {
      if (!seasonId) return;

      setLoading(true);
      try {
        const data = await getAllTeams(parseInt(seasonId)); // Use seasonId to fetch teams
        setTeams(data);
      } catch (error) {
        message.error("Failed to fetch teams.");
        console.error(error);
      } finally {
        setLoading(false);
      }
    };

    fetchTeams();
  }, [seasonId]);

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
          Teams for Season {seasonId}
        </Typography.Title>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => navigate(`/seasons/${seasonId}/teams/create`)} // Navigate to team creation
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
              <Button
                key="edit"
                onClick={() =>
                  navigate(`/seasons/${seasonId}/teams/${team.id}/edit`)
                } // Navigate to edit page
              >
                Edit
              </Button>,
            ]}
          >
            <List.Item.Meta title={team.name} />
          </List.Item>
        )}
      />
    </div>
  );
};

export default TeamsPage;
