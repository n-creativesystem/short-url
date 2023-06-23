import { Dispatch, SetStateAction } from 'react';

type Modal = {
  open: boolean;
};

export type ModalContext = {
  open: boolean;
  onOpen: Dispatch<SetStateAction<boolean>>;
};

export type ErrorModal = {
  title: string;
  description: string;
} & Modal;

export type ErrorModalContext = {
  value: ErrorModal;
  setValue: (value: ErrorModal) => void;
};
