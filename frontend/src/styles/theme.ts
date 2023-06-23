import style from '@/styles/variables/color.module.scss';
import { LinkProps } from '@mui/material/Link';
import { createTheme } from '@mui/material/styles';
import { LinkBehavior } from './link';

const theme = createTheme({
  components: {
    MuiLink: {
      defaultProps: {
        component: LinkBehavior,
      } as LinkProps,
    },
    MuiButtonBase: {
      defaultProps: {
        LinkComponent: LinkBehavior,
      },
    },
  },
  palette: {
    primary: {
      main: style['primaryColor-v1_primary'],
    },
    secondary: {
      main: style['secondaryColor-v1_secondary'],
    },
    error: {
      main: style['errorColor-v1_error'],
    },
  },
});

export default theme;
