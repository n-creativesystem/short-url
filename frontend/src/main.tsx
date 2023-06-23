import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import * as React from 'react';
import * as ReactDOM from 'react-dom/client';
import { App } from './app';
import './index.scss';

const rootElement = document.getElementById('root');
const root = ReactDOM.createRoot(rootElement as HTMLElement);

root.render(
  /**
   NOTE: https://ja.legacy.reactjs.org/docs/strict-mode.html
   開発環境の場合2回レンダリングされる mount → unmount → mount
   */
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
