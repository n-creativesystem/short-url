import { useRouter } from '@/components/Parts/Navigation';
import { Presenter } from '@/components/pages/oauth/Contents/Presenter';
import type { Input } from '@/components/pages/oauth/Contents/index.d';
import { useCreateOAuthApplicationMutation } from '@/components/pages/oauth/Register/graphql';
import { FC, useCallback } from 'react';

export const OAuthAppContainer: FC = () => {
  const router = useRouter();
  const [update] = useCreateOAuthApplicationMutation();
  const onClick = useCallback(
    (input: Input): Promise<void> => {
      return new Promise(async (resolve, reject) => {
        try {
          await update({
            variables: {
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
    [router]
  );
  return <Presenter onClick={onClick} />;
};
