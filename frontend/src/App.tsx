import { Route, Routes } from 'react-router-dom';
import { ErrorBoundary } from './components/ErrorBoundary';
import { Content } from './components/Layout/Content';
import { Footer } from './components/Layout/Footer';
import { Header } from './components/Layout/Header';
import { ROUTES } from './constants/routes';
import { HomePage } from './pages/HomePage';
import { LoginPage } from './pages/LoginPage';
import { NotFoundPage } from './pages/NotFoundPage';
import { ProfilePage } from './pages/ProfilePage';
import { ResultsPage } from './pages/ResultsPage';
import { SeasonsPage } from './pages/SeasonsPage';
import { SignupPage } from './pages/SignupPage';
import { TeamsPage } from './pages/TeamsPage';
import { AllProviders } from './providers/AllProviders';
import './styles/global.css';

export default function App() {
  return (
    <AllProviders>
      <Header />
      <Content>
        <ErrorBoundary>
          <Routes>
            <Route
              path={ROUTES.HOME}
              element={<HomePage />}
            />
            <Route
              path={ROUTES.LOGIN}
              element={<LoginPage />}
            />
            <Route
              path={ROUTES.SIGNUP}
              element={<SignupPage />}
            />
            <Route
              path={ROUTES.PROFILE}
              element={<ProfilePage />}
            />
            <Route
              path={ROUTES.SEASONS}
              element={<SeasonsPage />}
            />
            <Route
              path={ROUTES.TEAMS(':seasonId')}
              element={<TeamsPage />}
            />
            <Route
              path={ROUTES.RESULTS(':seasonId', ':teamId')}
              element={<ResultsPage />}
            />
            <Route
              path={ROUTES.NOT_FOUND}
              element={<NotFoundPage />}
            />
          </Routes>
        </ErrorBoundary>
      </Content>
      <Footer />
    </AllProviders>
  );
}
