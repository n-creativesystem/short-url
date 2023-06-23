type TLogin = {
  state: string;
  code: string;
};

export type ContainerProps = Partial<TLogin> &
  Partial<{
    isCallback: true;
  }>;

export type GetOAuthButton = {
  buttons: string[];
  isLoading: boolean;
};
