import { TrackmaniaTrackDto } from '../trackmaniaTracks/trackmaniaTrackDto';
import { UserMinimalDto } from '../users/userMinimalDto';

export type LeagueDto = {
  id: string;
  name: string;
  logoUrl: string;
  startDate: Date;
  endDate: Date;

  tracks: TrackmaniaTrackDto[];
  user: UserMinimalDto;
};
