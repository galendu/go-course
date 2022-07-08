<script setup>
import { inject } from "vue";

const count = inject("count");

const doClick = () => {
  count.value++;
};
</script>

<template>
  <button @click="doClick">You clicked me {{ count }} times.</button>
</template>

<style scoped></style>
