<script setup lang="ts">
import { LINKS } from "../constant";
import LinkElement from "./LinkElement.vue";
</script>

<template>
  <main>
    <ul>
      <li v-for="{ title, href } in LINKS" :key="title">
        <LinkElement :href="href" :title="title" />
      </li>
    </ul>
  </main>
</template>

<style scoped>
main {
  margin: 3rem 0;
}

ul {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1.5rem;
}

ul li {
  display: flex;
  justify-content: center;
  padding: 0.5rem;
  width: 100%;
  background: #1147bb;
  border-radius: 5px;
  transition: background 500ms ease-in-out;
  cursor: pointer;
}

ul li:hover {
  background: hsl(221, 83%, 20%);
}
</style>
