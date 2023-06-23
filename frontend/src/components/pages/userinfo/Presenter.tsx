import { FC, memo } from 'react';

type Props = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [key: string]: any;
};

const Presenter: FC<Props> = memo((props) => {
  return <div>{props ? JSON.stringify(props) : ''}</div>;
});

Presenter.displayName = 'UserInfoPresenter';

export default Presenter;
