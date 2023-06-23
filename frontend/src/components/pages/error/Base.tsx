import Container from '@mui/material/Container';
import Link from '@mui/material/Link';
import Stack from '@mui/material/Stack';
import { FC, ReactNode, memo } from 'react';
import { Item } from './Item';

type Props = {
  children: ReactNode;
  showTopPage?: boolean;
};

export const BaseError: FC<Props> = memo(({ children, showTopPage = true }) => {
  return (
    <Container>
      <Stack spacing={1}>
        {children}
        {showTopPage && (
          <Item>
            <Link href="/" variant="h5" underline="none" color="primary">
              トップページへ
            </Link>
          </Item>
        )}
      </Stack>
    </Container>
  );
});

BaseError.displayName = 'BaseError';
