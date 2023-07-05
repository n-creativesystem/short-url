import { internalErrorPageVar } from '@/components/hooks/Context';
import { useEffect, useState } from 'react';
import { openApiClient } from './useFetch';

export const useCSRFToken = (): string => {
  const [token, setToken] = useState('');
  useEffect(() => {
    let ignore = false;
    const fetch = async () => {
      try {
        const data = await openApiClient.csrf_token.$get();
        if (!ignore && data) {
          setToken(data.csrf_token ?? '');
        }
      } catch (error) {
        internalErrorPageVar({
          show: true,
        });
        return;
      }
    };
    fetch();
    return () => {
      ignore = true;
    };
  }, []);
  return token;
};
