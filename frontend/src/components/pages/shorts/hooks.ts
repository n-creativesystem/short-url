import { errorModalVar } from '@/components/hooks/Context';
import { useCallback } from 'react';
import { UrlsDocument, useDeleteUrlMutation } from './graphql';

export const useDeleteHandler = () => (key: string) => {
  const [update] = useDeleteUrlMutation();
  const handler = useCallback(async () => {
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
  }, [key]);
  return handler;
};
