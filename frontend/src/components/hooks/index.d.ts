export type UserInfo = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [key: string]: any;
};

export type FetchResponse<T> = {
  response?: T;
  isLoading: boolean;
  error?: Error;
};
