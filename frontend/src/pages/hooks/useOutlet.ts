import { useOutletContext as useBaseOutletContext } from 'react-router-dom';

type OutletContext = {
  setTitle: (title: string) => void;
};

export const useOutletContext = () => {
  return useBaseOutletContext() as OutletContext;
};
