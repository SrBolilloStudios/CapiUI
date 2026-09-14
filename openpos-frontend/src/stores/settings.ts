import { defineStore } from 'pinia'
import { ref } from 'vue'

import { defaultLocale, type SupportedLocale } from '@/i18n'

export const useSettingsStore = defineStore(
  'settings',
  () => {
    const locale = ref<SupportedLocale>(defaultLocale)

    return { locale }
  },
  {
    persist: true,
  },
)
