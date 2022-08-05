<template>
  <div class="about">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { Greet } from "../../wailsjs/go/main/App";

const msg = ref("hello");

onMounted(async () => {
  msg.value = await Greet("hello");
});
</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
