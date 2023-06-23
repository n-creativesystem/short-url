import { MouseEvent } from 'react';

type onClickEvent =
  | MouseEvent<HTMLAnchorElement> & MouseEvent<HTMLButtonElement>;

type onClickHandler = (e?: onClickEvent) => Promise<void> | undefined;
