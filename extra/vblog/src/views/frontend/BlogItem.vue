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
    <div class="summary">
      {{ summary }}
    </div>
    <div class="more">
      <a-button size="mini" type="dashed" @click="jumpToDetail"
        >阅读全文 »</a-button
      >
    </div>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";

defineProps({
  title: String,
  publishAt: Number,
  author: String,
  summary: String,
});

const router = useRouter();
const jumpToDetail = () => {
  router.push({ path: `frontend/blogs/xxx` });
};
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

.summary {
  color: #333;
  font-size: 14px;
  line-height: 1.6;
  margin: 10px 0 10px 0;
}

.more {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
}
</style>
