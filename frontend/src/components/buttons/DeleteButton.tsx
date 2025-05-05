import { Delete } from '@mui/icons-material';
import { IconButton } from '@mui/material';

type DeleteButtonProps = {
  onClick: () => void;
};

export default function DeleteButton({ onClick }: DeleteButtonProps) {
  return (
    <IconButton
      edge="end"
      onClick={onClick}
      sx={{ color: 'error.main' }}
    >
      <Delete />
    </IconButton>
  );
}
