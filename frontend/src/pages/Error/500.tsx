import { InternalError } from '@/components/pages/error';
import { FC, memo } from 'react';

const InternalErrorPage: FC = memo(() => {
  return <InternalError />;
});

InternalErrorPage.displayName = 'InternalErrorPage';

export default InternalErrorPage;
