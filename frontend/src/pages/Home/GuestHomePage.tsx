import { Button, Typography, Box } from "@mui/material";
import { Link } from "react-router-dom";

const GuestHomePage = () => {
  return (
    <Box sx={{ textAlign: "center", mt: 1, mb: 1 }}>
      <Typography variant="h3" gutterBottom fontFamily={"Roboto Condensed"}>
        Welcome to MatchMania!
      </Typography>
      <Typography variant="h6" gutterBottom>
        A matchmaking platform for Trackmania
      </Typography>
      <Box sx={{ mt: 3 }}>
        <Button
          component={Link}
          to="/login"
          variant="contained"
          color="primary"
          sx={{ mx: 1 }}
        >
          Login
        </Button>
        <Button
          component={Link}
          to="/signup"
          variant="outlined"
          color="secondary"
          sx={{ mx: 1 }}
        >
          Sign Up
        </Button>
      </Box>
    </Box>
  );
};

export default GuestHomePage;
