import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import MenuIcon from "@mui/icons-material/Menu";
import IconButton from "@mui/material/IconButton";
import AccountCircle from "@mui/icons-material/AccountCircle";
import MenuItem from "@mui/material/MenuItem";
import Menu from "@mui/material/Menu";
import { UseAuth } from "../Auth/AuthContext";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { logout } from "../../api/auth";
import { useMediaQuery } from "@mui/material";

export default function Header() {
  const { user, setUser } = UseAuth();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [mobileMenuOpen, setMobileMenuOpen] = useState<boolean>(false);
  const navigate = useNavigate();

  const handleLogout = () => {
    console.log("Logging out...");
    logout();
    setUser(null);
    navigate("/");
  };

  const handleProfile = () => {
    console.log("Viewing profile...");
    navigate("/profile");
  };

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const isMobile = useMediaQuery("(max-width: 768px)");

  return (
    <AppBar position="static">
      <Toolbar>
        {isMobile && (
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
            onClick={() => setMobileMenuOpen(true)}
          >
            <MenuIcon />
          </IconButton>
        )}

        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
          MatchMania
        </Typography>

        {!isMobile && user && (
          <Box sx={{ flexGrow: 1, display: { xs: "none", md: "flex" } }}>
            <Button
              color="inherit"
              onClick={() => navigate("/")}
              sx={{ mr: 1 }}
            >
              Home
            </Button>
            <Button
              color="inherit"
              onClick={() => navigate("/seasons")}
              sx={{ mr: 1 }}
            >
              Seasons
            </Button>
          </Box>
        )}

        {user && (
          <div>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleMenu}
              color="inherit"
            >
              <AccountCircle />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorEl}
              anchorOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              open={Boolean(anchorEl)}
              onClose={handleClose}
            >
              <MenuItem onClick={handleProfile}>Profile</MenuItem>
              <MenuItem onClick={handleLogout}>Logout</MenuItem>
            </Menu>
          </div>
        )}

        {isMobile && mobileMenuOpen && (
          <Box
            sx={{
              position: "absolute",
              top: "64px",
              right: 0,
              backgroundColor: "white",
              boxShadow: 3,
              zIndex: 1300,
            }}
          >
            <Button
              color="inherit"
              onClick={() => {
                navigate("/");
                setMobileMenuOpen(false);
              }}
            >
              Home
            </Button>
            <Button
              color="inherit"
              onClick={() => {
                navigate("/seasons");
                setMobileMenuOpen(false);
              }}
            >
              Seasons
            </Button>
            <Button color="inherit" onClick={handleLogout}>
              Logout
            </Button>
            <Button color="inherit" onClick={handleProfile}>
              Profile
            </Button>
          </Box>
        )}
      </Toolbar>
    </AppBar>
  );
}
