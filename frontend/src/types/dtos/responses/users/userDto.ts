import { RoleDto } from '../roles/roleDto';
import { TrackmaniaTrackDto } from '../trackmaniaTracks/trackmaniaTrackDto';

export type UserDto = {
  id: string;
  username: string;
  email: string;
  profilePhotoUrl: string | null;
  country: string | null;
  roles: RoleDto[];
  permissions: string[];
  trackmaniaId: string | null;
  trackmaniaName: string | null;
  tracks: TrackmaniaTrackDto[];
};
