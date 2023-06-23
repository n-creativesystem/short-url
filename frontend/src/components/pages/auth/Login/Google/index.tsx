import { AsyncButton } from '@/components/Parts/Button';
import { onClickEvent } from '@/components/Parts/Button/index.d';
import { FC } from 'react';
import { Normal } from './Logo';

type Props = {
  onClick: (e?: onClickEvent) => Promise<void>;
};

const GoogleLogin: FC<Props> = ({ onClick }) => {
  return (
    <AsyncButton loading={false} variant="text" size="large" onClick={onClick}>
      <Normal />
    </AsyncButton>
  );
};

export default GoogleLogin;
