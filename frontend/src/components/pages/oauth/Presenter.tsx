import { OAuthApplication } from '@/graphql/generated';
import { FC, memo } from 'react';
import { Table } from './Table';

type Props = {
  data: OAuthApplication[];
  deleteHandler: (id: string) => () => Promise<void>;
};

const Presenter: FC<Props> = memo(({ data, deleteHandler }) => {
  return (
    <>
      <Table data={data} deleteHandler={deleteHandler} />
    </>
  );
});

Presenter.displayName = 'OAuthPresenter';

export default Presenter;
