import { Link } from '@/components/Parts/Navigation';
import { DataTable, GridColDef } from '@/components/Parts/Table';
import { FC, memo } from 'react';
import { Url } from '../graphql';

type DataType = Omit<Url, '__typename'>;

type Props = {
  data: DataType[];
  deleteHandler: (key: string) => () => Promise<void>;
};

type columnsProps = {
  deleteHandler: (key: string) => () => Promise<void>;
};

type TColumns = (props: columnsProps) => GridColDef[];

const columns: TColumns = ({ deleteHandler }) => {
  return [
    DataTable.DeleteActionColumn({
      handler(param) {
        return deleteHandler(param.row.key);
      },
    }),
    {
      field: 'key',
      headerName: '短縮パス',
      width: 400,
      renderCell: (param) => {
        return <Link to={`/shorts/${param.value}`}>{param.value}</Link>;
      },
    },
    {
      field: 'url',
      headerName: 'オリジナルURL',
      width: 500,
    },
  ];
};

export const Table: FC<Props> = memo(({ data, deleteHandler }) => (
  <DataTable
    columns={columns({
      deleteHandler: deleteHandler,
    })}
    rows={data}
    getRowId={(row) => row.key}
  />
));

Table.displayName = 'ShortsTable';
