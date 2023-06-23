import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import Link from '@mui/material/Link';
import Stack from '@mui/material/Stack';
import { styled } from '@mui/material/styles';
import { FC, memo } from 'react';

const Item = styled(Box)(({ theme }) => ({
  padding: theme.spacing(1),
  textAlign: 'center',
}));

export const NotFound: FC = memo(() => {
  return (
    <Container>
      <Stack spacing={1}>
        <Item>アクセスしようとしたページはすでに存在しないか</Item>
        <Item>またはURLが間違っている可能性があります。</Item>
        <Item>
          <Link href="/" variant="h5" underline="none" color="primary">
            トップページへ
          </Link>
        </Item>
      </Stack>
    </Container>
  );
});

NotFound.displayName = 'NotFound';
