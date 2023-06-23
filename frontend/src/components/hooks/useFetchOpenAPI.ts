/* eslint-disable @typescript-eslint/no-explicit-any*/

import {
  ApiError,
  ErrorCodes,
  Methods,
  Paths,
  RequestData,
  RequestParameters,
  ResponseData,
  SuccessResponseData,
} from '@/openapi/schema.helper.d';
import axios, { AxiosError, AxiosRequestConfig, AxiosResponse } from 'axios';
import { useCallback, useEffect, useMemo, useState } from 'react';
import { axiosInstance } from './useFetch';

type Header = { [key: string]: string };

type RefetchByOpenAPI<Path extends Paths, Method extends Methods> = ({
  url,
  params,
}?: RefetchArgsByOpenAPI<Path, Method>) => Promise<SuccessResponseData<
  Path,
  Method
> | null>;

type RefetchArgsByOpenAPI<Path extends Paths, Method extends Methods> = {
  url: string;
  method: Method;
  params?: {
    query?: RequestParameters<Path, Method>;
    body?: RequestData<Path, Method>;
  };
  headers?: Header;
  onSuccess?: (data?: SuccessResponseData<Path, Method>) => void;
  onError?: (err: AxiosError) => void;
};

type FetchResponseByOpenAPI<
  Path extends Paths,
  Method extends Methods,
  Code extends number
> = {
  data?: ResponseData<Path, Method, Code> | null;
  refetch: RefetchByOpenAPI<Path, Method>;
  error: any;
  hasError: boolean;
  isLoading: boolean;
};

type FetchRequestByOpenAPI<Path extends Paths, Method extends Methods> = {
  url: Path;
  method: Method;
  params?: {
    query?: RequestParameters<Path, Method>;
    body?: RequestData<Path, Method>;
  };
  headers?: Header;
  onError?: ((err: AxiosError) => void) | undefined;
};

export const useFetchByOpenAPI = <P extends Paths, M extends Methods>({
  url,
  method,
  params,
  headers,
  onError,
}: FetchRequestByOpenAPI<P, M>): FetchResponseByOpenAPI<P, M, 200> => {
  const [data, setData] = useState<ResponseData<P, M, 200> | null>(null);
  const [isLoading, setLoading] = useState(false);
  const [error, setError] = useState<ApiError<P, M> | null>(null);
  const [hasError, setHasError] = useState(false);
  const memoUrl = useMemo(() => url, [url]);
  const refetch = useCallback<RefetchByOpenAPI<P, M>>(
    async (args?: RefetchArgsByOpenAPI<P, M>) => {
      try {
        setLoading(true);
        const res = await fetchByOpenAPI<P, M>(
          args ?? {
            url: memoUrl,
            method,
            params,
            headers,
          }
        );
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
        const res = await refetch();
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

export const fetchByOpenAPI = async <P extends Paths, M extends Methods>({
  url,
  method,
  params,
  headers,
}: RefetchArgsByOpenAPI<P, M>): Promise<
  AxiosResponse<SuccessResponseData<P, M>, any>
> => {
  const instance = axiosInstance[method];
  const config: AxiosRequestConfig = {
    params: params?.query,
    data: params?.body,
    headers: {
      ...headers,
    },
  };
  try {
    return await instance<SuccessResponseData<P, M>>(`${url}`, config);
  } catch (error) {
    if (axios.isAxiosError(error) && !!error.response) {
      const errorData = {
        status: error.response.status,
        data: error.response.data,
      };
      if (isExpectedError<P, M>(errorData)) {
        throw errorData;
      }
    }
    throw error;
  }
};

const isExpectedError = <Path extends Paths, Method extends Methods>(res: {
  status: number;
  data: any;
}): res is ApiError<Path, Method> => {
  return ErrorCodes.map(Number).includes(res.status);
};
