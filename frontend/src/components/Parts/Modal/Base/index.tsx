import BaseModal from '@mui/material/Modal';
import { FC, ReactElement } from 'react';

type Props = {
  open: boolean;
  onClose: () => {};
  children: ReactElement;
};

export const Modal: FC<Props> = ({ open, onClose, children }) => {
  return (
    <BaseModal open={open} onClose={onClose}>
      {children}
    </BaseModal>
  );
};
