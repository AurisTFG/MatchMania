import api from "./api";
import { handleApiError } from "./utils";
import { Season, SeasonResponse, SeasonsResponse } from "../types";

export async function getAllSeasons(): Promise<Season[]> {
  try {
    const { data } = await api.get<SeasonsResponse>("/seasons");

    return data.seasons || [];
  } catch (error) {
    handleApiError(error);
  }
}

export async function getSeason(seasonID: number): Promise<Season> {
  try {
    const { data } = await api.get<SeasonResponse>(`/seasons/${seasonID}`);

    return data.season || ({} as Season);
  } catch (error) {
    handleApiError(error);
  }
}

export async function createSeason(season: {
  name: string;
  startDate: Date;
  endDate: Date;
}): Promise<Season> {
  try {
    const { data } = await api.post<SeasonResponse>("/seasons", season);

    return data.season || ({} as Season);
  } catch (error) {
    handleApiError(error);
  }
}

export async function updateSeason(
  seasonID: number,
  season: {
    name: string;
    startDate: Date;
    endDate: Date;
  }
): Promise<Season> {
  try {
    const { data } = await api.patch<SeasonResponse>(
      `/seasons/${seasonID}`,
      season
    );

    return data.season || ({} as Season);
  } catch (error) {
    handleApiError(error);
  }
}

export async function deleteSeason(seasonID: number): Promise<void> {
  try {
    await api.delete(`/seasons/${seasonID}`);
  } catch (error) {
    handleApiError(error);
  }
}
