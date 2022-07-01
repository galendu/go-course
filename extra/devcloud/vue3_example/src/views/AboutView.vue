<template>
  <div class="about">
    <!-- 这里为啥没有使用 name.value来访问喃?
    当 ref 在模板中作为顶层 property 被访问时，它们会被自动“解包”，所以不需要使用 .value -->
    <h2>{{ name }}</h2>
    <input v-model="name" type="text" />
  </div>
</template>

<script setup>
// 当启用响应性语法糖时，这些宏函数都是全局可用的、无需手动导入。但如果你想让它更明显，你也可以选择从 vue/macros 中引入它们
import { $ref } from "vue/macros";

// 使用ref来为基础类型 构造响应式变量
let name = $ref("老喻");

// 通过value来设置 基础类型的值(Setter方式)
name = "张三";
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
