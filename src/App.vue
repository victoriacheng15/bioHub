<script setup lang="ts">
import { css } from "../styled-system/css"
import { Header, LinkContainer, Footer } from "./components"
import { NAME, HEADLINE } from "./constant"


</script>

<template>
  <div :class="css({
    h: '100vh',
    maxW: '2xl',
    mx: 'auto',
    pt: '10',
    px: { base: '6', md: '10' },
    display: 'flex',
    flexDir: 'column',
    alignItems: 'center'
  })">
    <Header :name="NAME" :headline="HEADLINE" />
    <LinkContainer />
    <Footer :name="NAME" />
  </div>
</template>
