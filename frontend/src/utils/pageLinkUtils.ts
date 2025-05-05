import PARAMS from 'constants/params';
import ROUTE_PATHS from 'constants/route_paths';

export const getTeamsPageLink = (leagueId: string) => ({
  pathname: ROUTE_PATHS.TEAMS,
  search: `?${PARAMS.LEAGUE_ID}=${leagueId}`,
});

export const getResultsPageLink = (leagueId: string, teamId: string) => ({
  pathname: ROUTE_PATHS.RESULTS,
  search: `?${PARAMS.LEAGUE_ID}=${leagueId}&${PARAMS.TEAM_ID}=${teamId}`,
});
