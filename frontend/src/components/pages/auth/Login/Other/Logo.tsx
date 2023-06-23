import LoginIcon from '@/styles/icons/login.svg';
import { FC } from 'react';

type SafeNumber = number | `${number}`;

export type Props = {
  width?: SafeNumber | undefined;
  height?: SafeNumber | undefined;
};

export const Login: FC<Props> = (props) => (
  <img src={LoginIcon} alt="" {...props} />
);

export default Login;
