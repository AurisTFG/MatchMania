import type { TrackmaniaOAuthTracksDto } from '../trackmaniaOAuth/tracklmaniaOAuthTracksDto';

export type UserDto = {
  id: string;
  username: string;
  email: string;
  role: string;
  country: string | null;
  permissions: string[];
  trackmaniaId: string | null;
  trackmaniaName: string | null;
  tracks: TrackmaniaOAuthTracksDto[];
};
