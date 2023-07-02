import { DeleteAction } from './Actions';
import {
  DataTable as BaseDataTable,
  GridColDef,
  GridRenderCellParams,
} from './DataTable';

type InternalDataTable = typeof BaseDataTable;

type TDeleteActionColumn = (args: {
  handler: (param: GridRenderCellParams) => () => Promise<void>;
  fieldName?: string;
  width?: number;
  headerName?: string;
}) => GridColDef;

type DataTableComponent = InternalDataTable & {
  DeleteAction: typeof DeleteAction;
  DeleteActionColumn: TDeleteActionColumn;
};

const DataTable = BaseDataTable as DataTableComponent;

DataTable.DeleteAction = DeleteAction;
DataTable.DeleteActionColumn = ({
  handler,
  fieldName = 'delete-action',
  headerName = '操作',
  width = 80,
}) => {
  return {
    field: fieldName,
    headerName: headerName,
    width: width,
    renderCell: (param) => {
      return <DeleteAction handler={handler(param)} />;
    },
  };
};

export { RegisterButton } from './Actions';
export type { InternalDataTable, DataTableComponent as Component, GridColDef };
export { DataTable };
