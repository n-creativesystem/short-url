import { TestingExports } from '@/lib/tests';
import { AxiosResponseHeaders, RawAxiosResponseHeaders } from 'axios';

const isHeaders = (headers: HeadersInit): headers is Headers =>
  headers.constructor?.name === 'Headers';

const arrToObject = (arr: Array<[string, string]>) => {
  const record: Record<string, string> = {};
  for (let i = 0; arr[i]; i++) {
    if (arr[i][1]) record[arr[i][0].toLocaleLowerCase()] = arr[i][1];
  }
  return record;
};

const headersToObject = (headers: Headers) => {
  return arrToObject(Array.from(headers.entries()));
};

const objectToObjectWithValueFilter = (headers: Record<string, string>) => {
  return arrToObject(Object.entries(headers));
};

const headersInitToObject = (headers: HeadersInit) => {
  if (isHeaders(headers)) {
    return headersToObject(headers);
  } else if (Array.isArray(headers)) {
    return arrToObject(headers);
  } else {
    return objectToObjectWithValueFilter(headers);
  }
};

export const createAxiosHeaders = (
  headers: HeadersInit = {}
): Record<string, string> => headersInitToObject(headers);

export const createFetchHeaders = (
  axiosHeaders: RawAxiosResponseHeaders | AxiosResponseHeaders = {}
): [string, string][] => {
  const headers: [string, string][] = [];
  Object.entries(axiosHeaders).forEach(([name, value]) => {
    headers.push([name, value]);
  });
  return headers;
};

export const testingExports = new TestingExports(
  isHeaders,
  arrToObject,
  headersToObject,
  objectToObjectWithValueFilter,
  headersInitToObject
);
