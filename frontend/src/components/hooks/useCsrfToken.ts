import { internalErrorPageVar } from '@/components/hooks/Context';
import { useFetchByOpenAPI } from '@/components/hooks/useFetchOpenAPI';
import { useEffect, useState } from 'react';

export const useCSRFToken = (): string => {
  const [token, setToken] = useState('');
  const { data, isLoading, error, hasError } = useFetchByOpenAPI({
    url: '/csrf_token',
    method: 'get',
  });
  useEffect(() => {
    let ignore = false;
    if (isLoading) {
      return;
    }
    if (hasError) {
      internalErrorPageVar({
        show: true,
      });
      return;
    }
    if (!ignore && data) {
      setToken(data.csrf_token ?? '');
    }
    return () => {
      ignore = true;
    };
  }, [data, isLoading, hasError, error]);
  return token;
};
