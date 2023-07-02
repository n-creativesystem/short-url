export type RouteProps = {
  path: string;
  auth?: boolean;
  Component?: LazyExoticComponent<FC>;
  routes?: RouteProps[];
} & Omit<BaseRouteProps, 'children' | 'element'>;
