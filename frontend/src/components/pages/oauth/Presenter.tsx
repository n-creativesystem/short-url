import { RegisterButton } from '@/components/Parts/Table';
import { OAuthApplication } from '@t/graphql';
import { FC, memo } from 'react';
import { Table } from './Table';

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
