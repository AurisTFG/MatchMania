import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { TextField, Button, Box } from "@mui/material";
import { useAuth } from "../../components/Auth/AuthContext";
import { login } from "../../api/auth";
import { getCurrentUser } from "../../api/users";

const Login = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { setUser } = useAuth();

  const handleLogin = async () => {
    try {
      const result = await login(email, password);
      const user = await getCurrentUser();

      console.log("Login success:", result, user);

      setUser(user);

      navigate("/");
    } catch (error) {
      console.error("Login failed:", error);
    }
  };

  return (
    <>
      <Box sx={{ maxWidth: 400, mx: "auto", mt: 5 }}>
        <h1>Login</h1>
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

export default Login;
