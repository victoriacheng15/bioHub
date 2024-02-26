<script setup lang="ts">
import { css } from "../styled-system/css"
import { Header, LinkContainer, Footer } from "./components"
import { NAME, HEADLINE } from "./constant"


</script>

<template>
  <div :class="css({
    h: '100vh',
    maxW: '3xl',
    mx: 'auto',
    pt: '12',
    px: '8',
    display: 'flex',
    flexDir: 'column',
    justifyContent: 'center',
    alignItems: 'center'
  })">
    <Header :name="NAME" :headline="HEADLINE" />
    <LinkContainer />
    <Footer :name="NAME" />
  </div>
</template>
