import { useCSRFToken } from '@/components/hooks/useCsrfToken';
import { FC, ReactNode, createContext, useContext } from 'react';

const CsrfTokenContext = createContext<string>('');

type Props = {
  children: ReactNode;
};

const CsrfTokenProvider: FC<Props> = ({ children }) => {
  const token = useCSRFToken();
  return (
    <CsrfTokenContext.Provider value={token}>
      {children}
    </CsrfTokenContext.Provider>
  );
};

export const useCsrfToken = () => useContext(CsrfTokenContext);

export default CsrfTokenProvider;
