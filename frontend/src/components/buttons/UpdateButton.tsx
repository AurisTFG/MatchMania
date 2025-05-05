import { Edit } from '@mui/icons-material';
import { IconButton } from '@mui/material';

type UpdateButtonProps = {
  onClick: () => void;
};

export default function UpdateButton({ onClick }: UpdateButtonProps) {
  return (
    <IconButton
      edge="end"
      onClick={onClick}
    >
      <Edit />
    </IconButton>
  );
}
