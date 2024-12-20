import api from "./api";
import { handleApiError } from "./utils";
import { Result, ResultResponse, ResultsResponse } from "../types";

export async function getAllResults(
  seasonID: number,
  teamID: number
): Promise<Result[]> {
  try {
    const { data } = await api.get<ResultsResponse>(
      `seasons/${seasonID}/teams/${teamID}/results`
    );

    return data.results || [];
  } catch (error) {
    handleApiError(error);
  }
}

export async function getResult(
  seasonID: number,
  teamID: number,
  resultID: number
): Promise<Result> {
  try {
    const { data } = await api.get<ResultResponse>(
      `seasons/${seasonID}/teams/${teamID}/results/${resultID}`
    );

    return data.result || ({} as Result);
  } catch (error) {
    handleApiError(error);
  }
}

export async function createResult(
  seasonID: number,
  teamID: number,
  result: {
    matchStartDate: Date;
    matchEndDate: Date;
    score: string;
    opponentScore: string;
    opponentTeamId: number;
  }
): Promise<Result> {
  try {
    const { data } = await api.post<ResultResponse>(
      `seasons/${seasonID}/teams/${teamID}/results`,
      result
    );

    return data.result || ({} as Result);
  } catch (error) {
    handleApiError(error);
  }
}

export async function updateResult(
  seasonID: number,
  teamID: number,
  resultID: number,
  result: {
    matchStartDate: Date;
    matchEndDate: Date;
    score: string;
    opponentScore: string;
  }
): Promise<Result> {
  try {
    const { data } = await api.patch<ResultResponse>(
      `seasons/${seasonID}/teams/${teamID}/results/${resultID}`,
      result
    );

    return data.result || ({} as Result);
  } catch (error) {
    handleApiError(error);
  }
}

export async function deleteResult(
  seasonID: number,
  teamID: number,
  resultID: number
): Promise<void> {
  try {
    await api.delete(`seasons/${seasonID}/teams/${teamID}/results/${resultID}`);
  } catch (error) {
    handleApiError(error);
  }
}
