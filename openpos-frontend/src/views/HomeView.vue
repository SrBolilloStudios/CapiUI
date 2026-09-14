<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
</script>

<template>
  <div class="flex min-h-screen flex-col items-center justify-center gap-4">
    <h1 class="text-2xl font-bold">{{ t('home.title') }}</h1>
    <p class="text-neutral-600 dark:text-neutral-400">{{ t('home.subtitle') }}</p>
    <UButton>{{ t('home.cta') }}</UButton>
  </div>
</template>
