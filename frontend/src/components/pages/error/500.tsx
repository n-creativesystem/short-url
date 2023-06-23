import { FC, memo } from 'react';
import { BaseError } from './Base';
import { Item } from './Item';

export const InternalError: FC = memo(() => {
  return (
    <BaseError>
      <Item>ご不便をおかけして申し訳ございません。</Item>
      <Item>正常にご覧いただけるよう、解決に取り組んでいます。</Item>
    </BaseError>
  );
});

InternalError.displayName = 'InternalError';
