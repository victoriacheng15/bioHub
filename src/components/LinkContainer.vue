<script setup lang="ts">
import { LINKS } from "../constant";
</script>

<template>
  <main>
    <ul>
      <li v-for="{ title, href } in LINKS" :key="title">
        <a :href="href" target="_blank" rel="noopener noreferrer">
          {{ title }}
        </a>
      </li>
    </ul>
  </main>
</template>

<style scoped>
main {
  margin-block: 2rem;
}

ul {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1.75rem;
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

a {
  color: #f1f5f9;
  font-size: 1.5rem;
  font-weight: bold;
  letter-spacing: 2px;
}
</style>
