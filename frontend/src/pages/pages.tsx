import CsrfTokenProvider from '@/components/Parts/Layout/CsrfToken';
import UserInfoProvider from '@/components/Parts/Layout/UserInfo';
import { internalErrorPageVar } from '@/components/hooks/Context';
import { useReactiveVar } from '@/components/hooks/reactive';
import Routing from '@/lib/routing';
import { RouteProps } from '@/lib/routing/types.d';
import { FC, lazy, memo } from 'react';

const routes: RouteProps[] = [
  {
    path: '',
    Component: lazy(() => import('./index')),
  },
  {
    path: 'auth',
    Component: lazy(() => import('./Auth')),
  },
  {
    path: 'oauth2/app',
    Component: lazy(() => import('./OAuth2/App')),
    auth: true,
    routes: [
      {
        path: 'register',
        Component: lazy(() => import('./OAuth2/App/Register')),
      },
      {
        path: ':id',
        Component: lazy(() => import('./OAuth2/App/[id]')),
      },
    ],
  },
  {
    path: 'shorts',
    Component: lazy(() => import('./Shorts')),
    auth: true,
    routes: [],
  },
  {
    path: '*',
    Component: lazy(() => import('./Error/404')),
  },
];

const Pages: FC = memo(() => {
  const internalError = useReactiveVar(internalErrorPageVar);
  if (internalError.show) {
    return <></>;
  }
  return (
    <CsrfTokenProvider>
      <UserInfoProvider>
        <Routing children={routes} />
      </UserInfoProvider>
    </CsrfTokenProvider>
  );
});

Pages.displayName = 'Pages';

export default Pages;
