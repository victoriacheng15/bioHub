<script setup lang="ts">
import SocialList from "./SocialList.vue";
import avatar from "/avatar.jpg"

const FIXEDSIZE = 175

const { name, headline } = defineProps<{ name: string; headline: string }>();
const imageInfo = {
  src: avatar,
  alt: name,
  width: FIXEDSIZE,
  height: FIXEDSIZE,
}
</script>

<template>
  <header>
    <img :="imageInfo" />
    <h1>{{ name }} </h1>
    <p> {{ headline }}</p>
    <SocialList />
  </header>
</template>

<style scoped>
header {
  margin-top: 4rem;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.25rem;
  text-align: center;
}

header h1 {
  font-size: 2.5rem;
  color: #ffa375;
  letter-spacing: 2px;
}

header p {
  font-size: 1.25rem;
}

img {
  display: block;
  border-radius: 100%;
}
</style>