import { Add } from '@mui/icons-material';
import { Button, ButtonProps } from '@mui/material';

type CreateButtonProps = {
  title?: string;
  canCreate: boolean;
  onClick: () => void;
} & Omit<ButtonProps, 'onClick' | 'disabled' | 'startIcon' | 'children'>;

export default function CreateButton({
  title = 'Create',
  canCreate,
  onClick,
}: CreateButtonProps) {
  if (!canCreate) {
    return null;
  }

  return (
    <Button
      variant="contained"
      startIcon={<Add />}
      onClick={onClick}
      sx={{
        filter: 'none',
        cursor: 'pointer',
      }}
    >
      {title}
    </Button>
  );
}
