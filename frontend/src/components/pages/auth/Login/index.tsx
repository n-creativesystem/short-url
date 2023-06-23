import { FC, memo } from 'react';
import Presenter from './Presenter';
import { useAuthLogin, useEnabledAuth } from './hooks';
import { ContainerProps } from './index.d';

const LoginContainer: FC<ContainerProps> = memo(() => {
  const enabledAuth = useEnabledAuth();
  return <Presenter useAuthLogin={useAuthLogin} enabledAuth={enabledAuth} />;
});

LoginContainer.displayName = '';

export default LoginContainer;
