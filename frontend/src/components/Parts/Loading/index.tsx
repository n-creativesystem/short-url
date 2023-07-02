import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import BaseSkeleton from '@mui/material/Skeleton';
import { SkeletonProps } from '@mui/material/Skeleton/Skeleton';
import { FC, memo } from 'react';

export type Props = {};

export const Loading: FC<Props> = memo(({}) => {
  return (
    <Box
      sx={{
        display: 'flex',
        width: '100%',
        height: '100%',
        justifyContent: 'center',
        alignItems: 'center',
      }}
    >
      <CircularProgress />
    </Box>
  );
});
Loading.displayName = 'Loading';

export const Skeleton: FC<SkeletonProps> = memo(({ ...props }) => {
  return <BaseSkeleton {...props} animation="wave" variant="rectangular" />;
});
