import LoadingButton, { LoadingButtonProps } from '@mui/lab/LoadingButton';
import classnames from 'classnames';
import { FC, forwardRef } from 'react';
import styles from './button.module.scss';

const cx = classnames.bind(styles);

export type ButtonProps = {} & LoadingButtonProps;

export const Button: FC<ButtonProps> = forwardRef(({ ...props }, ref) => {
  return (
    <LoadingButton
      {...props}
      className={cx(styles['button-override'])}
      ref={ref}
    />
  );
});
