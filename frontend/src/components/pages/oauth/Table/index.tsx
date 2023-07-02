import { Link } from '@/components/Parts/Navigation';
import { DataTable, GridColDef } from '@/components/Parts/Table';
import { OAuthApplication } from '@t/graphql';
import { FC, memo } from 'react';

type DataType = Omit<OAuthApplication, 'domain' | 'secret'>;

type Props = {
  data: DataType[];
  deleteHandler: (id: string) => () => Promise<void>;
};

type columnsProps = {
  deleteHandler: (id: string) => () => Promise<void>;
};

type TColumns = (props: columnsProps) => GridColDef[];

const columns: TColumns = ({ deleteHandler }) => {
  return [
    DataTable.DeleteActionColumn({
      handler(param) {
        return deleteHandler(param.row.id);
      },
    }),
    {
      field: 'id',
      headerName: 'ID',
      width: 400,
      renderCell: (param) => {
        return <Link to={`/oauth2/app/${param.value}`}>{param.value}</Link>;
      },
    },
    {
      field: 'name',
      headerName: 'アプリ名',
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
  />
));

Table.displayName = 'OAuthAppTable';
