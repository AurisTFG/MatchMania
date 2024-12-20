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
import "./styles/theme.ts";
import "./App.css";

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

            <Route path="*" element={<NotFound />} />
          </Routes>
        </Content>
        <Footer />
      </Router>
    </div>
  );
}

export default App;
