import { usePathname } from '@/components/Parts/Navigation';
import varStyles from '@/styles/variables/index.module.scss';
import SettingsApplicationsIcon from '@mui/icons-material/SettingsApplications';
import MuiDrawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import MuiListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import { styled } from '@mui/material/styles';
import React from 'react';
import {
  Link as RouterLink,
  LinkProps as RouterLinkProps,
} from 'react-router-dom';

const Drawer = styled(MuiDrawer, {
  shouldForwardProp: (prop) => prop !== 'open',
})(({ theme, open }) => ({
  '& .MuiDrawer-paper': {
    position: 'relative',
    whiteSpace: 'nowrap',
    width: varStyles.drawerWidth,
    transition: theme.transitions.create('width', {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
    boxSizing: 'border-box',
    ...(!open && {
      overflowX: 'hidden',
      transition: theme.transitions.create('width', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      width: theme.spacing(7),
      [theme.breakpoints.up('sm')]: {
        width: theme.spacing(9),
      },
    }),
  },
}));

type ListItemButtonProps = {
  selected?: boolean;
} & React.PropsWithChildren<RouterLinkProps>;

const ListItemButton = (props: ListItemButtonProps) => {
  const { to, children, selected, ...rest } = props;
  return (
    <MuiListItemButton
      selected={selected}
      component={RouterLink}
      to={to}
      {...rest}
    >
      {children}
    </MuiListItemButton>
  );
};

export const MenuDrawer = () => {
  const pathname = usePathname();
  return (
    <Drawer
      sx={{
        mt: varStyles.appBarHeight,
      }}
      variant="permanent"
      open={true}
    >
      <List component="nav">
        {/* <ListItemButton to="/" selected={pathname === '/'}>
          <ListItemIcon>
            <HomeIcon />
          </ListItemIcon>
          <ListItemText primary="トップページ" />
        </ListItemButton> */}
        <ListItemButton
          to="/oauth2/app"
          selected={pathname.startsWith('/oauth2/app')}
        >
          <ListItemIcon>
            <SettingsApplicationsIcon />
          </ListItemIcon>
          <ListItemText primary="アプリケーション" />
        </ListItemButton>
      </List>
    </Drawer>
  );
};
