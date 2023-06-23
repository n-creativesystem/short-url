import { errorModalVar } from '@/components/hooks/Context';
import { useFetch } from '@/components/hooks/useFetch';
import { paths } from '@/openapi/schema';
import { GetOAuthButton } from './index.d';

export const useAuthLogin = (label: string) => async () => {
  window.location.href = `/api/auth/${label}/authorize`;
};

export const useEnabledAuth: () => GetOAuthButton = (): GetOAuthButton => {
  const { data, isLoading, error, hasError } = useFetch<
    paths['/auth/enabled']['get']['responses']['200']['schema']
  >({
    url: '/auth/enabled',
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
