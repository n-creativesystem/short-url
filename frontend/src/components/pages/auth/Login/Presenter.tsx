import { onClickEvent } from '@/components/Parts/Button/index.d';
import { FC, memo } from 'react';
import GoogleLogin from './Google';
import OtherAuthButton from './Other';
import { GetOAuthButton } from './index.d';
import style from './index.module.scss';

type Props = {
  useAuthLogin: (label: string) => (e?: onClickEvent) => Promise<void>;
  enabledAuth: GetOAuthButton;
};

const LoginPresenter: FC<Props> = memo(({ useAuthLogin, enabledAuth }) => {
  const googleLogin = (
    <div key="google" style={{ margin: '20px' }}>
      <GoogleLogin onClick={useAuthLogin('google')} />
    </div>
  );
  const useOtherLogin = (label: string) => (
    <div key={label} style={{ marginBottom: '10px', marginTop: '10px' }}>
      <OtherAuthButton
        button={{ size: 'medium' }}
        label={label}
        icon={{ width: 20, height: 20 }}
        onClick={useAuthLogin(label)}
      />
    </div>
  );
  const otherLogin = useOtherLogin;
  const mapLoginButton: { [key: string]: JSX.Element } = {
    google: googleLogin,
  };

  return (
    <div className={style.container}>
      <div className={style.content}>
        <h2>ソーシャルアカウントでログイン</h2>
        {enabledAuth.isLoading ? (
          <></>
        ) : (
          <div className={style.description}>
            <p>続けるにはソーシャルアカウントでのログインが必要です。</p>
            <p>ログインページに遷移します。</p>
          </div>
        )}
        <div className={style['auth-button-content']}>
          {enabledAuth.buttons.map((k) => mapLoginButton[k] || otherLogin(k))}
        </div>
      </div>
    </div>
  );
});

LoginPresenter.displayName = 'LoginPresenter';

export default LoginPresenter;
