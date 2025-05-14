export type CreateLeagueDto = {
  name: string;
  logoUrl: string | null;
  startDate: Date;
  endDate: Date;

  trackIds: string[];
};
