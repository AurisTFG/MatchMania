import { Typography, Box } from "@mui/material";
import { UseAuth } from "../../components/Auth/AuthContext";

const UserHomePage = () => {
  const { user } = UseAuth();

  return (
    <Box sx={{ textAlign: "center", mt: 1, mb: 1 }}>
      <Typography variant="h3" gutterBottom>
        Welcome to MatchMania, {user?.username}!
      </Typography>
    </Box>
  );
};

export default UserHomePage;
