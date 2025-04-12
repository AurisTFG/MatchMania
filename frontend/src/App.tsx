import { Route, Routes } from 'react-router-dom';
import Content from './components/Content/Content';
import Footer from './components/Footer/Footer';
import Header from './components/Header/Header';
import Login from './pages/Auth/Login';
import Signup from './pages/Auth/Signup';
import HomePage from './pages/Home/HomePage';
import NotFound from './pages/NotFound/NotFound';
import Profile from './pages/Profile/Profile';
import ResultsPage from './pages/Seasons/ResultsPage';
import SeasonsPage from './pages/Seasons/SeasonsPage';
import TeamsPage from './pages/Seasons/TeamsPage';
import Providers from './providers/Providers';
import './styles/global.css';

export default function App() {
  return (
    <Providers>
      <Header />
      <Content>
        <Routes>
          <Route
            path="/"
            element={<HomePage />}
          />
          <Route
            path="/login"
            element={<Login />}
          />
          <Route
            path="/signup"
            element={<Signup />}
          />
          <Route
            path="/profile"
            element={<Profile />}
          />
          <Route
            path="/seasons"
            element={<SeasonsPage />}
          />
          <Route
            path="/seasons/:seasonId/teams"
            element={<TeamsPage />}
          />
          <Route
            path="/seasons/:seasonId/teams/:teamId/results"
            element={<ResultsPage />}
          />
          <Route
            path="*"
            element={<NotFound />}
          />
        </Routes>
      </Content>
      <Footer />
    </Providers>
  );
}
