import { DeleteOutlined, EditOutlined, PlusOutlined } from '@ant-design/icons';
import { Grid } from '@mui/material';
import { Button, List, Modal, Space, Typography } from 'antd';
import dayjs from 'dayjs';
import { useState } from 'react';
import { Link } from 'react-router-dom';
import {
  useCreateSeason,
  useDeleteSeason,
  useFetchSeasons,
  useUpdateSeason,
} from 'api/hooks/seasonsHooks';
import { StatusHandler } from 'components/StatusHandler';
import withAuth from 'hocs/withAuth';
import withErrorBoundary from 'hocs/withErrorBoundary';
import { useAppForm } from 'hooks/form/useAppForm';
import { useAuth } from 'providers/AuthProvider';
import { SeasonDto } from 'types/dtos/responses/seasons/seasonDto';
import { getStartOfDay } from 'utils/dateUtils';
import { seasonDtoValidator } from 'validators/seasons/seasonDtoValidator';

function SeasonsPage() {
  const { user } = useAuth();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isEditing, setIsEditing] = useState(false);
  const [editingSeasonId, setEditingSeasonId] = useState<string | null>(null);

  const {
    data: seasons,
    isLoading: seasonsLoading,
    error: seasonsError,
  } = useFetchSeasons();
  const { mutateAsync: createSeasonMutation } = useCreateSeason();
  const { mutateAsync: updateSeasonMutation } = useUpdateSeason();
  const { mutateAsync: deleteSeasonMutation } = useDeleteSeason();

  const form = useAppForm({
    defaultValues: {
      name: '',
      startDate: getStartOfDay(),
      endDate: getStartOfDay(7),
    },
    validators: {
      onSubmit: seasonDtoValidator,
    },
    onSubmit: async ({ value }) => {
      if (isEditing && editingSeasonId) {
        await updateSeasonMutation({
          seasonId: editingSeasonId,
          payload: value,
        });
      } else {
        await createSeasonMutation(value);
      }

      closeModal();
    },
  });

  const openEditModal = (season: SeasonDto) => {
    setIsEditing(true);
    setEditingSeasonId(season.id);
    form.setFieldValue('name', season.name);
    form.setFieldValue('startDate', season.startDate);
    form.setFieldValue('endDate', season.endDate);
    setIsModalOpen(true);
  };

  const openCreateModal = () => {
    setIsEditing(false);
    form.reset();
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    form.reset();
    setEditingSeasonId(null);
  };

  const handleDelete = async (seasonId: string) => {
    await deleteSeasonMutation(seasonId);
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
        <Typography.Title level={4}>Seasons</Typography.Title>
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
          Create Season
        </Button>
      </Space>
      <StatusHandler
        isLoading={seasonsLoading}
        error={seasonsError}
        isEmpty={!seasons || (Array.isArray(seasons) && seasons.length === 0)}
      >
        <List
          loading={seasonsLoading}
          bordered
          dataSource={seasons}
          renderItem={(season) => (
            <List.Item
              actions={[
                user &&
                (user.role === 'moderator' ||
                  user.role === 'admin' ||
                  season.user.id === user.id) ? (
                  <EditOutlined
                    key="edit"
                    onClick={() => {
                      openEditModal(season);
                    }}
                  />
                ) : null,

                user &&
                (user.role === 'admin' || season.user.id === user.id) ? (
                  <DeleteOutlined
                    key="delete"
                    onClick={() => void handleDelete(season.id)}
                    style={{ color: 'red' }}
                  />
                ) : null,
              ]}
            >
              <List.Item.Meta
                title={
                  <Link to={`/seasons/${season.id.toString()}/teams`}>
                    {season.name}
                  </Link>
                }
                description={
                  <>
                    <Typography.Text>
                      {`${dayjs(season.startDate).format('YYYY-MM-DD')} - ${dayjs(
                        season.endDate,
                      ).format('YYYY-MM-DD')}`}
                    </Typography.Text>
                    <br />
                    <Typography.Text type="secondary">
                      {`By: ${season.user.username}`}
                    </Typography.Text>
                  </>
                }
              />
            </List.Item>
          )}
        />
      </StatusHandler>

      <Modal
        title={isEditing ? 'Edit Season' : 'Create Season'}
        open={isModalOpen}
        onCancel={closeModal}
        onOk={() => void form.handleSubmit()}
      >
        <form.AppField name="name">
          {(field) => <field.Text label="Season Name" />}
        </form.AppField>

        <Grid
          container
          spacing={2}
        >
          <Grid size={6}>
            <form.AppField name="startDate">
              {(field) => <field.DatePicker label="Start Date" />}
            </form.AppField>
          </Grid>
          <Grid size={6}>
            <form.AppField name="endDate">
              {(field) => <field.DatePicker label="End Date" />}
            </form.AppField>
          </Grid>
        </Grid>
      </Modal>
    </div>
  );
}

export default withAuth(withErrorBoundary(SeasonsPage), {
  isLoggedInOnly: true,
});
