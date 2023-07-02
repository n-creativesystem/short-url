import { LoadingContext } from '@/components/hooks/Context';
import { Skeleton } from '@/components/Parts/Loading';
import TextField, { TextFieldProps } from '@mui/material/TextField';
import { FC, forwardRef, useContext } from 'react';

type Props = {
  id?: string;
  name?: string;
  label: string;
} & TextFieldProps;

export const Input: FC<Props> = forwardRef(
  ({ variant = 'standard', size = 'small', ...props }, ref) => {
    const loading = useContext(LoadingContext);
    return loading ? (
      <Skeleton />
    ) : (
      <TextField {...props} ref={ref} variant="standard" size={size} />
    );
  }
);
