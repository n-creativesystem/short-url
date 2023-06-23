import { OAuthAppContainer } from '@/components/pages/oauth/Updater';
import { useOutletContext } from '@/pages/hooks/useOutlet';
import { FC } from 'react';

const Page: FC = () => {
  const { setTitle } = useOutletContext();
  setTitle('アプリケーション編集');
  return <OAuthAppContainer />;
};

Page.displayName = 'OAuthApplicationRegisterPage';

export default Page;
