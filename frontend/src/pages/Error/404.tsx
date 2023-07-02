import { useOutletContext } from '@/components/hooks/useOutlet';
import NotFound from '@/components/pages/notfound';
import { FC, memo } from 'react';

const NotFoundPage: FC = memo(() => {
  const context = useOutletContext();
  context.setTitle('お探しのページは見つかりませんでした');
  return <NotFound />;
});

NotFoundPage.displayName = 'NotFoundPage';

export default NotFoundPage;
