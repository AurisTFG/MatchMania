// make signup page

import { Button, TextField, Typography, Box } from "@mui/material";
import { Link } from "react-router-dom";
import { useState } from "react";
import { useMutation } from "react-query";
import api from "../../api";
import { useAuth } from "../../context/AuthContext";

interface SignupResponse {
  // Define the shape of the response data from the signup API
  token: string;
  user: {
    id: string;
    email: string;
    username: string;
  };
}

interface SignupError {
  response: {
    data: {
      message: string;
    };
  };
}

const Signup = () => {
  const { login } = useAuth();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [username, setUsername] = useState("");
  const [error, setError] = useState("");

  const { mutate } = useMutation(api, {
    onSuccess: (data: SignupResponse) => {
      login(data);
    },
    onError: (error: SignupError) => {
      setError(error.response.data.message);
    },
  });

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    mutate({ email, password, username });
  };

  return (
    <Box sx={{ textAlign: "center", mt: 5 }}>
      <Typography variant="h4" gutterBottom>
        Sign Up
      </Typography>
      <form onSubmit={handleSubmit}>
        <TextField
          label="Email"
          type="email"
          variant="outlined"
          fullWidth
          sx={{ mt: 2 }}
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <TextField
          label="Username"
          type="text"
          variant="outlined"
          fullWidth
          sx={{ mt: 2 }}
          value={username}
          onChange={(e) => setUsername(e.target.value)}
        />
        <TextField
          label="Password"
          type="password"
          variant="outlined"
          fullWidth
          sx={{ mt: 2 }}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        {error && <Typography color="error">{error}</Typography>}
        <Button
          type="submit"
          variant="contained"
          color="primary"
          sx={{ mt: 2 }}
        >
          Sign Up
        </Button>
      </form>
      <Typography sx={{ mt: 2 }}>
        Already have an account? <Link to="/login">Login</Link>
      </Typography>
    </Box>
  );
};

export default Signup;
