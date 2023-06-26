import { errorModalVar } from '@/components/hooks/Context';
import { useFetchByOpenAPI } from '@/components/hooks/useFetchOpenAPI';
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
  if (isLoading) {
    return {
      buttons: [],
      isLoading: true,
    };
  }
  if (hasError) {
    console.error(error);
    errorModalVar({
      open: true,
      title: 'ログイン',
      description: 'ログインボタン取得に失敗しました。',
    });
  }
  return {
    buttons: data?.socials || [],
    isLoading: false,
  };
};
