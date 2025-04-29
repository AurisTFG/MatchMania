import { Route, Routes } from 'react-router-dom';
import { ROUTES } from 'constants/routes';
import { AllProviders } from './providers/AllProviders';
import './styles/global.css';

export default function App() {
  return (
    <AllProviders>
      <Routes>
        {ROUTES.map(({ path, element }) => (
          <Route
            key={path}
            path={path}
            element={element}
          />
        ))}
      </Routes>
    </AllProviders>
  );
}
