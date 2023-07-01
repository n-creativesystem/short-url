import * as Types from '../../../../../types/graphql';

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

export type UpdateOAuthApplicationMutationVariables = Types.Exact<{
  id: Types.Scalars['String'];
  input: Types.OAuthApplicationInput;
}>;

export type UpdateOAuthApplicationMutation = {
  __typename?: 'Mutation';
  updateOAuthApplication: {
    __typename?: 'OAuthApplication';
    id: string;
    name: string;
    secret: string;
    domain: string;
  };
};

export const UpdateOAuthApplicationDocument = gql`
  mutation updateOAuthApplication(
    $id: String!
    $input: OAuthApplicationInput!
  ) {
    updateOAuthApplication(id: $id, input: $input) {
      id
      name
      secret
      domain
    }
  }
`;
export type UpdateOAuthApplicationMutationFn = Apollo.MutationFunction<
  UpdateOAuthApplicationMutation,
  UpdateOAuthApplicationMutationVariables
>;

/**
 * __useUpdateOAuthApplicationMutation__
 *
 * To run a mutation, you first call `useUpdateOAuthApplicationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateOAuthApplicationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateOAuthApplicationMutation, { data, loading, error }] = useUpdateOAuthApplicationMutation({
 *   variables: {
 *      id: // value for 'id'
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateOAuthApplicationMutation(
  baseOptions?: Apollo.MutationHookOptions<
    UpdateOAuthApplicationMutation,
    UpdateOAuthApplicationMutationVariables
  >
) {
  const options = { ...defaultOptions, ...baseOptions };
  return Apollo.useMutation<
    UpdateOAuthApplicationMutation,
    UpdateOAuthApplicationMutationVariables
  >(UpdateOAuthApplicationDocument, options);
}
export type UpdateOAuthApplicationMutationHookResult = ReturnType<
  typeof useUpdateOAuthApplicationMutation
>;
export type UpdateOAuthApplicationMutationResult =
  Apollo.MutationResult<UpdateOAuthApplicationMutation>;
export type UpdateOAuthApplicationMutationOptions = Apollo.BaseMutationOptions<
  UpdateOAuthApplicationMutation,
  UpdateOAuthApplicationMutationVariables
>;
