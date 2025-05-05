import { Avatar } from '@mui/material';
import { useTheme } from '@mui/material/styles';

type UserAvatarProps = {
  imageUrl?: string | null;
  name?: string | null;
  size?: number;
  marginLeft?: number;
};

export default function UserAvatar({
  imageUrl,
  name,
  size = 36,
  marginLeft = 0,
}: UserAvatarProps) {
  const theme = useTheme();

  return imageUrl ? (
    <Avatar
      src={imageUrl}
      sx={{ width: size, height: size }}
    />
  ) : (
    <Avatar
      sx={{
        width: size,
        height: size,
        bgcolor: theme.palette.primary.main,
        fontSize: size * 0.7,
        marginLeft: `${String(marginLeft)}px`,
      }}
    >
      {name?.[0]?.toUpperCase() ?? 'G'}
    </Avatar>
  );
}
