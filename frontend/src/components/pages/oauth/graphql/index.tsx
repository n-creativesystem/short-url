import * as Types from '../../../../types/graphql';

import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]?: Maybe<T[SubKey]>;
};
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & {
  [SubKey in K]: Maybe<T[SubKey]>;
};
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: CustomTime;
  URL: CustomURL;
};

export type CreateUrlInput = {
  url: Scalars['URL'];
};

export type MetadataType = {
  __typename?: 'MetadataType';
  count: Scalars['Int'];
  next: Scalars['String'];
  prev: Scalars['String'];
  self: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  create: Url;
  createOAuthApplication: OAuthApplication;
  deleteOAuthApplication: Scalars['Boolean'];
  updateOAuthApplication: OAuthApplication;
};

export type MutationCreateArgs = {
  input: CreateUrlInput;
};

export type MutationCreateOAuthApplicationArgs = {
  input: OAuthApplicationInput;
};

export type MutationDeleteOAuthApplicationArgs = {
  id: Scalars['String'];
};

export type MutationUpdateOAuthApplicationArgs = {
  id: Scalars['String'];
  input: OAuthApplicationInput;
};

export type OAuthApplication = {
  __typename?: 'OAuthApplication';
  domain: Scalars['String'];
  id: Scalars['String'];
  name: Scalars['String'];
  secret: Scalars['String'];
};

export type OAuthApplicationInput = {
  name: Scalars['String'];
};

export type OAuthApplicationType = {
  __typename?: 'OAuthApplicationType';
  _metadata: MetadataType;
  result: Array<OAuthApplication>;
};

export type Query = {
  __typename?: 'Query';
  oauthApplication: OAuthApplication;
  oauthApplications: OAuthApplicationType;
  url: Url;
  urls: UrlType;
};

export type QueryOauthApplicationArgs = {
  id: Scalars['String'];
};

export type QueryOauthApplicationsArgs = {
  token?: InputMaybe<Scalars['String']>;
};

export type QueryUrlArgs = {
  key: Scalars['String'];
};

export type Url = {
  __typename?: 'Url';
  created_at: Scalars['Time'];
  key: Scalars['String'];
  updated_at: Scalars['Time'];
  url: Scalars['URL'];
};

export type UrlType = {
  __typename?: 'UrlType';
  result: Array<Url>;
};

export type OAuthApplicationsQueryVariables = Types.Exact<{
  token?: Types.InputMaybe<Types.Scalars['String']>;
}>;

export type OAuthApplicationsQuery = {
  __typename?: 'Query';
  oauthApplications: {
    __typename?: 'OAuthApplicationType';
    result: Array<{
      __typename?: 'OAuthApplication';
      id: string;
      name: string;
      secret: string;
      domain: string;
    }>;
    _metadata: {
      __typename?: 'MetadataType';
      prev: string;
      self: string;
      next: string;
      count: number;
    };
  };
};

export type OAuthApplicationQueryVariables = Types.Exact<{
  id: Types.Scalars['String'];
}>;

export type OAuthApplicationQuery = {
  __typename?: 'Query';
  oauthApplication: {
    __typename?: 'OAuthApplication';
    id: string;
    name: string;
    secret: string;
    domain: string;
  };
};

export const OAuthApplicationsDocument = gql`
  query OAuthApplications($token: String) {
    oauthApplications(token: $token) {
      result {
        id
        name
        secret
        domain
      }
      _metadata {
        prev
        self
        next
        count
      }
    }
  }
`;

/**
 * __useOAuthApplicationsQuery__
 *
 * To run a query within a React component, call `useOAuthApplicationsQuery` and pass it any options that fit your needs.
 * When your component renders, `useOAuthApplicationsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useOAuthApplicationsQuery({
 *   variables: {
 *      token: // value for 'token'
 *   },
 * });
 */
export function useOAuthApplicationsQuery(
  baseOptions?: Apollo.QueryHookOptions<
    OAuthApplicationsQuery,
    OAuthApplicationsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useQuery<
    OAuthApplicationsQuery,
    OAuthApplicationsQueryVariables
  >(OAuthApplicationsDocument, options);
}
export function useOAuthApplicationsLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    OAuthApplicationsQuery,
    OAuthApplicationsQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useLazyQuery<
    OAuthApplicationsQuery,
    OAuthApplicationsQueryVariables
  >(OAuthApplicationsDocument, options);
}
export type OAuthApplicationsQueryHookResult = ReturnType<
  typeof useOAuthApplicationsQuery
>;
export type OAuthApplicationsLazyQueryHookResult = ReturnType<
  typeof useOAuthApplicationsLazyQuery
>;
export type OAuthApplicationsQueryResult = Apollo.QueryResult<
  OAuthApplicationsQuery,
  OAuthApplicationsQueryVariables
>;
export const OAuthApplicationDocument = gql`
  query OAuthApplication($id: String!) {
    oauthApplication(id: $id) {
      id
      name
      secret
      domain
    }
  }
`;

/**
 * __useOAuthApplicationQuery__
 *
 * To run a query within a React component, call `useOAuthApplicationQuery` and pass it any options that fit your needs.
 * When your component renders, `useOAuthApplicationQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useOAuthApplicationQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useOAuthApplicationQuery(
  baseOptions: Apollo.QueryHookOptions<
    OAuthApplicationQuery,
    OAuthApplicationQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useQuery<OAuthApplicationQuery, OAuthApplicationQueryVariables>(
    OAuthApplicationDocument,
    options
  );
}
export function useOAuthApplicationLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<
    OAuthApplicationQuery,
    OAuthApplicationQueryVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useLazyQuery<
    OAuthApplicationQuery,
    OAuthApplicationQueryVariables
  >(OAuthApplicationDocument, options);
}
export type OAuthApplicationQueryHookResult = ReturnType<
  typeof useOAuthApplicationQuery
>;
export type OAuthApplicationLazyQueryHookResult = ReturnType<
  typeof useOAuthApplicationLazyQuery
>;
export type OAuthApplicationQueryResult = Apollo.QueryResult<
  OAuthApplicationQuery,
  OAuthApplicationQueryVariables
>;
