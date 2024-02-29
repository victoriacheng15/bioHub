<script setup lang="ts">
import { css } from "../../styled-system/css"
import { LINKS } from "../constant";
import LinkElement from "./LinkElement.vue";

const styles = css({
  w: '100%',
  h: '100%',
  display: 'grid',
  placeItems: 'center'
})

</script>

<template>
  <main :class="css({ w: '100%', mb: '14' })">
    <ul :class="css({ display: 'grid', gap: { base: '4', sm: '6', md: '8' } })">
      <li v-for="{ title, href } in LINKS" :key="title" :class="styles">
        <LinkElement :href="href" :title="title" />
      </li>
    </ul>
  </main>
</template>