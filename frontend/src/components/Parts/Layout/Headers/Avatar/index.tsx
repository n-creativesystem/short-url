import { useUserInfoContext } from '@/components/Parts/Layout/UserInfo';
import IconButton from '@mui/material/IconButton';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import { FC, MouseEvent, ReactNode, useState } from 'react';
import style from './index.module.scss';

type AvatarImageProps = {
  src: string | undefined;
};

const AvatarImage: FC<AvatarImageProps> = ({ src }) => {
  return (
    <img
      className={style.avatar}
      alt=""
      src={src || ''}
      width="100%"
      height="100%"
    />
  );
};

type Item = {
  key: string;
  label: ReactNode;
};

const AvatarParts: FC = () => {
  const context = useUserInfoContext();
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const handlerMenu = (e: MouseEvent<HTMLElement>) => {
    setAnchorEl(e.currentTarget);
  };
  const handlerClose = () => {
    setAnchorEl(null);
  };
  const items: Item[] = [
    {
      key: 'logout',
      label: <a href="/api/auth/logout">ログアウト</a>,
    },
  ];
  return context?.userInfo ? (
    <>
      <IconButton
        size="large"
        aria-label="account of current user"
        aria-controls="menu-appbar"
        aria-haspopup="true"
        onClick={handlerMenu}
        color="inherit"
        className={style.button}
      >
        {<AvatarImage src={context?.userInfo?.picture} />}
      </IconButton>
      <Menu
        id="menu-appbar"
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={handlerClose}
      >
        {items.map((item, idx) => (
          <MenuItem key={`${item.label}_${idx}`} onClick={handlerClose}>
            {item.label}
          </MenuItem>
        ))}
      </Menu>
    </>
  ) : (
    <></>
  );
};

export default AvatarParts;
