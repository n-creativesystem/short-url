import { usePathname } from '@/components/Parts/Navigation';
import varStyles from '@/styles/variables/index.module.scss';
import MuiDrawer from '@mui/material/Drawer';
import List from '@mui/material/List';
import MuiListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import { styled } from '@mui/material/styles';
import React, { FC, ReactNode } from 'react';
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
  path?: string;
} & React.PropsWithChildren<RouterLinkProps>;

const ListItemButton: FC<ListItemButtonProps> = ({
  to,
  children,
  selected,
  path,
  ...rest
}) => {
  const pathname = usePathname();
  selected = selected || (path && pathname.startsWith(path)) || false;
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

type MenuProps = {
  children?: ReactNode;
};

const MenuDrawer: FC<MenuProps> = ({ children }) => {
  return (
    <Drawer
      sx={{
        mt: varStyles.appBarHeight,
      }}
      variant="permanent"
      open={true}
    >
      <List component="nav">{children}</List>
    </Drawer>
  );
};

type TListItemButton = typeof ListItemButton & {
  ListItemIcon: typeof ListItemIcon;
  ListItemText: typeof ListItemText;
};

type TMenu = typeof MenuDrawer & {
  ListItemButton: TListItemButton;
};

const Menu = MenuDrawer as TMenu;
Menu.ListItemButton = ListItemButton as TListItemButton;
Menu.ListItemButton.ListItemIcon = ListItemIcon;
Menu.ListItemButton.ListItemText = ListItemText;

export { Menu };
