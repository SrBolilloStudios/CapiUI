import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import ui from '@nuxt/ui/vite'
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
        colors: {
          primary: 'teal',
          neutral: 'olive',
        },
        icons: {
          arrowDown: 'i-iconoir-arrow-down',
          arrowLeft: 'i-iconoir-arrow-left',
          arrowRight: 'i-iconoir-arrow-right',
          arrowUp: 'i-iconoir-arrow-up',
          caution: 'i-iconoir-warning-circle',
          check: 'i-iconoir-check',
          chevronDoubleLeft: 'i-iconoir-fast-arrow-left',
          chevronDoubleRight: 'i-iconoir-fast-arrow-right',
          chevronDown: 'i-iconoir-nav-arrow-down',
          chevronLeft: 'i-iconoir-nav-arrow-left',
          chevronRight: 'i-iconoir-nav-arrow-right',
          chevronUp: 'i-iconoir-nav-arrow-up',
          close: 'i-iconoir-xmark',
          copy: 'i-iconoir-copy',
          copyCheck: 'i-iconoir-clipboard-check',
          dark: 'i-iconoir-half-moon',
          drag: 'i-iconoir-menu',
          ellipsis: 'i-iconoir-more-horiz',
          error: 'i-iconoir-xmark-circle',
          external: 'i-iconoir-arrow-up-right',
          eye: 'i-iconoir-eye',
          eyeOff: 'i-iconoir-eye-closed',
          file: 'i-iconoir-page',
          folder: 'i-iconoir-folder',
          folderOpen: 'i-iconoir-folder',
          hash: 'i-iconoir-hashtag',
          info: 'i-iconoir-info-circle',
          light: 'i-iconoir-sun-light',
          loading: 'i-iconoir-refresh',
          menu: 'i-iconoir-menu',
          minus: 'i-iconoir-minus',
          panelClose: 'i-iconoir-sidebar-collapse',
          panelOpen: 'i-iconoir-sidebar-expand',
          plus: 'i-iconoir-plus',
          reload: 'i-iconoir-undo',
          search: 'i-iconoir-search',
          stop: 'i-iconoir-square',
          star: 'i-iconoir-star',
          success: 'i-iconoir-check-circle',
          system: 'i-iconoir-computer',
          tip: 'i-iconoir-light-bulb',
          upload: 'i-iconoir-upload',
          warning: 'i-iconoir-warning-triangle',
        },
        input: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
          slots: {
            base: 'rounded-full',
          },
        },
        select: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
          slots: {
            base: 'rounded-full',
          },
        },
        textarea: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
        },
        selectMenu: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
          slots: {
            base: 'rounded-full',
          },
        },
        inputMenu: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
          slots: {
            base: 'rounded-full',
          },
        },
        inputNumber: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
        },
        inputTags: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
        },
        inputDate: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
        },
        inputTime: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
        },
        pinInput: {
          defaultVariants: {
            variant: 'soft',
            size: 'lg',
          },
        },
        button: {
          defaultVariants: {
            size: 'lg',
          },
          slots: {
            base: 'rounded-full',
          },
        },
        badge: {
          defaultVariants: {
            size: 'lg',
          },
        },
        inputRating: {
          defaultVariants: {
            size: 'lg',
          },
        },
        tabs: {
          defaultVariants: {
            size: 'lg',
          },
        },
        checkbox: {
          defaultVariants: {
            size: 'lg',
          },
        },
        checkboxGroup: {
          defaultVariants: {
            size: 'lg',
          },
        },
        radioGroup: {
          defaultVariants: {
            size: 'lg',
          },
        },
        switch: {
          defaultVariants: {
            size: 'lg',
          },
        },
        slider: {
          defaultVariants: {
            size: 'lg',
          },
        },
        stepper: {
          defaultVariants: {
            size: 'lg',
          },
        },
        calendar: {
          defaultVariants: {
            size: 'lg',
          },
        },
        colorPicker: {
          defaultVariants: {
            size: 'lg',
          },
        },
        fileUpload: {
          defaultVariants: {
            size: 'lg',
          },
        },
        formField: {
          defaultVariants: {
            size: 'lg',
          },
        },
        fieldGroup: {
          defaultVariants: {
            size: 'lg',
          },
        },
        dropdownMenu: {
          defaultVariants: {
            size: 'lg',
          },
        },
        contextMenu: {
          defaultVariants: {
            size: 'lg',
          },
        },
        commandPalette: {
          defaultVariants: {
            size: 'lg',
          },
        },
        listbox: {
          defaultVariants: {
            size: 'lg',
          },
        },
      },
    }),
  ],
  server: {
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
})
