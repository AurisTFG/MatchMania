import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { UseAuth } from "./components/Auth/AuthContext.tsx";
import { useEffect } from "react";
import Header from "./components/Header/Header.tsx";
import Content from "./components/Content/Content.tsx";
import Footer from "./components/Footer/Footer.tsx";
import NotFound from "./pages/NotFound/NotFound";
import HomePage from "./pages/Home/HomePage.tsx";
import Login from "./pages/Auth/Login";
import Signup from "./pages/Auth/Signup";
import Profile from "./pages/Profile/Profile.tsx";
import SeasonsPage from "./pages/Seasons/SeasonsPage.tsx";
import TeamsPage from "./pages/Seasons/TeamsPage.tsx";
import ResultsPage from "./pages/Seasons/ResultsPage.tsx";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import "./App.css";
import { getMe } from "./api/auth.ts";

const theme = createTheme({
  typography: {
    fontFamily: "'Lato'",
  },
});

function App() {
  const { setUser } = UseAuth();

  useEffect(() => {
    const initializeAuth = async () => {
      try {
        const user = await getMe();

        console.log("Current user:", user);

        setUser(user);
      } catch (error) {
        console.error("Get current user failed:", error);

        setUser(null);
      }
    };

    initializeAuth();
  }, [setUser]);

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <div className="App">
        <Router>
          <Header />
          <Content>
            <Routes>
              <Route path="/" element={<HomePage />} />
              <Route path="/login" element={<Login />} />
              <Route path="/signup" element={<Signup />} />
              <Route path="/profile" element={<Profile />} />
              <Route path="/seasons" element={<SeasonsPage />} />
              <Route path="/seasons/:seasonId/teams" element={<TeamsPage />} />
              <Route
                path="/seasons/:seasonId/teams/:teamId/results"
                element={<ResultsPage />}
              />

              <Route path="*" element={<NotFound />} />
            </Routes>
          </Content>
          <Footer />
        </Router>
      </div>
    </ThemeProvider>
  );
}

export default App;
