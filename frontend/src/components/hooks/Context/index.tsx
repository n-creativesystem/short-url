import { makeVar } from '@/components/hooks/reactive';
import { createContext } from 'react';
import type { ErrorModal, InternalErrorPage } from './types';

export const initialErrorModal: ErrorModal = {
  open: false,
  title: '',
  description: '',
};

export const errorModalVar = makeVar<ErrorModal>(initialErrorModal);

export const initialInternalErrorPage: InternalErrorPage = {
  show: false,
};

export const internalErrorPageVar = makeVar<InternalErrorPage>(
  initialInternalErrorPage
);

export const LoadingContext = createContext(false);
