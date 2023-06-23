import { useParams, useRouter } from '@/components/Parts/Navigation';
import { errorModalVar } from '@/components/hooks/Context';
import NotFound from '@/components/pages/notfound';
import { Presenter } from '@/components/pages/oauth/Contents/Presenter';
import type { Input } from '@/components/pages/oauth/Contents/index.d';
import {
  useOAuthApplicationQuery,
  useUpdateOAuthApplicationMutation,
} from '@/graphql/generated';
import { getGraphQLStatusCode } from '@/utils/errors';
import { FC } from 'react';

export const OAuthAppContainer: FC = () => {
  const router = useRouter();
  const { id = '' } = useParams();

  const { data, loading, error } = useOAuthApplicationQuery({
    variables: {
      id: id,
    },
  });

  const [update] = useUpdateOAuthApplicationMutation();
  const onClick = (input: Input): Promise<void> => {
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
  };
  if (loading) {
    return <></>;
  }
  if (error) {
    const code = getGraphQLStatusCode(error);
    if (code === 404) {
      return <NotFound />;
    }
    console.error(error);
    errorModalVar({
      open: true,
      title: 'OAuthApplication',
      description: 'OAuthApplicationの取得時にエラーが発生しました。',
    });
  }
  return <Presenter onClick={onClick} data={data?.oauthApplication} />;
};
