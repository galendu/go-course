<template>
  <div>
    <h2 class="title">
      {{ title }}
    </h2>
    <div class="describe">
      <div>发表于: {{ publishAt }}</div>
      <a-divider direction="vertical" />
      <div>作者: {{ author }}</div>
    </div>
    <div style="width: 700px">
      <md-editor
        v-model="text"
        preview-only
        style="height: calc(100vh - 85px)"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import MdEditor from "md-editor-v3";
import "md-editor-v3/lib/style.css";

const title = "测试";
const publishAt = 10000;
const author = "测试";
const text = ref(`## 参考

+ [博客样式参考](http://www.arccode.net/)
+ [md-editor-v3 官方页面](https://imzbf.github.io/md-editor-v3/index)
+ [md-editor-v3 Github](https://github.com/imzbf/md-editor-v3)`);
</script>

<style lang="less" scoped>
.title {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  color: #555;
}

.describe {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  color: #999;
  font-size: 12px;
}
</style>
