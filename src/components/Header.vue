<script setup lang="ts">
import { css } from "../../styled-system/css"
import avatar from "/avatar.jpg"
const { name, headline } = defineProps<{ name: string; headline: string }>();
const imageInfo = {
  src: avatar,
  alt: name,
  width: 200,
  height: 200,
}
</script>

<template>
  <header :class="css({
    display: 'flex', flexDir: 'column', justifyContent: 'center',
    alignItems: 'center', gap: '6', marginBlockEnd: '12'
  })">
    <img :="imageInfo" :class="css({ marginInline: 'auto', borderRadius: '50%' })" />
    <h1 :class="css({ textAlign: 'center', fontSize: '4xl' })">{{ name }}</h1>
    <p :class="css({ textAlign: 'center', fontSize: 'xl' })">{{ headline }}</p>
    <h2>test, this will have a list of social media links</h2>
  </header>
</template>