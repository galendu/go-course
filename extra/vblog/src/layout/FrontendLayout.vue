<script setup>
import { RouterView } from "vue-router";
</script>

<template>
  <div>
    <!-- 顶部导航 -->
    <div class="header">
      <div class="logo">我的博客</div>
      <div class="right-header">
        <div>
          <!-- 登录后台进行博客管理 -->
          <a-button size="mini" type="text">登录</a-button>
        </div>
      </div>
    </div>
    <!-- 显示博客列表 -->
    <div class="content">
      <RouterView />
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

.content {
  margin: 20px;
  display: flex;
  align-content: center;
  justify-content: center;
  align-items: center;
}
</style>
