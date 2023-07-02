import { useOutletContext } from '@/components/hooks/useOutlet';
import { OAuthAppContainer } from '@/components/pages/oauth/Updater';
import { FC } from 'react';

const Page: FC = () => {
  const { setTitle } = useOutletContext();
  setTitle('アプリケーション編集');
  return <OAuthAppContainer />;
};

Page.displayName = 'OAuthApplicationRegisterPage';

export default Page;
