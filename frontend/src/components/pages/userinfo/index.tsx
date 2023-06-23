import { useUserInfoContext } from '@/components/Parts/Layout/UserInfo';
import { FC } from 'react';
import Presenter from './Presenter';
const UserInfoContainer: FC = () => {
  const data = useUserInfoContext();
  return <Presenter {...(data || undefined)} />;
};

export default UserInfoContainer;
