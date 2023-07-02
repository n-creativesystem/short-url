import { useRouter } from '@/components/Parts/Navigation';
import { LoadingContext } from '@/components/hooks/Context';
import { useOAuthApplicationsQuery } from '@/components/pages/oauth/graphql';
import { FC, memo, useCallback } from 'react';
import Presenter from './Presenter';
import { useDeleteHandler } from './Table/Actions';
type Props = {};

const OAuth2ApplicationsContainer: FC<Props> = memo(({}) => {
  const router = useRouter();
  const { data, loading } = useOAuthApplicationsQuery({
    variables: {
      token: '',
    },
  });

  const onRegisterClick = useCallback(() => {
    router.push('/oauth2/app/register');
  }, [router]);
  const deleteHandler = useDeleteHandler();

  return (
    <LoadingContext.Provider value={loading}>
      <Presenter
        data={data?.oauthApplications?.result || []}
        registerHandler={onRegisterClick}
        deleteHandler={deleteHandler}
      />
    </LoadingContext.Provider>
  );
});

OAuth2ApplicationsContainer.displayName = 'OAuth2ApplicationContainer';

export default OAuth2ApplicationsContainer;
