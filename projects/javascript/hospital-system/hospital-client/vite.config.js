import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  // 开发环境下把 /api 请求代理到本地后端，避免在代码中硬编码 host:port
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:3000',
        changeOrigin: true,
        secure: false,
        // 保持路径不变
        rewrite: (path) => path,
      },
    },
  },
})
