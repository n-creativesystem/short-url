import { Menu } from '@/components/Parts/MenuDrawer';
import LinkIcon from '@mui/icons-material/Link';
import SettingsApplicationsIcon from '@mui/icons-material/SettingsApplications';
import { FC } from 'react';

export const MenuItem: FC = () => {
  return (
    <>
      {/* <ListItemButton to="/" selected={pathname === '/'}>
          <ListItemIcon>
            <HomeIcon />
          </ListItemIcon>
          <ListItemText primary="トップページ" />
        </ListItemButton> */}
      <Menu.ListItemButton to="/shorts" path="/shorts">
        <Menu.ListItemButton.ListItemIcon>
          <LinkIcon />
        </Menu.ListItemButton.ListItemIcon>
        <Menu.ListItemButton.ListItemText primary="生成済みURL" />
      </Menu.ListItemButton>
      <Menu.ListItemButton to="/oauth2/app" path={'/oauth2/app'}>
        <Menu.ListItemButton.ListItemIcon>
          <SettingsApplicationsIcon />
        </Menu.ListItemButton.ListItemIcon>
        <Menu.ListItemButton.ListItemText primary="アプリケーション" />
      </Menu.ListItemButton>
    </>
  );
};
