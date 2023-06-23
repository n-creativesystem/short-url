import Box from '@mui/material/Box';
import { styled } from '@mui/material/styles';

export const Item = styled(Box)(({ theme }) => ({
  padding: theme.spacing(1),
  textAlign: 'center',
}));
