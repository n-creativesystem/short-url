import { internalErrorPageVar } from '@/components/hooks/Context';
import { useReactiveVar } from '@/components/hooks/reactive';
import Routing from '@/lib/routing';
import type { RouteProps } from '@/lib/routing/types';
import { FC, lazy, memo } from 'react';

const children: RouteProps[] = [
  {
    path: '*',
    Component: lazy(() => import('@/pages/Error/500')),
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
