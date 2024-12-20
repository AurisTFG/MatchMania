import api from "./api";
import { handleApiError } from "./utils";
import { Team, TeamResponse, TeamsResponse } from "../types";

export async function getAllTeams(seasonID: number): Promise<Team[]> {
  try {
    const { data } = await api.get<TeamsResponse>(`seasons/${seasonID}/teams`);

    return data.teams || [];
  } catch (error) {
    handleApiError(error);
  }
}

export async function getTeam(seasonID: number, teamID: number): Promise<Team> {
  try {
    const { data } = await api.get<TeamResponse>(
      `seasons/${seasonID}/teams/${teamID}`
    );

    return data.team || ({} as Team);
  } catch (error) {
    handleApiError(error);
  }
}

export async function createTeam(
  seasonID: number,
  team: { name: string }
): Promise<Team> {
  try {
    const { data } = await api.post<TeamResponse>(
      `seasons/${seasonID}/teams`,
      team
    );

    return data.team || ({} as Team);
  } catch (error) {
    handleApiError(error);
  }
}

export async function updateTeam(
  seasonID: number,
  teamID: number,
  team: { name: string }
): Promise<Team> {
  try {
    const { data } = await api.patch<TeamResponse>(
      `seasons/${seasonID}/teams/${teamID}`,
      team
    );

    return data.team || ({} as Team);
  } catch (error) {
    handleApiError(error);
  }
}

export async function deleteTeam(
  seasonID: number,
  teamID: number
): Promise<void> {
  try {
    await api.delete(`seasons/${seasonID}/teams/${teamID}`);
  } catch (error) {
    handleApiError(error);
  }
}
