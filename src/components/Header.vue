<script setup lang="ts">
import { css } from "../../styled-system/css"
import SocialList from "./SocialList.vue";
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
  <header :class="css({ w: '100%', mb: '14' })">
    <div :class="css({ display: 'grid', gap: '4', mb: '8' })">
      <img :="imageInfo" :class="css({ mx: 'auto', rounded: '50%' })" />
      <h1 :class="css({ textAlign: 'center', fontSize: '4xl' })">{{ name }}</h1>
      <p :class="css({ textAlign: 'center', fontSize: 'xl', })">{{ headline }}</p>
    </div>
    <SocialList />
  </header>
</template>