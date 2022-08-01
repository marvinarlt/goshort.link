import { resolve, join } from 'path';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@app': resolve(__dirname, 'src')
    }
  },
  build: {
    emptyOutDir: true,
    outDir: join(__dirname, 'dist')
  }
});
