<script setup lang="ts">
import { css } from "../../styled-system/css"
const { href, title } = defineProps<{ href: string, title: string }>()

const styles = css({
  w: '100%',
  h: '100%',
  bg: 'gray.900',
  p: '4',
  fontSize: { base: 'xl', md: '2xl' },
  fontWeight: 'semibold',
  letterSpacing: 'widest',
  textAlign: 'center',
  rounded: 'lg',
  transition: 'ease-in-out',
  transitionDuration: '400ms',
  _hover: { bg: 'blue.400', color: 'gray.900', transition: 'ease-in-out', transitionDuration: '400ms' }
})
</script>

<template>
  <a :href="href" target="_blank" rel="noopener noreferrer" :class="styles">
    {{ title }}
  </a>
</template>