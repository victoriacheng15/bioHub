<script setup lang="ts">
import { css } from "../../styled-system/css"
import { SCOIALS } from "../constant";
import SocialElement from "./SocialElement.vue"

const styles = css({
  m: '1',
  p: '1.5',
  bg: 'blue.400',
  rounded: 'md',
  transition: 'ease-in-out',
  transitionDuration: '400ms',
  _hover: { scale: '1.25', transition: 'ease-in-out', transitionDuration: '400ms' }
})
</script>

<template>
  <ul :class="css({ display: 'flex', justifyContent: 'center', gap: { base: '4', sm: '6' }, flexWrap: 'wrap' })">
    <li v-for="{ title, path, href } in  SCOIALS " :key="title" :class="styles">
      <SocialElement :href="href" :path="path" :altText="title" />
    </li>
  </ul>
</template>