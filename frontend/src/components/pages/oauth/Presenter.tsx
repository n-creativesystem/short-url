import { RegisterButton } from '@/components/Parts/Table';
import { OAuthApplication } from '@t/graphql';
import classNames from 'classnames';
import { FC, memo } from 'react';
import { Table } from './Table';
import styles from './index.module.scss';

const cx = classNames.bind(styles);

type Props = {
  data: OAuthApplication[];
  registerHandler: () => void;
  deleteHandler: (id: string) => () => Promise<void>;
};

const Presenter: FC<Props> = memo(
  ({ data, registerHandler, deleteHandler }) => {
    return (
      <>
        <RegisterButton handler={registerHandler} />
        <Table data={data} deleteHandler={deleteHandler} />
      </>
    );
  }
);

Presenter.displayName = 'OAuthPresenter';

export default Presenter;
