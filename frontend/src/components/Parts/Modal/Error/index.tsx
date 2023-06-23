import { errorModalVar, initialErrorModal } from '@/components/hooks/Context';
import { useReactiveVar } from '@/components/hooks/reactive';
import WarningIcon from '@mui/icons-material/WarningAmberOutlined';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import { FC } from 'react';
import { Modal } from '../Base';

export const ErrorModal: FC = () => {
  const value = useReactiveVar(errorModalVar);
  const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    boxShadow: 24,
    p: 4,
  };
  return (
    <Modal open={value.open} onClose={() => errorModalVar(initialErrorModal)}>
      <Box sx={{ ...style, maxWidth: 600 }}>
        <Typography sx={{ color: 'red' }} variant="h4" component="h2">
          <WarningIcon /> {value.title}
        </Typography>
        <Typography variant="subtitle1" sx={{ mt: 2, color: 'red' }}>
          {value.description}
        </Typography>
      </Box>
    </Modal>
  );
};
