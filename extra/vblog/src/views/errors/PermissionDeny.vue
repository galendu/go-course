<template>
  <div class="content">
    <a-result
      class="result"
      status="403"
      :subtitle="'你无权限访问该页面, 请登陆后重试'"
    >
      <template #extra>
        <a-button key="back" type="primary" @click="back"> 返回主页 </a-button>
      </template>
    </a-result>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";

const router = useRouter();
const back = () => {
  // warning： Go to the node that has the permission
  router.push({ name: "home" });
};
</script>

<style scoped lang="less">
.content {
  position: absolute;
  top: 50%;
  left: 50%;
  margin-left: -95px;
  margin-top: -121px;
  text-align: center;
}
</style>
