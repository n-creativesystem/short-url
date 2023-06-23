import FormControl from '@mui/material/FormControl';
import { FC, ReactNode } from 'react';

type Props = {
  children: ReactNode[];
};

export const Form: FC<Props> = ({ children }) => {
  return <FormControl>{children.map((child) => child)}</FormControl>;
};
