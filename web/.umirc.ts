import { defineConfig } from '@umijs/max';

export default defineConfig({
  routes: [
    { path: '/', component: 'index' },
    { path: '/docs', component: 'docs' },
    { path: '/login', component: 'login' },
    { path: '/callback', component: 'callback' },
  ],
  npmClient: 'pnpm',
  utoopack: {},
  styles: ['@/global.less'],
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true,
      secure: false,
    },
  },
  locale: {
    default: 'zh-CN',
    baseSeparator: '-',
  },
});
