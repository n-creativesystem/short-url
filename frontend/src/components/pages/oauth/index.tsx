import { Button } from '@/components/Parts/Button';
import { useRouter } from '@/components/Parts/Navigation';
import { useOAuthApplicationsQuery } from '@/graphql/generated';
import classNames from 'classnames';
import { FC, memo } from 'react';
import Presenter from './Presenter';
import { useDeleteHandler } from './Table/Actions';
import styles from './index.module.scss';

const cx = classNames.bind(styles);

type Props = {};

const OAuth2ApplicationsContainer: FC<Props> = memo(({}) => {
  const router = useRouter();
  const { data, loading } = useOAuthApplicationsQuery({
    variables: {
      token: '',
    },
  });
  const onRegisterClick = () => {
    router.push('/oauth2/app/register');
  };

  const deleteHandler = useDeleteHandler();

  return loading ? (
    <></>
  ) : (
    <>
      <div className={cx(styles['register-button'])}>
        <Button variant="contained" color="primary" onClick={onRegisterClick}>
          新規作成
        </Button>
      </div>
      <Presenter
        data={data?.oauthApplications?.result || []}
        deleteHandler={deleteHandler}
      />
    </>
  );
});

OAuth2ApplicationsContainer.displayName = 'OAuth2ApplicationContainer';

export default OAuth2ApplicationsContainer;
