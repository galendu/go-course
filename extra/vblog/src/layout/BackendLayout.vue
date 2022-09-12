<script setup>
import { RouterView, useRouter } from "vue-router";

const router = useRouter();

const clickMenu = (key) => {
  router.push(key);
};

const jumpToFrontend = () => {
  router.push("/frontend");
};
</script>

<template>
  <div>
    <div class="header">
      <div class="logo">我的博客</div>
      <div class="right-header">
        <div>
          <a-button @click="jumpToFrontend" size="mini" type="text"
            >前台</a-button
          >
        </div>
      </div>
    </div>
    <div class="main">
      <div class="sidebar">
        <a-menu
          :style="{ width: '200px', height: '100%' }"
          :default-open-keys="['0']"
          :default-selected-keys="['0_0']"
          show-collapse-button
          breakpoint="xl"
          @menu-item-click="clickMenu"
        >
          <a-sub-menu key="0">
            <template #icon><icon-apps></icon-apps></template>
            <template #title>文章管理</template>
            <a-menu-item key="/backend/blogs">文章列表</a-menu-item>
            <a-menu-item key="/backend/tags">标签管理</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </div>
      <div class="content">
        <RouterView />
      </div>
    </div>
  </div>
</template>

<style scoped>
.header {
  display: flex;
  align-content: center;
  justify-content: flex-start;
  align-items: center;
  border-bottom: 1px solid rgb(229, 230, 235);
  height: 45px;
}

.logo {
  margin-left: 8px;
  font-size: 14px;
  font-weight: 500;
}

.right-header {
  margin-left: auto;
}

.main {
  display: flex;
  align-content: center;
  justify-content: flex-start;
  align-items: flex-start;
  height: calc(100vh - 45px);
}

.sidebar {
  height: 100%;
  border-right: 1px solid rgb(229, 230, 235);
}

.content {
  margin: 8px;
}
</style>
