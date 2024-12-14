import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Header from "./components/Header";
import Footer from "./components/Footer";
import NotFound from "./pages/NotFound/NotFound";
import Home from "./pages/Home/Home";
import Login from "./pages/Auth/Login";
import Signup from "./pages/Auth/Signup";
// import SeasonList from "./pages/Seasons/SeasonList";
// import SeasonDetail from "./pages/Seasons/SeasonDetail";
import "./styles/theme.ts";
import "./App.css";

function App() {
  return (
    <div className="App">
      <Router>
        <Header />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/signup" element={<Signup />} />
          {/* <Route path="/seasons" element={<SeasonList />} />
          <Route path="/seasons/:seasonId" element={<SeasonDetail />} /> */}
          <Route path="*" element={<NotFound />} />
        </Routes>
        <Footer />
      </Router>
    </div>
  );
}

export default App;
