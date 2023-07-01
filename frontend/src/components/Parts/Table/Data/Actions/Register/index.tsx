import { Button, ButtonProps } from '@/components/Parts/Button';
import classNames from 'classnames';
import { FC } from 'react';
import styles from './index.module.scss';

const cx = classNames.bind(styles);

type Props = {
  handler: ButtonProps['onClick'];
  className?: string;
};

export const RegisterButton: FC<Props> = ({ handler, className }) => {
  return (
    <div className={cx(styles['button'], className)}>
      <Button variant="contained" color="primary" onClick={handler}>
        新規作成
      </Button>
    </div>
  );
};
