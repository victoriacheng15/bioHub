<script setup lang="ts">
import { SCOIALS } from "../constant";

const FIXEDSIZE = 45
</script>

<template>
  <ul>
    <li v-for="{ title, path, href } in SCOIALS " :key="title">
      <a :href="href" target="_blank" rel="noopener noreferrer">
        <img :src="path" :alt="title" :width="FIXEDSIZE" :height="FIXEDSIZE" />
      </a>
    </li>
  </ul>
</template>

<style scoped>
ul {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-wrap: wrap;
  gap: 2rem;
}

ul li {
  padding: 0.5rem;
  background: #60a5fa;
  border-radius: 10px;
  transition: background 500ms ease-in-out;
}

ul li:hover {
  background: #1147bb;
}
</style>