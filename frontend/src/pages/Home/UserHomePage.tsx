import { Typography, Box } from "@mui/material";
import { UseAuth } from "../../components/Auth/AuthContext";

const UserHomePage = () => {
  const { user } = UseAuth();

  return (
    <Box
      sx={{
        textAlign: "center",
        height: "83.5vh",
        backgroundImage: 'url("https://i.imgur.com/2k81qri.jpeg")',
        backgroundSize: "cover",
        backgroundPosition: "center",
        backgroundRepeat: "no-repeat",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <Typography
        variant="h3"
        gutterBottom
        sx={{
          wordWrap: "break-word",
          overflowWrap: "break-word",
          width: "100%",
          color: "rgba(94, 140, 192, 0.9)",
          fontSize: "6rem",
          textShadow: "5px 5px 5px rgba(0, 0, 0, 0.8)",
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
