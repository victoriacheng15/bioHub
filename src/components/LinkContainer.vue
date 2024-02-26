<script setup lang="ts">
import { css } from "../../styled-system/css"
import { LINKS } from "../constant";

const styles = css({
  w: '100%',
  h: '100%',
  bg: 'gray.900',
  p: '4',
  fontSize: '2xl',
  fontWeight: 'bold',
  letterSpacing: 'widest',
  textAlign: 'center',
  rounded: 'lg',
  transition: 'ease-in-out',
  transitionDuration: '400ms',
  _hover: { bg: 'blue.400', color: 'gray.900', transition: 'ease-in-out', transitionDuration: '400ms' }
})
</script>

<template>
  <main :class="css({ w: '100%', mb: '14' })">
    <ul :class="css({ display: 'grid', gap: '8' })">
      <li v-for="{ title, href } in LINKS" :key="title"
        :class="css({ w: '100%', h: '100%', display: 'grid', placeItems: 'center' })">
        <a :href="href" target="_blank" rel="noopener noreferrer" :class="styles">
          {{ title }}
        </a>
      </li>
    </ul>
  </main>
</template>