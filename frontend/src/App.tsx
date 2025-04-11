import { ThemeProvider } from "@mui/material/styles";
import { QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Content from "./components/Content/Content.tsx";
import Footer from "./components/Footer/Footer.tsx";
import Header from "./components/Header/Header.tsx";
import { queryClient } from "./configs/queryClient.ts";
import Login from "./pages/Auth/Login";
import Signup from "./pages/Auth/Signup";
import HomePage from "./pages/Home/HomePage.tsx";
import NotFound from "./pages/NotFound/NotFound";
import Profile from "./pages/Profile/Profile.tsx";
import ResultsPage from "./pages/Seasons/ResultsPage.tsx";
import SeasonsPage from "./pages/Seasons/SeasonsPage.tsx";
import TeamsPage from "./pages/Seasons/TeamsPage.tsx";
import { AuthProvider } from "./providers/AuthProvider.tsx";
import theme from "./styles/theme.ts";
import "./styles/global.css";

export default function App() {
  return (
    <ThemeProvider theme={theme}>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <BrowserRouter>
            <Header />
            <Content>
              <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="/login" element={<Login />} />
                <Route path="/signup" element={<Signup />} />
                <Route path="/profile" element={<Profile />} />
                <Route path="/seasons" element={<SeasonsPage />} />
                <Route
                  path="/seasons/:seasonId/teams"
                  element={<TeamsPage />}
                />
                <Route
                  path="/seasons/:seasonId/teams/:teamId/results"
                  element={<ResultsPage />}
                />
                <Route path="*" element={<NotFound />} />
              </Routes>
            </Content>
            <Footer />
          </BrowserRouter>
        </AuthProvider>
        <ReactQueryDevtools initialIsOpen={false} />
      </QueryClientProvider>
    </ThemeProvider>
  );
}
