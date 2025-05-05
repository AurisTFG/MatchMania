import { useState } from 'react';
import { useDeleteLeague } from 'api/hooks/leaguesHooks';
import DeleteButton from 'components/buttons/DeleteButton';
import ConfirmationDialog from 'components/dialogs/ConfirmationDialog';
import withAuth from 'hocs/withAuth';
import { LeagueDto } from 'types/dtos/responses/leagues/leagueDto';
import { Permission } from 'types/enums/permission';

type DeleteLeagueButtonProps = {
  league: LeagueDto;
};

function DeleteLeagueButton({ league }: DeleteLeagueButtonProps) {
  const [open, setOpen] = useState(false);

  const { mutateAsync: deleteLeagueAsync } = useDeleteLeague();

  const handleDeleteAsync = async () => {
    await deleteLeagueAsync(league.id);
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
        title="Delete League"
        description={`Are you sure you want to delete the league "${league.name}"? This action cannot be undone.`}
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

export default withAuth(DeleteLeagueButton, {
  permission: Permission.ManageLeague,
  dataOwnerUserId: (props) => (props as DeleteLeagueButtonProps).league.user.id,
  redirect: false,
});
