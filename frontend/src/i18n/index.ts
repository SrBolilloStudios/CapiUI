import { createI18n } from 'vue-i18n'

import en from './locales/en.json'
import es from './locales/es.json'

export const defaultLocale = 'es'
export const supportedLocales = ['es', 'en'] as const

export type SupportedLocale = (typeof supportedLocales)[number]

export const i18n = createI18n({
  legacy: false,
  locale: defaultLocale,
  fallbackLocale: 'en',
  messages: { es, en },
})
