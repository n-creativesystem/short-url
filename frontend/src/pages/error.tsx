import { internalErrorPageVar } from '@/components/hooks/Context';
import { useReactiveVar } from '@/components/hooks/reactive';
import { FC, lazy, memo } from 'react';
import Routing, { RouteProps } from './routing';

const children: RouteProps[] = [
  {
    path: '*',
    Component: lazy(() => import('./Error/500')),
  },
];

const ErrorPages: FC = memo(() => {
  const value = useReactiveVar(internalErrorPageVar);
  if (!value.show) {
    return <></>;
  }
  return <Routing children={children} />;
});

ErrorPages.displayName = 'ErrorPages';

export default ErrorPages;
