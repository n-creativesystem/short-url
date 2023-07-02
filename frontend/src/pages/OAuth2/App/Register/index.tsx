import { useOutletContext } from '@/components/hooks/useOutlet';
import { OAuthAppContainer } from '@/components/pages/oauth/Register';
import { FC } from 'react';

const Page: FC = () => {
  const { setTitle } = useOutletContext();
  setTitle('アプリケーション登録');
  return <OAuthAppContainer />;
};

Page.displayName = 'OAuthApplicationRegisterPage';

export default Page;
