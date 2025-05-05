import { useState } from 'react';
import { useDeleteResult } from 'api/hooks/resultsHooks';
import DeleteButton from 'components/buttons/DeleteButton';
import ConfirmationDialog from 'components/dialogs/ConfirmationDialog';
import withAuth from 'hocs/withAuth';
import { ResultDto } from 'types/dtos/responses/results/resultDto';
import { Permission } from 'types/enums/permission';

type DeleteResultButtonProps = {
  result: ResultDto;
};

function DeleteResultButton({ result }: DeleteResultButtonProps) {
  const [open, setOpen] = useState(false);

  const { mutateAsync: deleteResultAsync } = useDeleteResult();

  const handleDeleteAsync = async () => {
    await deleteResultAsync(result.id);
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
        title="Delete Result"
        description={`Are you sure you want to delete this result? This action cannot be undone.`}
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

export default withAuth(DeleteResultButton, {
  permission: Permission.ManageResult,
  dataOwnerUserId: (props) => (props as DeleteResultButtonProps).result.user.id,
  redirect: false,
});
