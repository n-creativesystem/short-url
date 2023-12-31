import { ErrorModal } from '@/components/Parts/Modal';
import Pages from '@/pages/pages';
import theme from '@/styles/theme.ts';
import { Box } from '@mui/material';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import { FC } from 'react';
import ErrorPages from './error';

export const App: FC = () => {
  return (
    <ThemeProvider theme={theme}>
      <Box sx={{ display: 'flex' }}>
        <CssBaseline />
        <Pages />
        <ErrorModal />
        <ErrorPages />
      </Box>
    </ThemeProvider>
  );
};
