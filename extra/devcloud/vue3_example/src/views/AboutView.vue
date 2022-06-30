<template>
  <div class="about">
    <h2>{{ name }}</h2>
    <input v-model="name" type="text" />
  </div>
</template>

<script setup>
// 以库的形式来使用vue实例提供的API
import {
  ref,
  onBeforeMount,
  onMounted,
  onBeforeUpdate,
  onUpdated,
  onBeforeUnmount,
  onUnmounted,
} from "vue";

// 响应式状态
// 相当于 data里面的 name属性
const name = ref("老喻");

onBeforeMount(() => {
  console.log("before mount");
});
onMounted(() => {
  console.log("mounted");
});
onBeforeUpdate(() => {
  console.log("before update");
});
onUpdated(() => {
  console.log("on updated");
});
onBeforeUnmount(() => {
  console.log("before unmount");
});
onUnmounted(() => {
  console.log("unmounted");
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
