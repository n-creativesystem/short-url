import { STATUS_CODES, TStatusCodes } from '@/constants/status';
import { ApolloError, ServerError, ServerParseError } from '@apollo/client';
import { GraphQLError, GraphQLErrorExtensions } from 'graphql';

type GraphQLNetworkError = Error | ServerParseError | ServerError;

export const isGraphQLError = (
  error: unknown
): error is ApolloError | GraphQLError[] =>
  isApolloError(error) || error instanceof Array;

export const isApolloError = (error: unknown): error is ApolloError =>
  error instanceof ApolloError;

export const isNetworkError = (
  error: GraphQLNetworkError | null
): error is GraphQLNetworkError => isApolloError(error) && !!error.networkError;

export const isNetworkErrorWithServerParseError = (
  error: GraphQLNetworkError
): error is ServerParseError => {
  const arg = error as ServerParseError;
  return (
    typeof arg.bodyText === 'string' &&
    typeof arg.message === 'string' &&
    typeof arg.name === 'string' &&
    typeof arg.statusCode === 'number' &&
    arg.response instanceof Response
  );
};

export const isNetworkErrorWithServerError = (
  error: GraphQLNetworkError
): error is ServerError => {
  const arg = error as ServerError;
  return (
    (typeof arg.result === 'string' || typeof arg.result === 'object') &&
    typeof arg.message === 'string' &&
    typeof arg.name === 'string' &&
    typeof arg.statusCode === 'number' &&
    arg.response instanceof Response
  );
};

export const isGraphQLExtensions = (
  error: GraphQLErrorExtensions
): error is GraphQLErrorExtensions => {
  return !!error;
};

export const getGraphQLStatusCode = (error: unknown): TStatusCodes => {
  let result: TStatusCodes = STATUS_CODES.UNKNOWN.code;
  if (!isApolloError(error)) {
    return result;
  }
  // networkErrorの場合
  if (isNetworkError(error.networkError)) {
    // ServerErrorにキャストしているが違う場合もあるため注意
    // (Errorの場合は{ message, stack }が渡ってくる)
    const networkError = error.networkError;
    let reason = '';
    if (isNetworkErrorWithServerError(networkError)) {
      if (typeof networkError.result === 'string') {
        reason = networkError.message;
      } else {
        reason = networkError.result?.reason || networkError.result?.message;
      }
    }

    switch (reason) {
      // 未認証の場合
      case STATUS_CODES.UNAUTHORIZED.key:
        result = STATUS_CODES.UNAUTHORIZED.code;
        break;
      // メンテナンスの場合
      case STATUS_CODES.MAINTENANCE.key:
        result = STATUS_CODES.MAINTENANCE.code;
        break;
    }

    return result;
  }

  // GraphQLエラーの場合
  const graphqlError = getGraphQLError(error);
  if (graphqlError && isGraphQLExtensions(graphqlError.extensions)) {
    const extension = convertGraphQLExtensionsToCustomGraphQLErrorExtensions(
      graphqlError.extensions
    );
    switch (extension.code) {
      case STATUS_CODES.INVALID.code:
        result = STATUS_CODES.INVALID.code;
        break;
      case STATUS_CODES.NOT_FOUND.code:
        result = STATUS_CODES.NOT_FOUND.code;
        break;
    }
  }

  return result;
};

export const getGraphQLError = (
  error: ApolloError | GraphQLError[]
): GraphQLError => {
  let result;

  // GraphQLエラーの場合
  if (error instanceof Array) {
    // graphQLErrorsの中身が返ってくる場合
    const _error = error as GraphQLError[];
    result = _error?.[0];
  } else {
    // 配列で渡って来ないときがある (apolloの不具合?)
    result = error?.graphQLErrors?.[0];
  }
  return result;
};

interface CustomGraphQLErrorExtensions {
  code: number;
}

export const convertGraphQLExtensionsToCustomGraphQLErrorExtensions = (
  errorExtensions: GraphQLErrorExtensions
): CustomGraphQLErrorExtensions => {
  const result: CustomGraphQLErrorExtensions = {
    code: 0,
  };
  if (typeof errorExtensions.code === 'number') {
    result.code = errorExtensions.code;
  }
  return result;
};
