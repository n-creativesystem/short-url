import TextField, { TextFieldProps } from '@mui/material/TextField';
import { FC, forwardRef } from 'react';

type Props = {
  id?: string;
  name?: string;
  label: string;
} & TextFieldProps;

export const Input: FC<Props> = forwardRef(
  ({ variant = 'standard', size = 'small', ...props }, ref) => {
    return <TextField {...props} ref={ref} variant="standard" size={size} />;
  }
);
