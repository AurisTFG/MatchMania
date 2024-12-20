import { UseAuth } from "../../components/Auth/AuthContext";

const Profile = () => {
  const { user } = UseAuth();

  return (
    <div>
      <h1>Profile</h1>
      <p>Username: {user?.username}</p>
      <p>Email: {user?.email}</p>
    </div>
  );
};

export default Profile;
