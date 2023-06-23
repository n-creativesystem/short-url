export const STATUS_CODES = {
  SUCCESS: {
    code: 200,
    key: 'SUCCESS',
  },
  NOT_FOUND: {
    code: 404,
    key: 'NOT_FOUND',
  },
  INTERNAL_SERVER_ERROR: {
    code: 500,
    key: 'INTERNAL_SERVER_ERROR',
  },
  UNAUTHORIZED: {
    code: 401,
    key: 'UNAUTHORIZED',
  },
  FORBIDDEN: {
    code: 403,
    key: 'FORBIDDEN',
  },
  MAINTENANCE: {
    code: 503,
    key: 'maintenance',
  },
  INVALID: {
    code: 400,
    key: 'invalid',
  },
  UNKNOWN: {
    code: 0,
    key: 'unknown',
  },
} as const;

const _STATUS_CODES = Object.values(STATUS_CODES).map((item) => item.code);
export type TStatusCodes = (typeof _STATUS_CODES)[number];
