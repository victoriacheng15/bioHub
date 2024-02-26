<script setup lang="ts">
import { css } from "../styled-system/css"
import { Header, LinkContainer, Footer } from "./components"
import { NAME, HEADLINE } from "./constant"


</script>

<template>
  <div :class="css({
    height: '100vh',
    maxW: '2xl',
    marginInline: 'auto',
    paddingTop: '12',
    paddingInline: '8',
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
