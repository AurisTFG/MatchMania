import { Button } from "@mui/material";
import { UseAuth } from "../../components/Auth/AuthContext";

const Profile = () => {
  const { user } = UseAuth();

  const handleEditProfile = () => {
    console.log(user);
    console.log("Editing profile...");
  };

  return (
    <div>
      <h1>Profile</h1>
      <p>Username: {user?.username}</p>
      <p>Email: {user?.email}</p>
      <Button variant="contained" color="primary" onClick={handleEditProfile}>
        Edit Profile
      </Button>
    </div>
  );
};

export default Profile;
