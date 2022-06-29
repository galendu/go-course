<template>
  <div class="about">
    <h2>{{ name }}</h2>
    <input v-model="name" type="text" />
  </div>
</template>

<script setup>
// 以库的形式来使用vue实例提供的API
import { ref, onMounted, getCurrentInstance } from "vue";

// 响应式状态
// 相当于 data里面的 name属性
const name = ref("老喻");

// 使用构造函数onMounted
// 想到于mounted选项
onMounted(() => {
  // 通过getCurrentInstance获取到当前组件的vue实例
  const _this = getCurrentInstance();
  console.log(_this);
  console.log(_this.proxy.$data);
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
