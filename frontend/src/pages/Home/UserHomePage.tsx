import { Typography, Box } from "@mui/material";
import { UseAuth } from "../../components/Auth/AuthContext";

const UserHomePage = () => {
  const { user } = UseAuth();

  return (
    <Box sx={{ textAlign: "center" }}>
      <Typography
        variant="h3"
        gutterBottom
        sx={{
          wordWrap: "break-word",
          overflowWrap: "break-word",
          width: "100%",
          marginTop: "35vh",
        }}
      >
        Welcome to MatchMania,
        <br />
        {user?.username}!
      </Typography>
    </Box>
  );
};

export default UserHomePage;
