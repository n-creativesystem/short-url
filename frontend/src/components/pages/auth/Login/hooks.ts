import { errorModalVar } from '@/components/hooks/Context';
import { openApiClient } from '@/components/hooks/useFetch';
import { useEffect, useState } from 'react';
import { GetOAuthButton } from './index.d';

export const useAuthLogin = (label: string) => async () => {
  window.location.href = `/api/auth/${label}/authorize`;
};

export const useEnabledAuth: () => GetOAuthButton = (): GetOAuthButton => {
  const [state, setState] = useState<GetOAuthButton>({
    buttons: [],
    isLoading: true,
  });

  useEffect(() => {
    const innerValue: GetOAuthButton = {
      buttons: [],
      isLoading: true,
    };
    const fetch = async () => {
      try {
        const data = await openApiClient.auth.enabled.$get();
        innerValue.buttons = data?.socials ?? [];
      } catch (error) {
        console.error(error);
        errorModalVar({
          open: true,
          title: 'ログイン',
          description: 'ログインボタン取得に失敗しました。',
        });
      } finally {
        innerValue.isLoading = false;
      }
      setState(innerValue);
    };
    fetch();
  }, []);
  return state;
};
