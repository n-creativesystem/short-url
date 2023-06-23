import { internalErrorPageVar } from '@/components/hooks/Context';
import { useFetchByOpenAPI } from '@/components/hooks/useFetchOpenAPI';
import { SuccessResponseData } from '@/openapi/schema.helper';
import { useEffect, useState } from 'react';

export const useManifest = (): SuccessResponseData<'/manifest', 'get'> => {
  const initialize: Required<SuccessResponseData<'/manifest', 'get'>> = {
    header_name: '',
    token_base: false,
  };
  const [manifest, setManifest] = useState(initialize);
  const { data, isLoading, error, hasError } = useFetchByOpenAPI({
    url: '/manifest',
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
      setManifest({
        header_name: data.header_name ?? '',
        token_base: data.token_base ?? false,
      });
    }
    return () => {
      ignore = true;
    };
  }, [data, isLoading, hasError, error]);
  return manifest;
};
