<script setup lang="ts">
const { href, title } = defineProps<{ href: string, title: string }>()

</script>

<template>
  <a :href="href" target="_blank" rel="noopener noreferrer">
    {{ title }}
  </a>
</template>

<style scoped>
a {
  font-size: 1.5rem;
  color: #ffa375;
  font-weight: bold;
  letter-spacing: 2px;
}
</style>