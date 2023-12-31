import react from '@vitejs/plugin-react';
import { fileURLToPath, URL } from 'node:url';
import { defineConfig } from 'vite';

// https://vitejs.dev/config/
const config = defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
      '@t': fileURLToPath(new URL('./src/types', import.meta.url)),
      '@gql': fileURLToPath(new URL('./src/graphql', import.meta.url)),
    },
  },
  server: {
    port: 3000,
    watch: {
      usePolling: true,
    },
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  preview: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
    open: true,
  },
});

export default config;
