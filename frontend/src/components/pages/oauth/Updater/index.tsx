import { useParams, useRouter } from '@/components/Parts/Navigation';
import { LoadingContext, errorModalVar } from '@/components/hooks/Context';
import NotFound from '@/components/pages/notfound';
import { Presenter } from '@/components/pages/oauth/Contents/Presenter';
import type { Data, Input } from '@/components/pages/oauth/Contents/index.d';
import { useUpdateOAuthApplicationMutation } from '@/components/pages/oauth/Updater/graphql';
import { useOAuthApplicationQuery } from '@/components/pages/oauth/graphql';
import { getGraphQLStatusCode } from '@/utils/errors';
import { FC, useCallback, useEffect, useState } from 'react';

export const OAuthAppContainer: FC = () => {
  const router = useRouter();
  const { id = '' } = useParams();

  const { data, loading, error } = useOAuthApplicationQuery({
    variables: {
      id: id,
    },
  });

  const [update] = useUpdateOAuthApplicationMutation();
  const onClick = useCallback(
    (input: Input): Promise<void> => {
      return new Promise(async (resolve, reject) => {
        try {
          await update({
            variables: {
              id: id,
              input: {
                name: input.name,
              },
            },
          });
          router.push('/oauth2/app');
          resolve(undefined);
        } catch (error) {
          reject(error);
        }
      });
    },
    [update, router]
  );

  const [isNotFound, setIsNotFound] = useState(false);
  const [initialValues, setInitialValue] = useState<Data>({
    name: '',
    id: '',
    secret: '',
  });

  useEffect(() => {
    if (error) {
      const code = getGraphQLStatusCode(error);
      if (code === 404) {
        setIsNotFound(true);
        return;
      }
      console.error(error);
      errorModalVar({
        open: true,
        title: 'OAuthApplication',
        description: 'OAuthApplicationの取得時にエラーが発生しました。',
      });
    }
    if (!!data && !loading) {
      setInitialValue(data.oauthApplication);
    }
  }, [data, loading, error]);
  return (
    <LoadingContext.Provider value={loading}>
      {isNotFound ? (
        <NotFound />
      ) : (
        <Presenter onClick={onClick} data={initialValues} />
      )}
    </LoadingContext.Provider>
  );
};
