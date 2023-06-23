export type { GridColDef, GridValueGetterParams } from '@mui/x-data-grid';
import {
  DataGrid,
  DataGridProps,
  GridPaginationInitialState,
  jaJP,
} from '@mui/x-data-grid';
import classNames from 'classnames';
import { FC } from 'react';
import styles from './index.module.scss';

const cx = classNames.bind(styles);

export type Props = {
  className?: string;
} & DataGridProps;

export const DataTable: FC<Props> = ({ className, ...props }) => {
  const pagination: GridPaginationInitialState = {
    paginationModel: {
      page: 0,
      pageSize: 50,
      ...(props.initialState?.pagination?.paginationModel || {}),
    },
  };
  return (
    <div className={cx(styles.container, className)}>
      <DataGrid
        {...props}
        initialState={{
          pagination: pagination,
          ...props.initialState,
        }}
        localeText={jaJP.components.MuiDataGrid.defaultProps.localeText}
      />
    </div>
  );
};
