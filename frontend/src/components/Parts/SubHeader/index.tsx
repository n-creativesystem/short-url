import { FC, memo } from 'react';
import styles from './index.module.scss';

type Props = {
  title: string;
};

export const SubHeader: FC<Props> = memo(({ title }) => {
  if (!title) return <></>;
  return (
    <div className={styles.container}>
      <h2 className={styles.title} data-testid="page-title">
        {title}
      </h2>
    </div>
  );
});

SubHeader.displayName = 'SubHeader';

export default memo(SubHeader);
