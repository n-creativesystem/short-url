declare module '*.svg' {
  import React = require('react');
  export const ReactComponent: React.FC<React.SVGProps<SVGSVGElement>>;
  const src: string;
  export default src;
}

type Expand<T> = T extends object
  ? T extends infer O
    ? { [K in keyof O]: Expand<O[K]> }
    : never
  : T;

type CustomTime = string & { readonly brand: unique symbol };
type CustomURL = URL & { readonly brand: unique symbol };
type TitleProps = {
  setTitle: (value: string) => void;
};
