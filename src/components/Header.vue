<script setup lang="ts">
import avatar from "/avatar.jpg"
const { name, headline } = defineProps<{ name: string; headline: string }>();
const imageInfo = {
  src: avatar,
  alt: name,
  width: 200,
  height: 200,
}
</script>

<template>
  <header>
    <div>
      <img :="imageInfo" />
      <h1>{{ name }}</h1>
      <p>{{ headline }}</p>
    </div>
  </header>
</template>