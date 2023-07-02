import { RegisterButton } from '@/components/Parts/Table';
import { FC, memo } from 'react';
import { Table } from './Table';
import { ResultFragment } from './graphql/fragment';
type Props = {
  data: Array<ResultFragment>;
  registerHandler: () => void;
  deleteHandler: (key: string) => () => Promise<void>;
};

export const Presenter: FC<Props> = memo(
  ({ data, registerHandler, deleteHandler }) => {
    return (
      <>
        <RegisterButton handler={registerHandler} />
        <Table data={data} deleteHandler={deleteHandler} />
      </>
    );
  }
);
Presenter.displayName = 'Presenter';
