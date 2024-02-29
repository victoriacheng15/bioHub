<script setup lang="ts">
import { css } from "../../styled-system/css"
import SocialList from "./SocialList.vue";
import avatar from "/avatar.jpg"

const FIXEDSIZE = 200

const { name, headline } = defineProps<{ name: string; headline: string }>();
const imageInfo = {
  src: avatar,
  alt: name,
  width: FIXEDSIZE,
  height: FIXEDSIZE,
}
</script>

<template>
  <header :class="css({ w: '100%', mb: { base: '8', md: '12' } })">
    <div :class="css({ display: 'grid', gap: '4', mb: '8' })">
      <img :="imageInfo" :class="css({ mx: 'auto', rounded: '50%' })" />
      <h1 :class="css({ textAlign: 'center', fontSize: '4xl', fontWeight: 'semibold', letterSpacing: 'wide' })">{{ name
      }}
      </h1>
      <p
        :class="css({ textAlign: 'center', fontSize: { base: 'xl', sm: '2xl' }, transition: 'ease-in-out', transitionDuration: '400ms' })">
        {{ headline }}</p>
    </div>
    <SocialList />
  </header>
</template>