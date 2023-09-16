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
  createOAuthApplication: OAuthApplication;
  deleteOAuthApplication: Scalars['Boolean'];
  deleteURL: Scalars['Boolean'];
  generateURL: Url;
  updateOAuthApplication: OAuthApplication;
  updateURL: Url;
};

export type MutationCreateOAuthApplicationArgs = {
  input: OAuthApplicationInput;
};

export type MutationDeleteOAuthApplicationArgs = {
  id: Scalars['String'];
};

export type MutationDeleteUrlArgs = {
  key: Scalars['String'];
};

export type MutationGenerateUrlArgs = {
  input: CreateUrlInput;
};

export type MutationUpdateOAuthApplicationArgs = {
  id: Scalars['String'];
  input: OAuthApplicationInput;
};

export type MutationUpdateUrlArgs = {
  key: Scalars['String'];
  url: Scalars['URL'];
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

export type UrlsQueryVariables = Types.Exact<{ [key: string]: never }>;

export type UrlsQuery = {
  __typename?: 'Query';
  urls: {
    __typename?: 'UrlType';
    result: Array<{
      __typename?: 'Url';
      key: string;
      url: CustomURL;
      created_at: CustomTime;
      updated_at: CustomTime;
    }>;
  };
};

export type UpdateUrlMutationVariables = Types.Exact<{
  key: Types.Scalars['String'];
  url: Types.Scalars['URL'];
}>;

export type UpdateUrlMutation = {
  __typename?: 'Mutation';
  updateURL: {
    __typename?: 'Url';
    key: string;
    url: CustomURL;
    created_at: CustomTime;
    updated_at: CustomTime;
  };
};

export type DeleteUrlMutationVariables = Types.Exact<{
  key: Types.Scalars['String'];
}>;

export type DeleteUrlMutation = { __typename?: 'Mutation'; deleteURL: boolean };

export const UrlsDocument = gql`
  query urls {
    urls {
      result {
        key
        url
        created_at
        updated_at
      }
    }
  }
`;

/**
 * __useUrlsQuery__
 *
 * To run a query within a React component, call `useUrlsQuery` and pass it any options that fit your needs.
 * When your component renders, `useUrlsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUrlsQuery({
 *   variables: {
 *   },
 * });
 */
export function useUrlsQuery(
  baseOptions?: Apollo.QueryHookOptions<UrlsQuery, UrlsQueryVariables>
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useQuery<UrlsQuery, UrlsQueryVariables>(UrlsDocument, options);
}
export function useUrlsLazyQuery(
  baseOptions?: Apollo.LazyQueryHookOptions<UrlsQuery, UrlsQueryVariables>
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useLazyQuery<UrlsQuery, UrlsQueryVariables>(
    UrlsDocument,
    options
  );
}
export type UrlsQueryHookResult = ReturnType<typeof useUrlsQuery>;
export type UrlsLazyQueryHookResult = ReturnType<typeof useUrlsLazyQuery>;
export type UrlsQueryResult = Apollo.QueryResult<UrlsQuery, UrlsQueryVariables>;
export const UpdateUrlDocument = gql`
  mutation updateUrl($key: String!, $url: URL!) {
    updateURL(key: $key, url: $url) {
      key
      url
      created_at
      updated_at
    }
  }
`;
export type UpdateUrlMutationFn = Apollo.MutationFunction<
  UpdateUrlMutation,
  UpdateUrlMutationVariables
>;

/**
 * __useUpdateUrlMutation__
 *
 * To run a mutation, you first call `useUpdateUrlMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateUrlMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateUrlMutation, { data, loading, error }] = useUpdateUrlMutation({
 *   variables: {
 *      key: // value for 'key'
 *      url: // value for 'url'
 *   },
 * });
 */
export function useUpdateUrlMutation(
  baseOptions?: Apollo.MutationHookOptions<
    UpdateUrlMutation,
    UpdateUrlMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useMutation<UpdateUrlMutation, UpdateUrlMutationVariables>(
    UpdateUrlDocument,
    options
  );
}
export type UpdateUrlMutationHookResult = ReturnType<
  typeof useUpdateUrlMutation
>;
export type UpdateUrlMutationResult = Apollo.MutationResult<UpdateUrlMutation>;
export type UpdateUrlMutationOptions = Apollo.BaseMutationOptions<
  UpdateUrlMutation,
  UpdateUrlMutationVariables
>;
export const DeleteUrlDocument = gql`
  mutation deleteURL($key: String!) {
    deleteURL(key: $key)
  }
`;
export type DeleteUrlMutationFn = Apollo.MutationFunction<
  DeleteUrlMutation,
  DeleteUrlMutationVariables
>;

/**
 * __useDeleteUrlMutation__
 *
 * To run a mutation, you first call `useDeleteUrlMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteUrlMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteUrlMutation, { data, loading, error }] = useDeleteUrlMutation({
 *   variables: {
 *      key: // value for 'key'
 *   },
 * });
 */
export function useDeleteUrlMutation(
  baseOptions?: Apollo.MutationHookOptions<
    DeleteUrlMutation,
    DeleteUrlMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useMutation<DeleteUrlMutation, DeleteUrlMutationVariables>(
    DeleteUrlDocument,
    options
  );
}
export type DeleteUrlMutationHookResult = ReturnType<
  typeof useDeleteUrlMutation
>;
export type DeleteUrlMutationResult = Apollo.MutationResult<DeleteUrlMutation>;
export type DeleteUrlMutationOptions = Apollo.BaseMutationOptions<
  DeleteUrlMutation,
  DeleteUrlMutationVariables
>;
