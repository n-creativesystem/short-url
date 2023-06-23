import NotFound from '@/components/pages/notfound';
import { useOutletContext } from '@/pages/hooks/useOutlet';
import { FC, memo } from 'react';

const NotFoundPage: FC = memo(() => {
  const context = useOutletContext();
  context.setTitle('お探しのページは見つかりませんでした');
  return <NotFound />;
});

NotFoundPage.displayName = 'NotFoundPage';

export default NotFoundPage;
