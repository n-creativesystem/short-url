export type UserInfo = {
  sub: string;
  profile: string;
  email: string;
  emailVerified: boolean;
  userName: string;
  picture: string;
};

export type UserInfoContext = {
  loading: boolean;
  error?: Error;
  userInfo?: BaseUserInfo;
};

export type FetchResponse<T> = {
  response?: T;
  isLoading: boolean;
  error?: Error;
};
