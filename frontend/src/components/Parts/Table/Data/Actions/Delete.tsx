import { AsyncButton } from '@/components/Parts/Button';
import DeleteIcon from '@mui/icons-material/Delete';
import { FC } from 'react';

type Props = {
  handler: () => Promise<void>;
};

export const DeleteAction: FC<Props> = ({ handler }) => {
  const onClick = (): Promise<void> => {
    return new Promise(async (resolve, reject) => {
      try {
        await handler();
        resolve(undefined);
      } catch (error) {
        reject(error);
      }
    });
  };
  return (
    <AsyncButton icon size="small" color="error" onClick={onClick}>
      <DeleteIcon />
    </AsyncButton>
  );
};

DeleteAction.displayName = 'DeleteAction';
