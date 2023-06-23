import { FC, ReactNode } from 'react';
import { Link as BaseLink } from 'react-router-dom';

type Props = {
  to: string;
  children?: ReactNode;
};

export const Link: FC<Props> = ({ to, children }) => {
  return (
    <BaseLink to={to} replace={false}>
      {children && children}
    </BaseLink>
  );
};
