import { useFetch } from '@/components/hooks/useFetch';
import {
  FC,
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from 'react';
import { UserInfoContext as TUserInfoContext } from './index.d';

const UserInfoContext = createContext<TUserInfoContext | null | undefined>(
  undefined
);

type responseUserInfo = {
  sub: string;
  profile: string;
  email: string;
  email_verified: boolean;
  userName: string;
  picture: string;
};

type Props = {
  children: ReactNode;
};

const UserInfoProvider: FC<Props> = ({ children }) => {
  const [value, setValue] = useState<TUserInfoContext>({ loading: true });
  const { data, hasError, isLoading, error } = useFetch<responseUserInfo>({
    url: '/auth/userinfo',
  });
  useEffect(() => {
    const value: TUserInfoContext = {
      loading: isLoading,
    };
    if (hasError || error) {
      value.error = error;
    }
    if (data) {
      value.userInfo = {
        ...data,
        emailVerified: data.email_verified,
      };
    }
    setValue(value);
  }, [data, hasError, isLoading]);
  return (
    <UserInfoContext.Provider value={value}>
      {children}
    </UserInfoContext.Provider>
  );
};

export const useUserInfoContext = () => useContext(UserInfoContext);

export default UserInfoProvider;
