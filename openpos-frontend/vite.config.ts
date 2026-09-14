import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import ui from "@nuxt/ui/vite";
import VueI18nPlugin from '@intlify/unplugin-vue-i18n/vite'

export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
    VueI18nPlugin({
      include: [fileURLToPath(new URL('./src/i18n/locales/**', import.meta.url))],
    }),
    ui({
      ui: {
        colors:
        {
          primary: 'indigo',
          neutral: 'slate'
        }
      },
      experimental: { componentDetection: true }
    })
  ],
  server: {
    port: 5173,
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true
      }
    }
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
