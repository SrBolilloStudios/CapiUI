<script setup lang="ts">
import { computed, watchEffect } from 'vue'
import { useI18n } from 'vue-i18n'
import { en, es } from '@nuxt/ui/locale'
import { useSettingsStore } from '@/stores/settings'

const settings = useSettingsStore()
const { locale } = useI18n()

locale.value = settings.locale

const uiLocales = { es, en } as const
const uiLocale = computed(() => uiLocales[locale.value as keyof typeof uiLocales] ?? es)

watchEffect(() => {
  settings.locale = locale.value as typeof settings.locale
  document.documentElement.lang = locale.value
})
</script>

<template>
  <UApp :locale="uiLocale">
    <RouterView />
  </UApp>
</template>
