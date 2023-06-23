import CircularProgress from '@mui/material/CircularProgress';
import IconButton from '@mui/material/IconButton';
import { FC, useState } from 'react';
import { Button, ButtonProps } from './Button';
import { onClickEvent, onClickHandler } from './index.d';

export type Props = {
  onClick?: onClickHandler;
  loadingIcon?: boolean;
  icon?: boolean;
} & Omit<ButtonProps, 'onClick'>;

export const AsyncButton: FC<Props> = ({
  onClick,
  loadingIcon = true,
  icon,
  ...props
}) => {
  const handler = useWrapClick(onClick)();
  if (icon) {
    return (
      <IconButton
        {...props}
        onClick={handler.handler}
        disabled={handler.loading}
      >
        {handler.loading ? <CircularProgress size={20} /> : props.children}
      </IconButton>
    );
  }
  return (
    <Button
      {...props}
      onClick={handler.handler}
      loading={loadingIcon ? handler.loading : false}
      disabled={handler.loading}
    />
  );
};

const useWrapClick = (onHandle?: onClickHandler) => {
  const [loading, setLoading] = useState(false);
  return () => {
    if (!onHandle)
      return {
        loading: false,
        handler: undefined,
      };
    const handler = async (e: onClickEvent) => {
      setLoading(true);
      try {
        await onHandle(e);
      } finally {
        setLoading(false);
      }
    };
    return {
      loading: loading,
      handler: handler,
    };
  };
};
