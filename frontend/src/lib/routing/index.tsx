import Layout from '@/components/Parts/Layout';
import { FC, LazyExoticComponent, memo } from 'react';
import {
  RouteProps as BaseRouteProps,
  BrowserRouter,
  Route,
  Routes,
} from 'react-router-dom';
import { Auth } from './authGuard';

export type RouteProps = {
  path: string;
  auth?: boolean;
  Component?: LazyExoticComponent<FC>;
  routes?: RouteProps[];
} & Omit<BaseRouteProps, 'children' | 'element'>;

type Props = {
  children: RouteProps[];
  layout?: boolean;
};

const routers = (children: RouteProps[]) => {
  return children.map((child) => {
    const nestedRouters =
      child.routes &&
      routers(
        child.routes.map((item) => {
          return {
            ...item,
            ...(child.auth ? { auth: true } : {}),
          };
        })
      );
    return child.auth ? (
      <Route
        path={child.path}
        key={child.path}
        element={<Auth key={child.path} Component={child.Component} />}
      >
        {nestedRouters}
      </Route>
    ) : (
      <Route path={child.path} key={child.path} Component={child.Component}>
        {nestedRouters}
      </Route>
    );
  });
};

const Routing: FC<Props> = memo(({ children, layout = true }) => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={layout && <Layout />}>
          {routers(children)}
        </Route>
      </Routes>
    </BrowserRouter>
  );
});

Routing.displayName = 'Routing';

export default Routing;
