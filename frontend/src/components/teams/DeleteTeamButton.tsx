import { useState } from 'react';
import { useDeleteTeam } from 'api/hooks/teamsHooks';
import DeleteButton from 'components/buttons/DeleteButton';
import ConfirmationDialog from 'components/dialogs/ConfirmationDialog';
import withAuth from 'hocs/withAuth';
import { TeamDto } from 'types/dtos/responses/teams/teamDto';
import { Permission } from 'types/enums/permission';

type DeleteTeamButtonProps = {
  team: TeamDto;
};

function DeleteTeamButton({ team }: DeleteTeamButtonProps) {
  const [open, setOpen] = useState(false);

  const { mutateAsync: deleteTeamAsync } = useDeleteTeam();

  const handleDeleteAsync = async () => {
    await deleteTeamAsync(team.id);
    setOpen(false);
  };

  return (
    <>
      <DeleteButton
        onClick={() => {
          setOpen(true);
        }}
      />

      <ConfirmationDialog
        open={open}
        title="Delete Team"
        description={`Are you sure you want to delete the team "${team.name}"? This action cannot be undone.`}
        confirmText="Delete"
        cancelText="Cancel"
        onConfirm={() => {
          void handleDeleteAsync();
        }}
        onCancel={() => {
          setOpen(false);
        }}
      />
    </>
  );
}

export default withAuth(DeleteTeamButton, {
  permission: Permission.ManageTeam,
  dataOwnerUserId: (props) => (props as DeleteTeamButtonProps).team.user.id,
  redirect: false,
});
