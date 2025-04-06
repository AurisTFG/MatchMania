import AccountCircle from "@mui/icons-material/AccountCircle";
import MenuIcon from "@mui/icons-material/Menu";
import { useMediaQuery } from "@mui/material";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Drawer from "@mui/material/Drawer";
import IconButton from "@mui/material/IconButton";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { logout } from "../../api/auth";
import { UseAuth } from "../Auth/AuthContext";

export default function Header() {
  const { user, setUser } = UseAuth();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const [mobileMenuOpen, setMobileMenuOpen] = useState<boolean>(false);
  const navigate = useNavigate();
  const location = useLocation();
  const isActive = (path: string) => location.pathname === path;

  const handleLogout = () => {
    console.log("Logging out...");
    setAnchorEl(null);
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
    <AppBar
      position="static"
      sx={{
        background: "linear-gradient(to right,rgb(64, 112, 145), #9b59b6)",
      }}
    >
      <Toolbar>
        {isMobile && (
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
            onClick={() => {
              setMobileMenuOpen(true);
            }}
          >
            <MenuIcon />
          </IconButton>
        )}

        <Typography
          variant="h6"
          component="div"
          sx={{ flexGrow: 1, display: "flex", alignItems: "center" }}
        >
          <img
            src="/car_icon.svg"
            alt="Logo"
            style={{ height: "40px", marginRight: "8px" }}
          />
          MatchMania
        </Typography>

        {!isMobile && (
          <Box
            sx={{ flexGrow: 1, mr: 10, display: { xs: "none", md: "flex" } }}
          >
            <Button
              onClick={() => navigate("/")}
              color="inherit"
              sx={{
                mr: 1,
                textDecoration: isActive("/") ? "underline" : "none",
                fontSize: "1.20rem",
              }}
            >
              Home
            </Button>
            <Button
              onClick={() => navigate("/seasons")}
              color="inherit"
              sx={{
                mr: 1,
                textDecoration: isActive("/seasons") ? "underline" : "none",
                fontSize: "1.20rem",
              }}
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

        {/* Mobile Menu Drawer */}
        <Drawer
          anchor="left"
          open={mobileMenuOpen}
          onClose={() => {
            setMobileMenuOpen(false);
          }}
          sx={{
            "& .MuiDrawer-paper": {
              width: "250px",
              height: "100vh",
              display: "flex",
              flexDirection: "column",
              justifyContent: "space-between",
              padding: "20px",
            },
          }}
        >
          <Box sx={{ display: "flex", flexDirection: "column", gap: "16px" }}>
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
          </Box>
        </Drawer>
      </Toolbar>
    </AppBar>
  );
}
