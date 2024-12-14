import { Typography, Box } from "@mui/material";
import { useAuth } from "../../components/Auth/AuthContext";

const UserHomePage = () => {
  const { user } = useAuth();

  return (
    <Box sx={{ textAlign: "center", mt: 1, mb: 1 }}>
      <Typography variant="h3" gutterBottom>
        Welcome to MatchMania, {user?.username}!
      </Typography>
    </Box>
  );
};

export default UserHomePage;
