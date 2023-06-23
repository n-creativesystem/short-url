import { usePathname } from '@/components/Parts/Navigation';
import OAuth2ApplicationsContainer from '@/components/pages/oauth';
import { useOutletContext } from '@/pages/hooks/useOutlet';
import { FC, memo } from 'react';
import { Outlet } from 'react-router-dom';

type Props = {};

const Index: FC<Props> = memo(() => {
  const context = useOutletContext();
  context.setTitle('アプリケーション管理');
  return <OAuth2ApplicationsContainer />;
});

Index.displayName = 'OAuth2AppIndex';

const NestedPage: FC<Props> = memo(() => {
  const context = useOutletContext();
  return <Outlet context={context} />;
});

NestedPage.displayName = 'OAuth2AppNestedPage';

const Page: FC<Props> = memo(() => {
  const pathname = usePathname();
  if (pathname === '/oauth2/app') {
    return <Index />;
  } else {
    return <NestedPage />;
  }
});

Page.displayName = 'OAuth2AppPage';

export default Page;
