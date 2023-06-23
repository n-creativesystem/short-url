import { OAuthAppContainer } from '@/components/pages/oauth/Register';
import { useOutletContext } from '@/pages/hooks/useOutlet';
import { FC } from 'react';

const Page: FC = () => {
  const { setTitle } = useOutletContext();
  setTitle('アプリケーション登録');
  return <OAuthAppContainer />;
};

Page.displayName = 'OAuthApplicationRegisterPage';

export default Page;
