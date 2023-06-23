/* eslint-disable @typescript-eslint/no-explicit-any*/

import { useCsrfToken } from '@/components/Parts/Layout/CsrfToken';
import { axiosInstance } from '@/components/hooks/useFetch';
import {
  ApolloClient,
  ApolloLink,
  HttpLink,
  InMemoryCache,
  NormalizedCacheObject,
  from,
} from '@apollo/client';
import { AxiosError, AxiosHeaders, AxiosResponse } from 'axios';
import { createAxiosHeaders, createFetchHeaders } from './http/header';

const instance = axiosInstance['post'];

type TClient = ApolloClient<NormalizedCacheObject>;

export const useApolloClient = () => {
  const token = useCsrfToken();
  return createClient(token);
};

const getTypePolicies = () => ({});
const getCache = () =>
  new InMemoryCache({
    typePolicies: getTypePolicies(),
  });

const createClient = (token: string): TClient => {
  const cache = getCache();
  const middleware = new ApolloLink((operation, forward) => {
    operation.setContext(({ headers = {} }) => {
      return {
        headers: {
          ...headers,
          'X-Csrf-Token': token,
        },
      };
    });

    return forward(operation);
  });

  const httpLink = new HttpLink({
    uri: '/graphql',
    credentials: 'include',
    fetch: async (input: RequestInfo | URL, init: RequestInit | undefined) => {
      const rawHeaders = createAxiosHeaders(init?.headers);
      const headers = new AxiosHeaders(rawHeaders);
      let result: AxiosResponse<any, any> = {
        data: undefined,
        status: 200,
        statusText: '',
        headers: {},
        config: {
          headers: headers,
        },
      };
      try {
        result = await instance(getUrl(input), init?.body, {
          responseType: 'arraybuffer',
          headers: headers,
        });
      } catch (err: any) {
        if (err instanceof AxiosError) {
          const axiosErr = err as AxiosError;
          if (axiosErr.response) {
            result = axiosErr.response;
          } else {
            throw err;
          }
        }
      }
      return new Response(result?.data, {
        status: result.status,
        statusText: result.statusText,
        headers: createFetchHeaders(result.headers),
      });
    },
  });

  return new ApolloClient({
    ssrMode: false,
    cache,
    connectToDevTools: process.env.NODE_ENV !== 'production',
    link: from([middleware, httpLink]),
    defaultOptions: {
      watchQuery: {
        fetchPolicy: 'cache-and-network',
      },
    },
  });
};

const getUrl = (input?: string | { href?: string; url?: string }): string => {
  if (typeof input === 'string') {
    return input;
  } else if (input?.href) {
    return input.href;
  } else if (input?.url) {
    return input.url;
  }
  return '';
};
