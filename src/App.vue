<script setup lang="ts">
import { Header, LinkContainer, Footer } from "./components"
import { NAME, HEADLINE } from "./constant"
</script>

<template>
  <Header :name="NAME" :headline="HEADLINE" />
  <LinkContainer />
  <Footer :name="NAME" />
</template>
