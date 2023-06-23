import {
  default as MuiAppBar,
  AppBarProps as MuiAppBarProps,
} from '@mui/material/AppBar';
import Link from '@mui/material/Link';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import { styled } from '@mui/material/styles';
import { ReactNode } from 'react';

interface MyAppBarProps extends MuiAppBarProps {
  title?: string;
  auth?: ReactNode;
}

const MyAppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== 'open',
})<MyAppBarProps>(({ theme }) => {
  return {
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  };
});

export const AppBar = ({ title, auth }: MyAppBarProps) => {
  return (
    <MyAppBar position="absolute">
      <Toolbar>
        <Typography variant="h4" component="div" sx={{ flexGrow: 1 }}>
          <Link href="/" component="button" underline="none" color="inherit">
            {title}
          </Link>
        </Typography>
        {auth && <div>{auth}</div>}
      </Toolbar>
    </MyAppBar>
  );
};
