<script setup lang="ts">
const { href, path, altText } = defineProps<{ href: string; path?: string, altText: string }>();

const FIXEDSIZE = 45
</script>

<template>
  <a :href="href" target="_blank" rel="noopener noreferrer">
    <img :src="path" :alt="altText" :width="FIXEDSIZE" :height="FIXEDSIZE" />
  </a>
</template>

<style scoped>
img {
  padding: 1px;
  background: #f1f5f9;
  border-radius: 5px;
}
</style>