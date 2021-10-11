<template>
  <div class="navbar">
    <!-- logo -->
    <div class="logo-container">
      <span class="title">极乐研发云</span>
    </div>
    <!-- 主导航栏 -->
    <div class="navbar-main">
      <span>首页</span>
      <span>产品运营</span>
      <span>资源管理</span>
      <span>研发交付</span>
      <span>监控告警</span>
    </div>
    <!-- 用户信息区 -->
    <div class="navbar-user">
      <el-dropdown>
        <el-button type="text" style="color:white">
          <span>老喻</span>
          <i class="el-icon-arrow-down el-icon--right dropdown-color" />
        </el-button>
        <el-dropdown-menu slot="dropdown">
          <!-- 个人信息 -->
          <el-dropdown-item>
            <svg-icon icon-class="person" />
            <span class="dropdown-item-text">基本信息</span>
          </el-dropdown-item>
          <!-- 项目地址 -->
          <el-dropdown-item>
            <svg-icon icon-class="github" />
            <span class="dropdown-item-text">Github地址</span>
          </el-dropdown-item>
          <!-- 退出系统 -->
          <el-dropdown-item divided>
            <span style="display:block;">退出登录</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Navbar',
  data() {
    return {
      activeIndex: 'dashboard'
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  display: flex;
  align-content: center;
  justify-content: space-between;
  align-items: center;
}

.logo-container {
  box-sizing: border-box;
  width: 210px;
  padding-left: 28px;
  color: rgba(255, 255, 255, 0.8);

  .title {
    font-size: 20px;
  }
}

.navbar ::v-deep .el-menu-item {
  height: 50px;
  line-height: 50px;
}

.navbar-user {
  margin-left: auto;
  padding-right: 12px;
}

.dropdown-item-text {
  margin-left: 12px;
  margin-right: 12px;
}
</style>