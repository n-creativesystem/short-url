import { errorModalVar } from '@/components/hooks/Context';
import { useFetchByOpenAPI } from '@/components/hooks/useFetchOpenAPI';
import { useEffect, useState } from 'react';
import { GetOAuthButton } from './index.d';

export const useAuthLogin = (label: string) => async () => {
  window.location.href = `/api/auth/${label}/authorize`;
};

export const useEnabledAuth: () => GetOAuthButton = (): GetOAuthButton => {
  const { data, isLoading, error, hasError } = useFetchByOpenAPI<
    '/auth/enabled',
    'get'
  >({
    url: '/auth/enabled',
    method: 'get',
  });
  const [state, setState] = useState<GetOAuthButton>({
    buttons: [],
    isLoading: isLoading,
  });

  useEffect(() => {
    if (hasError) {
      console.error(error);
      errorModalVar({
        open: true,
        title: 'ログイン',
        description: 'ログインボタン取得に失敗しました。',
      });
      setState((prev) => ({ ...prev, isLoading: false }));
    } else {
      setState({
        buttons: data?.socials || [],
        isLoading: isLoading,
      });
    }
  }, [data, isLoading, error, hasError]);
  return state;
};
