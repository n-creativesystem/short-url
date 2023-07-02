import { useOutletContext } from '@/components/hooks/useOutlet';
import { Container } from '@/components/pages/shorts';
import { FC, memo } from 'react';

const Index: FC = memo(() => {
  const { setTitle } = useOutletContext();
  setTitle('短縮URL');
  return <Container />;
});

Index.displayName = 'ShortsPage';

export default Index;
