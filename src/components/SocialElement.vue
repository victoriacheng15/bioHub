<script setup lang="ts">
const { href, path, altText } = defineProps<{ href: string; path?: string, altText: string }>();
</script>

<template>
  <a :href="href" target="_blank" rel="noopener noreferrer">
    <img :src="path" :alt="altText" width="45" height="45" />
  </a>
</template>