import { useRouter } from '@/components/Parts/Navigation';
import { LoadingContext } from '@/components/hooks/Context';
import { FC, memo, useCallback } from 'react';
import { Presenter } from './Presenter';
import { useUrlsQuery } from './graphql';
import { useDeleteHandler } from './hooks';

export type Props = {};

export const Container: FC<Props> = memo(() => {
  const router = useRouter();
  const { data, loading } = useUrlsQuery({});
  const onRegisterClick = useCallback(() => {
    router.push('/shorts/register');
  }, [router]);
  const handler = useDeleteHandler();
  return (
    <LoadingContext.Provider value={loading}>
      <Presenter
        data={data?.urls?.result ?? []}
        registerHandler={onRegisterClick}
        deleteHandler={handler}
      />
    </LoadingContext.Provider>
  );
});
Container.displayName = 'Container';
