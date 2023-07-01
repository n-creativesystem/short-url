import { MenuDrawer } from '@/components/Parts/MenuDrawer';
import SubHeader from '@/components/Parts/SubHeader';
import { useApolloClient } from '@/components/hooks/useApollo';
import varStyles from '@/styles/variables/index.module.scss';
import { ApolloProvider } from '@apollo/client';
import Box from '@mui/material/Box';
import { FC, Suspense, memo, useEffect, useState } from 'react';
import { Outlet } from 'react-router-dom';
import { useLocationChange } from '../Navigation';
import { AppBar } from './Appbar';
import AvatarParts from './Headers/Avatar';

const Layout: FC = memo(() => {
  useLocationChange((_location) => {
    setTitle('');
  });
  const client = useApolloClient();
  const [title, setTitle] = useState('');
  const useSetTitle = (title: string) => {
    useEffect(() => {
      setTitle(title);
    });
  };
  return (
    <ApolloProvider client={client}>
      <AppBar title="QuickLink" auth={<AvatarParts />} />
      <MenuDrawer />
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          height: '100vh',
          overflow: 'auto',
        }}
      >
        <SubHeader title={title} />
        <Box sx={{ p: 4, height: varStyles.heightPageContent }}>
          <Suspense fallback={<div>Loading...</div>}>
            <Outlet context={{ setTitle: useSetTitle }} />
          </Suspense>
        </Box>
      </Box>
    </ApolloProvider>
  );
});

Layout.displayName = 'Layout';

export default Layout;
