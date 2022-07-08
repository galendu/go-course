<script setup>
import { ref } from "vue";

const count = ref(0);

// 其中事件的名称可以使用数组来进行定义
const emit = defineEmits(["inFocus", "submit"]);

const doClick = () => {
  count.value++;
  emit("submit", count.value);
};
</script>

<template>
  <button @click="doClick">You clicked me {{ count }} times.</button>
</template>

<style scoped></style>
