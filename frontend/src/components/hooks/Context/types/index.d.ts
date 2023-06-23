type Modal = {
  open: boolean;
};

export type ErrorModal = {
  title: string;
  description: string;
} & Modal;

type ErrorPage = {
  show: boolean;
};

export type InternalErrorPage = ErrorPage;
