/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
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
export type MakeEmpty<
  T extends { [key: string]: unknown },
  K extends keyof T
> = { [_ in K]?: never };
export type Incremental<T> =
  | T
  | {
      [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never;
    };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string };
  String: { input: string; output: string };
  Boolean: { input: boolean; output: boolean };
  Int: { input: number; output: number };
  Float: { input: number; output: number };
  Time: { input: any; output: any };
  URL: { input: any; output: any };
};

export type CreateUrlInput = {
  url: Scalars['URL']['input'];
};

export type MetadataType = {
  __typename?: 'MetadataType';
  count: Scalars['Int']['output'];
  next: Scalars['String']['output'];
  prev: Scalars['String']['output'];
  self: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createOAuthApplication: OAuthApplication;
  deleteOAuthApplication: Scalars['Boolean']['output'];
  deleteURL: Scalars['Boolean']['output'];
  generateURL: Url;
  updateOAuthApplication: OAuthApplication;
  updateURL: Url;
};

export type MutationCreateOAuthApplicationArgs = {
  input: OAuthApplicationInput;
};

export type MutationDeleteOAuthApplicationArgs = {
  id: Scalars['String']['input'];
};

export type MutationDeleteUrlArgs = {
  key: Scalars['String']['input'];
};

export type MutationGenerateUrlArgs = {
  input: CreateUrlInput;
};

export type MutationUpdateOAuthApplicationArgs = {
  id: Scalars['String']['input'];
  input: OAuthApplicationInput;
};

export type MutationUpdateUrlArgs = {
  key: Scalars['String']['input'];
  url: Scalars['URL']['input'];
};

export type OAuthApplication = {
  __typename?: 'OAuthApplication';
  domain: Scalars['String']['output'];
  id: Scalars['String']['output'];
  name: Scalars['String']['output'];
  secret: Scalars['String']['output'];
};

export type OAuthApplicationInput = {
  name: Scalars['String']['input'];
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
  id: Scalars['String']['input'];
};

export type QueryOauthApplicationsArgs = {
  token?: InputMaybe<Scalars['String']['input']>;
};

export type QueryUrlArgs = {
  key: Scalars['String']['input'];
};

export type Url = {
  __typename?: 'Url';
  created_at: Scalars['Time']['output'];
  key: Scalars['String']['output'];
  updated_at: Scalars['Time']['output'];
  url: Scalars['URL']['output'];
};

export type UrlType = {
  __typename?: 'UrlType';
  result: Array<Url>;
};

export type CreateOAuthApplicationMutationVariables = Exact<{
  input: OAuthApplicationInput;
}>;

export type CreateOAuthApplicationMutation = {
  __typename?: 'Mutation';
  createOAuthApplication: {
    __typename?: 'OAuthApplication';
    id: string;
    name: string;
    secret: string;
    domain: string;
  };
};

export type DeleteOAuthApplicationMutationVariables = Exact<{
  id: Scalars['String']['input'];
}>;

export type DeleteOAuthApplicationMutation = {
  __typename?: 'Mutation';
  deleteOAuthApplication: boolean;
};

export type UpdateOAuthApplicationMutationVariables = Exact<{
  id: Scalars['String']['input'];
  input: OAuthApplicationInput;
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

export type OAuthApplicationsQueryVariables = Exact<{
  token?: InputMaybe<Scalars['String']['input']>;
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

export type OAuthApplicationQueryVariables = Exact<{
  id: Scalars['String']['input'];
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

export type ResultFragment = {
  __typename?: 'Url';
  key: string;
  url: any;
  created_at: any;
  updated_at: any;
} & { ' $fragmentName'?: 'ResultFragment' };

export type UrlsQueryVariables = Exact<{ [key: string]: never }>;

export type UrlsQuery = {
  __typename?: 'Query';
  urls: {
    __typename?: 'UrlType';
    result: Array<
      { __typename?: 'Url' } & {
        ' $fragmentRefs'?: { ResultFragment: ResultFragment };
      }
    >;
  };
};

export type UpdateUrlMutationVariables = Exact<{
  key: Scalars['String']['input'];
  url: Scalars['URL']['input'];
}>;

export type UpdateUrlMutation = {
  __typename?: 'Mutation';
  updateURL: {
    __typename?: 'Url';
    key: string;
    url: any;
    created_at: any;
    updated_at: any;
  };
};

export type DeleteUrlMutationVariables = Exact<{
  key: Scalars['String']['input'];
}>;

export type DeleteUrlMutation = { __typename?: 'Mutation'; deleteURL: boolean };

export const ResultFragmentDoc = {
  kind: 'Document',
  definitions: [
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'result' },
      typeCondition: {
        kind: 'NamedType',
        name: { kind: 'Name', value: 'Url' },
      },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'key' } },
          { kind: 'Field', name: { kind: 'Name', value: 'url' } },
          { kind: 'Field', name: { kind: 'Name', value: 'created_at' } },
          { kind: 'Field', name: { kind: 'Name', value: 'updated_at' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<ResultFragment, unknown>;
export const CreateOAuthApplicationDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'createOAuthApplication' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: {
            kind: 'Variable',
            name: { kind: 'Name', value: 'input' },
          },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'OAuthApplicationInput' },
            },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'createOAuthApplication' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'input' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'input' },
                },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'name' } },
                { kind: 'Field', name: { kind: 'Name', value: 'secret' } },
                { kind: 'Field', name: { kind: 'Name', value: 'domain' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  CreateOAuthApplicationMutation,
  CreateOAuthApplicationMutationVariables
>;
export const DeleteOAuthApplicationDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'deleteOAuthApplication' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'String' },
            },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'deleteOAuthApplication' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'id' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'id' },
                },
              },
            ],
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  DeleteOAuthApplicationMutation,
  DeleteOAuthApplicationMutationVariables
