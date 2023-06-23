import Stack from '@mui/material/Stack';
import { ResponsiveStyleValue, SxProps } from '@mui/system';
import { FC, FormEventHandler, ReactNode } from 'react';

type Props = {
  children?: ReactNode;
  noValidate?: boolean;
  onSubmit?: FormEventHandler;
  spacing?: ResponsiveStyleValue<number | string>;
  sx?: SxProps;
  autoComplete?: string;
};

export const Form: FC<Props> = ({
  children,
  autoComplete = 'off',
  ...props
}) => {
  return (
    <Stack {...props} component="form" autoComplete={autoComplete}>
      {children}
    </Stack>
  );
};
