import { errorModalVar } from '@/components/hooks/Context';
import { useCallback } from 'react';
import { UrlsDocument, useDeleteUrlMutation } from './graphql';

export const useDeleteHandler = () => {
  const [update] = useDeleteUrlMutation();
  const handler = useCallback(
    (key: string) => async () => {
      try {
        const { errors } = await update({
          variables: {
            key: key,
          },
          refetchQueries: [{ query: UrlsDocument }],
        });
        if (errors) throw errors;
      } catch (error) {
        console.error(error);
        errorModalVar({
          open: true,
          title: 'Shorts',
          description: '削除時にエラーが発生しました。',
        });
      }
    },
    [update]
  );
  return (key: string) => {
    return handler(key);
  };
};
