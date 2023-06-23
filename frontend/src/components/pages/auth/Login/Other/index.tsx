import {
  AsyncButton,
  Props as AsyncButtonType,
} from '@/components/Parts/Button';
import { onClickEvent } from '@/components/Parts/Button/index.d';
import { FC } from 'react';
import LoginIcon, { Props as LoginIconProps } from './Logo';

type Props = {
  label: string;
  onClick: (e?: onClickEvent) => Promise<void>;
  button?: Omit<AsyncButtonType, 'onClick'>;
  icon?: LoginIconProps;
};

const OtherAuthButton: FC<Props> = ({ label, button, icon, onClick }) => {
  return (
    <AsyncButton variant="outlined" size="large" onClick={onClick} {...button}>
      <LoginIcon {...icon} />{' '}
      <span style={{ marginLeft: '40px' }}>Sign in with {label}</span>
    </AsyncButton>
  );
};

export default OtherAuthButton;
