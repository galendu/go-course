<template>
  <div class="about">
    <!-- 这里为啥没有使用 name.value来访问喃?
    当 ref 在模板中作为顶层 property 被访问时，它们会被自动“解包”，所以不需要使用 .value -->
    <h2>{{ name }}</h2>
    <input v-model="name" type="text" />
  </div>
</template>

<script setup>
// 以库的形式来使用vue实例提供的API
import { ref } from "vue";

// 使用ref来为基础类型 构造响应式变量
const name = ref("老喻");

// 通过value来设置 基础类型的值(Setter方式)
name.value = "张三";
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
