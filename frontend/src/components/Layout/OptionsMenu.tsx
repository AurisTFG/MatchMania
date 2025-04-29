import LogoutRoundedIcon from '@mui/icons-material/LogoutRounded';
import MoreVertRoundedIcon from '@mui/icons-material/MoreVertRounded';
import { IconButton } from '@mui/material';
import Divider, { dividerClasses } from '@mui/material/Divider';
import { listClasses } from '@mui/material/List';
import ListItemIcon, { listItemIconClasses } from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import Menu from '@mui/material/Menu';
import MuiMenuItem from '@mui/material/MenuItem';
import { paperClasses } from '@mui/material/Paper';
import { styled } from '@mui/material/styles';
import * as React from 'react';
import { useNavigate } from 'react-router-dom';
import { useLogOut } from 'api/hooks/authHooks';
import { ROUTE_PATHS } from 'constants/route_paths';

const MenuItem = styled(MuiMenuItem)({
  margin: '2px 0',
});

export default function OptionsMenu() {
  const navigate = useNavigate();
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const { mutateAsync: logoutAsync } = useLogOut();

  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleProfileClick = async () => {
    await navigate(ROUTE_PATHS.PROFILE);
  };

  const handleLogout = async () => {
    setAnchorEl(null);
    await logoutAsync();
  };

  return (
    <>
      <IconButton
        size="small"
        aria-label="Open menu"
        onClick={handleClick}
        sx={{
          borderRadius: '10px',
          backgroundColor: 'action.hover',
          border: '1px solid',
          borderColor: 'divider',
          '&:hover': {
            backgroundColor: 'action.selected',
          },
        }}
      >
        <MoreVertRoundedIcon />
      </IconButton>
      <Menu
        anchorEl={anchorEl}
        id="menu"
        open={open}
        onClose={handleClose}
        onClick={handleClose}
        transformOrigin={{ horizontal: 'right', vertical: 'top' }}
        anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
        sx={{
          [`& .${listClasses.root}`]: {
            padding: '4px',
          },
          [`& .${paperClasses.root}`]: {
            padding: 0,
          },
          [`& .${dividerClasses.root}`]: {
            margin: '4px -4px',
          },
        }}
      >
        <MenuItem
          onClick={() => {
            void handleProfileClick();
          }}
        >
          Profile
        </MenuItem>
        <Divider />
        <MenuItem
          onClick={() => {
            void handleLogout();
          }}
          sx={{
            [`& .${listItemIconClasses.root}`]: {
              ml: 'auto',
              minWidth: 0,
            },
            minWidth: 120,
            display: 'flex',
            justifyContent: 'space-between',
            paddingRight: 1,
          }}
        >
          <ListItemText>Logout</ListItemText>
          <ListItemIcon>
            <LogoutRoundedIcon fontSize="small" />
          </ListItemIcon>
        </MenuItem>
      </Menu>
    </>
  );
}
