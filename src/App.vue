<script setup lang="ts">
import Header from "./components/Header.vue";
import LinkContainer from "./components/LinkContainer.vue";
import Footer from "./components/Footer.vue";
import { NAME, HEADLINE } from "./constant";
</script>

<template>
  <Header :name="NAME" :headline="HEADLINE" />
  <LinkContainer />
  <Footer :name="NAME" />
</template>
