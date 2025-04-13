import { Box, Button, TextField } from '@mui/material';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useSignUp } from '../../api/hooks/authHooks';

export default function SignupPage() {
  const navigate = useNavigate();
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const { mutateAsync: signupMutation, isPending, error } = useSignUp();

  const handleSignup = async () => {
    try {
      await signupMutation({ username, email, password });
      await navigate('/');
    } catch (error) {
      console.error('Signup failed:', error);
    }
  };

  return (
    <Box sx={{ maxWidth: 400, mx: 'auto', mt: 5 }}>
      <h1>Sign up</h1>
      <TextField
        fullWidth
        label="Username"
        value={username}
        onChange={(e) => {
          setUsername(e.target.value);
        }}
        margin="normal"
      />
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
        onClick={handleSignup}
        disabled={isPending}
      >
        {isPending ? 'Signing up...' : 'Sign up'}
      </Button>
      {error && (
        <Box sx={{ color: 'red', mt: 2 }}>
          <p>Signup failed: {error.message}</p>
        </Box>
      )}
    </Box>
  );
}
