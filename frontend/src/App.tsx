import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { UseAuth } from "./components/Auth/AuthContext.tsx";
import { useEffect } from "react";
import { getCurrentUser } from "./api/users.ts";
import Header from "./components/Header/Header.tsx";
import Content from "./components/Content/Content.tsx";
import Footer from "./components/Footer/Footer.tsx";
import NotFound from "./pages/NotFound/NotFound";
import GuestHomePage from "./pages/Home/GuestHomePage.tsx";
import UserHomePage from "./pages/Home/UserHomePage.tsx";
import Login from "./pages/Auth/Login";
import Signup from "./pages/Auth/Signup";
import Profile from "./pages/Profile/Profile.tsx";
import SeasonsPage from "./pages/Seasons/SeasonsPage.tsx";
import TeamsPage from "./pages/Seasons/TeamsPage.tsx";
import ResultsPage from "./pages/Seasons/ResultsPage.tsx";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import CssBaseline from "@mui/material/CssBaseline";
import "./App.css";

const theme = createTheme({
  typography: {
    fontFamily: "'Lato'", // Replace with your desired font(s)
  },
});

function App() {
  const { user, setUser } = UseAuth();

  useEffect(() => {
    const initializeAuth = async () => {
      try {
        const user = await getCurrentUser();

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
              <Route
                path="/"
                element={user ? <UserHomePage /> : <GuestHomePage />}
              />
              <Route
                path="/login"
                element={user ? <UserHomePage /> : <Login />}
              />
              <Route
                path="/signup"
                element={user ? <UserHomePage /> : <Signup />}
              />
              <Route
                path="/profile"
                element={!user ? <GuestHomePage /> : <Profile />}
              />
              <Route
                path="/seasons"
                element={!user ? <GuestHomePage /> : <SeasonsPage />}
              />
              <Route
                path="/seasons/:seasonId/teams"
                element={!user ? <GuestHomePage /> : <TeamsPage />}
              />
              <Route
                path="/seasons/:seasonId/teams/:teamId/results"
                element={!user ? <GuestHomePage /> : <ResultsPage />}
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
