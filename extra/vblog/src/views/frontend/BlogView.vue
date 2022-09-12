<script setup>
import BlogItem from "./BlogItem.vue";
</script>

<template>
  <div class="content">
    <a-space direction="vertical" style="width: 700px">
      <BlogItem
        title="测试"
        publishAt="1662978665"
        author="测试"
        summary="Java启动命令是所有java应用程序的入口，通过它来启动Java运行时环境，并加载相关的class。不过由于IDE的盛行，我们Java程序员中的大多数并不是非常的了解Java启动命令。本文希望做一个Java启动命令的汇总，和各位同道分享，也便于日后作为自己的参考"
      ></BlogItem>
      <BlogItem
        title="测试"
        publishAt="1662978665"
        author="测试"
        summary="测试"
      ></BlogItem>
      <BlogItem
        title="测试"
        publishAt="1662978665"
        author="测试"
        summary="测试"
      ></BlogItem>
    </a-space>
  </div>
</template>

<style lang="less" scoped>
.content {
  width: 100%;
}
</style>
