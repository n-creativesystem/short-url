import { useUserInfoContext } from '@/components/Parts/Layout/UserInfo';
import { useRouter } from '@/components/Parts/Navigation';
import { FC, useEffect } from 'react';

type Props = {
  Component?: FC;
};

export const Auth: FC<Props> = ({ Component }) => {
  const router = useRouter();
  const userInfo = useUserInfoContext();
  useEffect(() => {
    if (userInfo?.loading) {
      return;
    }
    if (userInfo?.error) {
      router.push('/auth');
      return;
    }
    if (!userInfo?.userInfo) {
      router.push('/auth');
      return;
    }
  }, [userInfo]);
  if (userInfo?.loading) {
    return <></>;
  }
  return (Component && <Component />) || <></>;
};
