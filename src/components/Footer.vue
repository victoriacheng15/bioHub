<script setup lang="ts">
const { name } = defineProps<{ name: string }>();
</script>

<template>
  <footer>
    <p>
      Copyright &copy; <span aria-label="penguin emoji">🐧</span> {{ name }}
    </p>
  </footer>
</template>