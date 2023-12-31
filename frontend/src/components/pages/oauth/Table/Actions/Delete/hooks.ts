import { errorModalVar } from '@/components/hooks/Context';
import { useDeleteOAuthApplicationMutation } from '@/components/pages/oauth/Table/Actions/Delete/graphql';
import { OAuthApplicationsDocument } from '@/components/pages/oauth/graphql';
import { useCallback } from 'react';

export const useDeleteHandler = () => (id: string) => {
  const [update] = useDeleteOAuthApplicationMutation();
  const handler = useCallback(async () => {
    try {
      const { errors } = await update({
        variables: {
          id: id,
        },
        refetchQueries: [
          { query: OAuthApplicationsDocument, variables: { token: '' } },
        ],
      });
      if (errors) throw errors;
    } catch (error) {
      console.error(error);
      errorModalVar({
        open: true,
        title: 'OAuthApplication',
        description: '削除時にエラーが発生しました。',
      });
    }
  }, [id]);
  return handler;
};
