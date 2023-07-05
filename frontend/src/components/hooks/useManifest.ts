import { internalErrorPageVar } from '@/components/hooks/Context';
import { Response_WebUIManifest } from '@openapi/@types';
import { useEffect, useState } from 'react';
import { openApiClient } from './useFetch';

export const useManifest = () => {
  const initialize: Required<Response_WebUIManifest> = {
    header_name: '',
    token_base: false,
  };
  const [manifest, setManifest] = useState(initialize);
  useEffect(() => {
    let ignore = false;
    const fetch = async () => {
      try {
        const data = await openApiClient.manifest.$get();
        if (!ignore && data) {
          setManifest({
            header_name: data.header_name ?? '',
            token_base: data.token_base ?? false,
          });
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
  return manifest;
};
