import { Avatar } from '@mui/material';
import { useTheme } from '@mui/material/styles';

type UserAvatarProps = {
  profilePhotoUrl?: string | null;
  username?: string | null;
  size?: number;
};

export default function UserAvatar({
  profilePhotoUrl,
  username,
  size = 36,
}: UserAvatarProps) {
  const theme = useTheme();

  return profilePhotoUrl ? (
    <Avatar
      src={profilePhotoUrl}
      sx={{ width: size, height: size }}
    />
  ) : (
    <Avatar
      sx={{
        width: size,
        height: size,
        bgcolor: theme.palette.primary.main,
        fontSize: size * 0.6,
      }}
    >
      {username?.[0]?.toUpperCase() ?? 'G'}
    </Avatar>
  );
}
