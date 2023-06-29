/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
  'mutation createOAuthApplication($input: OAuthApplicationInput!) {\n  createOAuthApplication(input: $input) {\n    id\n    name\n    secret\n    domain\n  }\n}':
    types.CreateOAuthApplicationDocument,
  'mutation deleteOAuthApplication($id: String!) {\n  deleteOAuthApplication(id: $id)\n}':
    types.DeleteOAuthApplicationDocument,
  'mutation updateOAuthApplication($id: String!, $input: OAuthApplicationInput!) {\n  updateOAuthApplication(id: $id, input: $input) {\n    id\n    name\n    secret\n    domain\n  }\n}':
    types.UpdateOAuthApplicationDocument,
  'query OAuthApplications($token: String) {\n  oauthApplications(token: $token) {\n    result {\n      id\n      name\n      secret\n      domain\n    }\n    _metadata {\n      prev\n      self\n      next\n      count\n    }\n  }\n}\n\nquery OAuthApplication($id: String!) {\n  oauthApplication(id: $id) {\n    id\n    name\n    secret\n    domain\n  }\n}':
    types.OAuthApplicationsDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: 'mutation createOAuthApplication($input: OAuthApplicationInput!) {\n  createOAuthApplication(input: $input) {\n    id\n    name\n    secret\n    domain\n  }\n}'
): (typeof documents)['mutation createOAuthApplication($input: OAuthApplicationInput!) {\n  createOAuthApplication(input: $input) {\n    id\n    name\n    secret\n    domain\n  }\n}'];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: 'mutation deleteOAuthApplication($id: String!) {\n  deleteOAuthApplication(id: $id)\n}'
): (typeof documents)['mutation deleteOAuthApplication($id: String!) {\n  deleteOAuthApplication(id: $id)\n}'];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: 'mutation updateOAuthApplication($id: String!, $input: OAuthApplicationInput!) {\n  updateOAuthApplication(id: $id, input: $input) {\n    id\n    name\n    secret\n    domain\n  }\n}'
): (typeof documents)['mutation updateOAuthApplication($id: String!, $input: OAuthApplicationInput!) {\n  updateOAuthApplication(id: $id, input: $input) {\n    id\n    name\n    secret\n    domain\n  }\n}'];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(
  source: 'query OAuthApplications($token: String) {\n  oauthApplications(token: $token) {\n    result {\n      id\n      name\n      secret\n      domain\n    }\n    _metadata {\n      prev\n      self\n      next\n      count\n    }\n  }\n}\n\nquery OAuthApplication($id: String!) {\n  oauthApplication(id: $id) {\n    id\n    name\n    secret\n    domain\n  }\n}'
): (typeof documents)['query OAuthApplications($token: String) {\n  oauthApplications(token: $token) {\n    result {\n      id\n      name\n      secret\n      domain\n    }\n    _metadata {\n      prev\n      self\n      next\n      count\n    }\n  }\n}\n\nquery OAuthApplication($id: String!) {\n  oauthApplication(id: $id) {\n    id\n    name\n    secret\n    domain\n  }\n}'];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> =
  TDocumentNode extends DocumentNode<infer TType, any> ? TType : never;
