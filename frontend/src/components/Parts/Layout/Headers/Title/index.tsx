import { FC, ReactNode, memo } from 'react';
import style from './index.module.scss';

type Props = {
  children: ReactNode;
};

export const Title: FC<Props> = memo(({ children }) => {
  return <h3 className={style.title}>{children}</h3>;
});

if (process.env.NODE_ENV === 'production') {
  Title.displayName = 'ContentTitle';
}
