import CsrfTokenProvider from '@/components/Parts/Layout/CsrfToken';
import UserInfoProvider from '@/components/Parts/Layout/UserInfo';
import { internalErrorPageVar } from '@/components/hooks/Context';
import { useReactiveVar } from '@/components/hooks/reactive';
import { FC, lazy, memo } from 'react';
import Routing, { RouteProps } from './routing';

const children: RouteProps[] = [
  {
    path: '',
    Component: lazy(() => import('./App')),
  },
  {
    path: 'auth',
    Component: lazy(() => import('./auth')),
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
        <Routing children={children} />
      </UserInfoProvider>
    </CsrfTokenProvider>
  );
});

Pages.displayName = 'Pages';

export default Pages;
