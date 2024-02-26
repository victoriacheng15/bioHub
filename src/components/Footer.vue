<script setup lang="ts">
import { css } from "../../styled-system/css"
const { name } = defineProps<{ name: string }>();
</script>

<template>
  <footer :class="css({ mt: 'auto', py: '4' })">
    <p>
      Copyright &copy; <span aria-label="penguin emoji">🐧</span> {{ name }}
    </p>
  </footer>
</template>