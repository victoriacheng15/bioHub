<script setup lang="ts">
import { css } from "../../styled-system/css"
import { LINKS } from "../constant";
import LinkElement from "./LinkElement.vue";
</script>

<template>
  <main :class="css({ w: '100%', mb: '14' })">
    <ul :class="css({ display: 'grid', gap: { base: '4', sm: '6', md: '8' } })">
      <li v-for="{ title, href } in LINKS" :key="title">
        <LinkElement :href="href" :title="title" />
      </li>
    </ul>
  </main>
</template>

<style scoped>
li {
  height: 60px;
  background: #111827;
  position: relative;
  transform: rotate(-1deg);
}

li:hover:before {
  content: "";
  position: absolute;
  top: 0;
  left: -5px;
  border-top: 30px solid transparent;
  border-left: 40px solid rgb(31, 41, 55);
  border-bottom: 30px solid transparent;
}

li:hover:after {
  content: "";
  position: absolute;
  top: 0;
  right: -5px;
  border-top: 30px solid transparent;
  border-right: 40px solid rgb(31, 41, 55);
  border-bottom: 30px solid transparent;

}
</style>