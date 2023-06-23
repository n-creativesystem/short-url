import { Get, UnionToIntersection } from 'type-fest';
import { paths } from './schema';

export type Paths = keyof paths;

export type Methods = keyof UnionToIntersection<paths[keyof paths]>;

export type MethodsFilterByPath<Path extends Paths> = Methods &
  keyof UnionToIntersection<paths[Path]>;

export type RequestParameters<Path extends Paths, Method extends Methods> = Get<
  paths,
  `${Path}.${Method}.parameters.query`
>;

export type RequestData<Path extends Paths, Method extends Methods> = Get<
  paths,
  `${Path}.${Method}.requestBody.content.application/json`
>;

export type ResponseCodes<Path extends Paths, Method extends Methods> = Get<
  paths,
  `${Path}.${Method}.responses`
>;

export type ResponseData<
  Path extends Paths,
  Method extends Methods,
  Code extends number
> = Get<paths, `${Path}.${Method}.responses.${Code}.content.application/json`>;

export type SuccessResponseData<
  Path extends Paths,
  Method extends Methods
> = ResponseData<Path, Method, 200>;

export const ErrorCodes = [400, 401, 403, 404, 500] as const;
type ErrorCode = (typeof ErrorCodes)[number];

export type ErrorResponseData<
  Path extends Paths,
  Method extends Methods,
  Code extends ErrorCode
> = Code extends number
  ? ResponseData<Path, Method, Code> extends never
    ? {
        status: Code;
        data: {
          message: string;
        };
      }
    : {
        status: Code;
        data: ResponseData<Path, Method, Code>;
      }
  : never;

export type ApiError<Path extends Paths, Method extends Methods> = ApiResponse<
  Path,
  Method,
  ErrorCode
>;