>;
export const UpdateOAuthApplicationDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'updateOAuthApplication' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'String' },
            },
          },
        },
        {
          kind: 'VariableDefinition',
          variable: {
            kind: 'Variable',
            name: { kind: 'Name', value: 'input' },
          },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'OAuthApplicationInput' },
            },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'updateOAuthApplication' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'id' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'id' },
                },
              },
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'input' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'input' },
                },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'name' } },
                { kind: 'Field', name: { kind: 'Name', value: 'secret' } },
                { kind: 'Field', name: { kind: 'Name', value: 'domain' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  UpdateOAuthApplicationMutation,
  UpdateOAuthApplicationMutationVariables
>;
export const OAuthApplicationsDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'OAuthApplications' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: {
            kind: 'Variable',
            name: { kind: 'Name', value: 'token' },
          },
          type: { kind: 'NamedType', name: { kind: 'Name', value: 'String' } },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'oauthApplications' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'token' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'token' },
                },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                {
                  kind: 'Field',
                  name: { kind: 'Name', value: 'result' },
                  selectionSet: {
                    kind: 'SelectionSet',
                    selections: [
                      { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                      { kind: 'Field', name: { kind: 'Name', value: 'name' } },
                      {
                        kind: 'Field',
                        name: { kind: 'Name', value: 'secret' },
                      },
                      {
                        kind: 'Field',
                        name: { kind: 'Name', value: 'domain' },
                      },
                    ],
                  },
                },
                {
                  kind: 'Field',
                  name: { kind: 'Name', value: '_metadata' },
                  selectionSet: {
                    kind: 'SelectionSet',
                    selections: [
                      { kind: 'Field', name: { kind: 'Name', value: 'prev' } },
                      { kind: 'Field', name: { kind: 'Name', value: 'self' } },
                      { kind: 'Field', name: { kind: 'Name', value: 'next' } },
                      { kind: 'Field', name: { kind: 'Name', value: 'count' } },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  OAuthApplicationsQuery,
  OAuthApplicationsQueryVariables
>;
export const OAuthApplicationDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'OAuthApplication' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'id' } },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'String' },
            },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'oauthApplication' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'id' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'id' },
                },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'id' } },
                { kind: 'Field', name: { kind: 'Name', value: 'name' } },
                { kind: 'Field', name: { kind: 'Name', value: 'secret' } },
                { kind: 'Field', name: { kind: 'Name', value: 'domain' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<
  OAuthApplicationQuery,
  OAuthApplicationQueryVariables
>;
export const UrlsDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'query',
      name: { kind: 'Name', value: 'urls' },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'urls' },
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                {
                  kind: 'Field',
                  name: { kind: 'Name', value: 'result' },
                  selectionSet: {
                    kind: 'SelectionSet',
                    selections: [
                      {
                        kind: 'FragmentSpread',
                        name: { kind: 'Name', value: 'result' },
                      },
                    ],
                  },
                },
              ],
            },
          },
        ],
      },
    },
    {
      kind: 'FragmentDefinition',
      name: { kind: 'Name', value: 'result' },
      typeCondition: {
        kind: 'NamedType',
        name: { kind: 'Name', value: 'Url' },
      },
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          { kind: 'Field', name: { kind: 'Name', value: 'key' } },
          { kind: 'Field', name: { kind: 'Name', value: 'url' } },
          { kind: 'Field', name: { kind: 'Name', value: 'created_at' } },
          { kind: 'Field', name: { kind: 'Name', value: 'updated_at' } },
        ],
      },
    },
  ],
} as unknown as DocumentNode<UrlsQuery, UrlsQueryVariables>;
export const UpdateUrlDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'updateUrl' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'key' } },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'String' },
            },
          },
        },
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'url' } },
          type: {
            kind: 'NonNullType',
            type: { kind: 'NamedType', name: { kind: 'Name', value: 'URL' } },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'updateURL' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'key' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'key' },
                },
              },
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'url' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'url' },
                },
              },
            ],
            selectionSet: {
              kind: 'SelectionSet',
              selections: [
                { kind: 'Field', name: { kind: 'Name', value: 'key' } },
                { kind: 'Field', name: { kind: 'Name', value: 'url' } },
                { kind: 'Field', name: { kind: 'Name', value: 'created_at' } },
                { kind: 'Field', name: { kind: 'Name', value: 'updated_at' } },
              ],
            },
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<UpdateUrlMutation, UpdateUrlMutationVariables>;
export const DeleteUrlDocument = {
  kind: 'Document',
  definitions: [
    {
      kind: 'OperationDefinition',
      operation: 'mutation',
      name: { kind: 'Name', value: 'deleteURL' },
      variableDefinitions: [
        {
          kind: 'VariableDefinition',
          variable: { kind: 'Variable', name: { kind: 'Name', value: 'key' } },
          type: {
            kind: 'NonNullType',
            type: {
              kind: 'NamedType',
              name: { kind: 'Name', value: 'String' },
            },
          },
        },
      ],
      selectionSet: {
        kind: 'SelectionSet',
        selections: [
          {
            kind: 'Field',
            name: { kind: 'Name', value: 'deleteURL' },
            arguments: [
              {
                kind: 'Argument',
                name: { kind: 'Name', value: 'key' },
                value: {
                  kind: 'Variable',
                  name: { kind: 'Name', value: 'key' },
                },
              },
            ],
          },
        ],
      },
    },
  ],
} as unknown as DocumentNode<DeleteUrlMutation, DeleteUrlMutationVariables>;
