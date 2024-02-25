<script setup lang="ts">
import { LINKS } from "../constant";
</script>

<template>
  <main>
    <section>
      <ul>
        <li v-for="{ title, href } in LINKS" :key="title">
          <a :href="href" target="_blank" rel="noopener noreferrer">
            <div>
              <span>{{ title }}</span>
            </div>
          </a>
        </li>
      </ul>
    </section>
  </main>
</template>