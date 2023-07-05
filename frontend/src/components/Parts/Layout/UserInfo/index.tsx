import { openApiClient } from '@/components/hooks/useFetch';
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
  useEffect(() => {
    const innerValue: TUserInfoContext = {
      loading: true,
    };
    const fetch = async () => {
      try {
        const data = await openApiClient.auth.userinfo.$get();
        if (data) {
          innerValue.userInfo = {
            ...data,
          };
        }
      } catch (error) {
        innerValue.error = error as Error;
      } finally {
        innerValue.loading = false;
      }
      setValue(innerValue);
    };
    fetch();
  }, []);
  return (
    <UserInfoContext.Provider value={value}>
      {children}
    </UserInfoContext.Provider>
  );
};

export const useUserInfoContext = () => useContext(UserInfoContext);

export default UserInfoProvider;
