import CheckIcon from '@mui/icons-material/Check';
import ContentCopyIcon from '@mui/icons-material/ContentCopy';
import Tooltip from '@mui/material/Tooltip';
import useCopy from '@react-hook/copy';
import { FC, forwardRef, useCallback } from 'react';
import { Button, ButtonProps } from './Button';

type Props = {
  value: string;
  toolTipText: string;
  visible: boolean;
  variant?: string;
} & ButtonProps;

export const CopyButton: FC<Props> = forwardRef(
  ({ value, toolTipText, visible, variant = 'text', ...props }, ref) => {
    const { copied, copy, reset } = useCopy(value);
    const onCopy = useCallback(() => {
      copy();
      setTimeout(() => {
        reset();
      }, 2000);
    }, [copy, reset]);

    return !visible ? (
      <></>
    ) : (
      <Tooltip placement="bottom" title={toolTipText}>
        <Button ref={ref} {...props} variant={variant} onClick={onCopy}>
          <Icon copied={copied} />
        </Button>
      </Tooltip>
    );
  }
);

type IconProps = {
  copied: boolean;
};

const Icon: FC<IconProps> = ({ copied }) => {
  return copied ? <CheckIcon /> : <ContentCopyIcon />;
};
