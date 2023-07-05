/* eslint-disable @typescript-eslint/no-explicit-any*/

import openapi from '@/openapi/$api';
import aspida from '@aspida/axios';
import axios, { AxiosError, AxiosResponse, Method } from 'axios';
import { useCallback, useEffect, useMemo, useState } from 'react';

type Header = { [key: string]: string };
type Params = { [key: string]: any };

type FetchRequest<T> = {
  method?: 'get' | 'post';
  url: string;
  params?: Params;
  headers?: Header;
  onSuccess?: (data?: T) => void;
  onError?: (err: AxiosError) => void;
};

type Refetch<T> = ({ url, params }?: RefetchArgs<T>) => Promise<T | null>;

type RefetchArgs<T> = {
  url?: string;
  method?: Pick<Method, 'get' & 'post'>;
  params?: Params;
  body?: Params;
  onSuccess?: (data?: T) => void;
  onError?: (err: AxiosError) => void;
};

const baseUrl = '/api';
export const axiosInstance = axios.create({
  baseURL: baseUrl,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
    withCredentials: true,
  },
  withCredentials: true,
});

export const openApiClient = openapi(aspida(axiosInstance));

type FetchResponse<T> = {
  data?: T | null;
  refetch: Refetch<T>;
  error: any;
  hasError: boolean;
  isLoading: boolean;
};

export const useFetch: <T>(req: FetchRequest<T>) => FetchResponse<T> = <T>({
  url,
  method = 'get',
  params,
  headers,
  onError,
}: FetchRequest<T>): FetchResponse<T> => {
  const [data, setData] = useState<T | null>(null);
  const [isLoading, setLoading] = useState(false);
  const [error, setError] = useState<AxiosError | null>(null);
  const [hasError, setHasError] = useState(false);
  const memoUrl = useMemo(() => url, [url]);
  const refetch = useCallback<Refetch<T>>(
    async <T>(args?: RefetchArgs<T>) => {
      try {
        setLoading(true);
        const res = await fetch<T>({ method, url: memoUrl, params, headers });
        const data = res.data as any;
        setData(data);
        args?.onSuccess?.(data);
        return data;
      } catch (error) {
        const err = error as AxiosError;
        onError?.(err);
        setError(err);
        setHasError(true);
        return null;
      } finally {
        setLoading(false);
      }
    },
    [headers, memoUrl, method, onError, params]
  );
  const clear = useCallback(() => {
    setData(null);
    setLoading(false);
    setHasError(false);
    setError(null);
  }, []);

  useEffect(() => {
    if (memoUrl) {
      (async () => {
        const res = await refetch({});
        res && setData(res);
      })();
    }
    return () => clear();
  }, [memoUrl, clear, refetch]);
  return {
    data,
    refetch,
    error,
    hasError,
    isLoading,
  };
};

export const fetch = async <T>({
  method = 'get',
  url,
  params,
  headers,
}: FetchRequest<T>): Promise<AxiosResponse<T, any>> => {
  const instance = axiosInstance[method];
  const res = await instance<T>(`${url}`, params, {
    headers: {
      ...headers,
    },
  });
  return res;
};
