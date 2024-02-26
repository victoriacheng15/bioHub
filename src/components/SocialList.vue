<script setup lang="ts">
import { css } from "../../styled-system/css"
import { SCOIALS } from "../constant";

const styles = css({
  m: '2',
  p: '1.5',
  bg: 'blue.400',
  rounded: 'lg',
  transition: 'ease-in-out',
  transitionDuration: '400ms',
  _hover: { scale: '1.25', transition: 'ease-in-out', transitionDuration: '400ms' }
})
</script>

<template>
  <ul :class="css({ display: 'flex', justifyContent: 'center', gap: '6', flexWrap: 'wrap' })">
    <li v-for="{ title, path, href } in  SCOIALS " :key="title" :class="styles">
      <a :href="href" target="_blank" rel="noopener noreferrer">
        <img :src="path" :alt="title" width="50" height="50" />
      </a>
    </li>
  </ul>
</template>