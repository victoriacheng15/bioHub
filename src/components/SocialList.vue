<script setup lang="ts">
import { SCOIALS } from "../constant";
import SocialElement from "./SocialElement.vue"
</script>

<template>
  <ul>
    <li v-for="{ title, path, href } in SCOIALS " :key="title">
      <SocialElement :href="href" :path="path" :altText="title" />
    </li>
  </ul>
</template>

<style scoped>
ul {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  gap: 1.25rem;
}

ul li {
  padding: 0.25rem;
  background: #1147bb;
  border-radius: 5px;
  transition: padding 500ms ease-in-out;
}

ul li:hover {
  padding-inline: 0.75rem;

}
</style>