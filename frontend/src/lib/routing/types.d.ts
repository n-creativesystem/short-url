import type { ReadonlyDeep } from 'type-fest';

export type RouteProps = {
  path: string;
  auth?: boolean;
  Component?: LazyExoticComponent<FC>;
  routes?: RouteProps[];
} & Omit<BaseRouteProps, 'children' | 'element'>;

export type RoutesDef = ReadonlyArray<ReadonlyDeep<RouteProps>>;
