import { defineConfig, mergeConfig } from 'vitest/config';
import viteConfig from './vite.config';

export default mergeConfig(
  viteConfig,
  defineConfig({
    esbuild: {
      jsxInject: "import React from 'react'",
    },
    test: {
      environment: 'jsdom',
      globals: true,
    },
  })
);
