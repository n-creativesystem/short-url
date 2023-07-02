import { Skeleton } from '@/components/Parts/Loading';
import { LoadingContext } from '@/components/hooks/Context';
import LoadingButton, { LoadingButtonProps } from '@mui/lab/LoadingButton';
import classnames from 'classnames';
import { FC, forwardRef, useContext } from 'react';
import styles from './button.module.scss';

const cx = classnames.bind(styles);

export type ButtonProps = {} & LoadingButtonProps;

export const Button: FC<ButtonProps> = forwardRef(({ ...props }, ref) => {
  const loading = useContext(LoadingContext);
  return loading ? (
    <>
      <Skeleton height={30} />
    </>
  ) : (
    <LoadingButton
      {...props}
      className={cx(styles['button-override'])}
      ref={ref}
    />
  );
});
