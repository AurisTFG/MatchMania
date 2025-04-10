import { Box, Button, TextField } from "@mui/material";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useLogIn } from "../../api/hooks/authHooks";
import { User } from "../../types/users";

export default function Login() {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  const { mutateAsync: loginMutation } = useLogIn();

  const handleLogin = async () => {
    setLoading(true);

    const user = (await loginMutation({ email, password })) as User;

    if (user.id) {
      await navigate("/");
    } else {
      console.error("User not found");
    }

    setLoading(false);
  };

  return (
    <Box sx={{ maxWidth: 400, mx: "auto", mt: 5 }}>
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
        disabled={loading}
      >
        {loading ? "Logging in..." : "Login"}
      </Button>
    </Box>
  );
}
