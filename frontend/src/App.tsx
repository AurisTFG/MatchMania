import { Route, Routes } from 'react-router-dom';
import { Content } from './components/Content';
import { ErrorBoundary } from './components/ErrorBoundary';
import { Footer } from './components/Footer';
import { Header } from './components/Header';
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
      <ErrorBoundary>
        <Content>
          <Routes>
            <Route
              path="/"
              element={<HomePage />}
            />
            <Route
              path="/login"
              element={<LoginPage />}
            />
            <Route
              path="/signup"
              element={<SignupPage />}
            />
            <Route
              path="/profile"
              element={<ProfilePage />}
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
              element={<NotFoundPage />}
            />
          </Routes>
        </Content>
      </ErrorBoundary>
      <Footer />
    </AllProviders>
  );
}
