import LoginContainer from '@/components/pages/auth/Login';
import { FC, memo } from 'react';

type Props = {};

const Page: FC<Props> = memo(() => {
  return <LoginContainer />;
});

Page.displayName = 'AuthPage';

export default Page;
