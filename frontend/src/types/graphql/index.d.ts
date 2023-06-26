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
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type MetadataType = {
  __typename?: 'MetadataType';
  count: Scalars['Int'];
  next: Scalars['String'];
  prev: Scalars['String'];
  self: Scalars['String'];
};

export type OAuth2ClientMutation = {
  __typename?: 'OAuth2ClientMutation';
  createOAuthApplication: OAuthApplication;
  deleteOAuthApplication: Scalars['Boolean'];
  updateOAuthApplication: OAuthApplication;
};

export type OAuth2ClientMutationCreateOAuthApplicationArgs = {
  input: OAuthApplicationInput;
};

export type OAuth2ClientMutationDeleteOAuthApplicationArgs = {
  id: Scalars['String'];
};

export type OAuth2ClientMutationUpdateOAuthApplicationArgs = {
  id: Scalars['String'];
  input: OAuthApplicationInput;
};

export type OAuth2ClientQuery = {
  __typename?: 'OAuth2ClientQuery';
  oauthApplication: OAuthApplication;
  oauthApplications: OAuthApplicationType;
};

export type OAuth2ClientQueryOauthApplicationArgs = {
  id: Scalars['String'];
};

export type OAuth2ClientQueryOauthApplicationsArgs = {
  token?: InputMaybe<Scalars['String']>;
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
