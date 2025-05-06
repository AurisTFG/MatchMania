import AccessAlarmIcon from '@mui/icons-material/AccessAlarm';
import GroupIcon from '@mui/icons-material/Group';
import HomeIcon from '@mui/icons-material/Home';
import HourglassBottomIcon from '@mui/icons-material/HourglassBottom';
import SportsEsportsIcon from '@mui/icons-material/SportsEsports';
import ForbiddenPage from 'pages/ForbiddenPage';
import HomePage from 'pages/HomePage';
import LeaguesPage from 'pages/LeaguesPage';
import LoginPage from 'pages/LoginPage';
import MatchmakingQueuePage from 'pages/MatchmakingQueuePage';
import NotFoundPage from 'pages/NotFoundPage';
import ProfilePage from 'pages/ProfilePage';
import ResultsPage from 'pages/ResultsPage';
import SignupPage from 'pages/SignupPage';
import TeamsPage from 'pages/TeamsPage';
import UnauthorizedPage from 'pages/UnauthorizedPage';
import { Permission } from 'types/enums/permission';
import ROUTE_PATHS from './route_paths';

export const ROUTES = [
  {
    label: 'Home',
    path: ROUTE_PATHS.HOME,
    element: <HomePage />,
    icon: <HomeIcon />,
    permission: null,
  },
  {
    label: 'Log In',
    path: ROUTE_PATHS.LOGIN,
    element: <LoginPage />,
    icon: null,
    permission: null,
  },
  {
    label: 'Sign Up',
    path: ROUTE_PATHS.SIGNUP,
    element: <SignupPage />,
    icon: null,
    permission: null,
  },
  {
    label: 'Profile',
    path: ROUTE_PATHS.PROFILE,
    element: <ProfilePage />,
    icon: null,
    permission: Permission.LoggedIn,
  },
  {
    label: 'Leagues',
    path: ROUTE_PATHS.LEAGUES,
    element: <LeaguesPage />,
    icon: <AccessAlarmIcon />,
    permission: Permission.ManageLeague,
  },
  {
    label: 'Teams',
    path: ROUTE_PATHS.TEAMS,
    element: <TeamsPage />,
    icon: <GroupIcon />,
    permission: Permission.ManageTeam,
  },
  {
    label: 'Results',
    path: ROUTE_PATHS.RESULTS,
    element: <ResultsPage />,
    icon: <HourglassBottomIcon />,
    permission: Permission.ViewResult,
  },
  {
    label: 'Matchmaking',
    path: ROUTE_PATHS.MATCHMAKING,
    element: <MatchmakingQueuePage />,
    icon: <SportsEsportsIcon />,
    permission: Permission.ManageQueue,
  },
  {
    label: 'Unauthorized',
    path: ROUTE_PATHS.UNAUTHORIZED,
    element: <UnauthorizedPage />,
    icon: null,
    permission: null,
  },
  {
    label: 'Forbidden',
    path: ROUTE_PATHS.FORBIDDEN,
    element: <ForbiddenPage />,
    icon: null,
    permission: null,
  },
  {
    label: 'Not Found',
    path: ROUTE_PATHS.NOT_FOUND,
    element: <NotFoundPage />,
    icon: null,
    permission: null,
  },
];
