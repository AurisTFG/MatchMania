import { useState } from "react";
import { TextField, Button, Box } from "@mui/material";
import { signup } from "../../api/auth";

const Signup = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    try {
      const result = await signup(username, email, password);
      console.log("Login success:", result);
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  return (
    <>
      <Box sx={{ maxWidth: 400, mx: "auto", mt: 5 }}>
        <h1>Sign up</h1>
        <TextField
          fullWidth
          label="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          margin="normal"
        />
        <TextField
          fullWidth
          label="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          margin="normal"
        />
        <TextField
          fullWidth
          label="Password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          margin="normal"
        />
        <Button
          fullWidth
          variant="contained"
          color="primary"
          onClick={handleLogin}
        >
          Login
        </Button>
      </Box>
    </>
  );
};

export default Signup;
