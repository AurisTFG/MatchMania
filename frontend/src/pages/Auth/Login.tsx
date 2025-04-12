import { Box, Button, TextField } from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useLogIn } from '../../api/hooks/authHooks';

export default function Login() {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const { mutateAsync: loginAsync, isPending: loginPending } = useLogIn();

  const handleLogin = async () => {
    await loginAsync({ email, password });

    await navigate('/');
  };

  return (
    <Box sx={{ maxWidth: 400, mx: 'auto', mt: 5 }}>
      <h1>Login</h1>
      <TextField
        fullWidth
        label="Email"
        value={email}
        onChange={(e) => {
          setEmail(e.target.value);
        }}
        margin="normal"
      />
      <TextField
        fullWidth
        label="Password"
        type="password"
        value={password}
        onChange={(e) => {
          setPassword(e.target.value);
        }}
        margin="normal"
      />
      <Button
        fullWidth
        variant="contained"
        color="primary"
        // eslint-disable-next-line @typescript-eslint/no-misused-promises
        onClick={handleLogin}
        disabled={loginPending}
      >
        {loginPending ? 'Logging in...' : 'Login'}
      </Button>
    </Box>
  );
}
