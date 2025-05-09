import { useEndMatch } from 'api/hooks/matchHooks';
import ConfirmationDialog from 'components/dialogs/ConfirmationDialog';
import { EndMatchDto } from 'types/dtos/requests/matches/endMatchDto';

type EndMatchDialogProps = {
  open: boolean;
  onClose: () => void;
};

export default function EndMatchDialog({ open, onClose }: EndMatchDialogProps) {
  const { mutateAsync: endMatchAsync } = useEndMatch();

  return (
    <ConfirmationDialog
      open={open}
      title="End Match"
      description="Are you sure you want to end this match? This action cannot be undone."
      confirmText="End Match"
      cancelText="Cancel"
      onConfirm={() => {
        void endMatchAsync({} as EndMatchDto);
        onClose();
      }}
      onCancel={() => {
        onClose();
      }}
    />
  );
}
