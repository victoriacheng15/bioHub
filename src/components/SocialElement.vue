<script setup lang="ts">
const { href, path, altText } = defineProps<{ href: string; path?: string, altText: string }>();

const FIXEDSIZE = 45
</script>

<template>
  <a :href="href" target="_blank" rel="noopener noreferrer">
    <img :src="path" :alt="altText" :width="FIXEDSIZE" :height="FIXEDSIZE" />
  </a>
</template>