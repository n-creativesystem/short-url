export type UserInfo = {
  email?: string | undefined;
  email_verified?: boolean | undefined;
  picture?: string | undefined;
  profile?: string | undefined;
  sub?: string | undefined;
  username?: string | undefined;
};

export type UserInfoContext = {
  loading: boolean;
  error?: Error;
  userInfo?: UserInfo;
};

export type FetchResponse<T> = {
  response?: T;
  isLoading: boolean;
  error?: Error;
};
