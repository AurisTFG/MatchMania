import { useAuth } from '../../providers/AuthProvider/AuthProvider';

export default function ProfilePage() {
  const { user } = useAuth();

  return (
    <div>
      <h1>Profile</h1>
      <p>Username: {user?.username}</p>
      <p>Email: {user?.email}</p>
    </div>
  );
}
