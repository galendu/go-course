<template>
  <div class="login-container">
      <el-form class="login-form">
        <!-- 切换 -->
        <div>
            <el-tabs>
            <el-tab-pane label="普通登陆">

            </el-tab-pane>
            <el-tab-pane label="LDAP登陆">

            </el-tab-pane>
            </el-tabs>
        </div>

        <!-- 账号输入框 -->
        <el-form-item>
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <el-input>

        </el-input>
        </el-form-item>

        <!-- 密码输入框 -->
        <el-form-item>
        <span class="svg-container">
          <svg-icon icon-class="password" />
        </span>
        <el-input>

        </el-input>
        </el-form-item>

        <!-- 登陆按钮 -->
        <el-button class="login-btn">
            登陆
        </el-button>
      </el-form>
  </div>
</template>

<script>
export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        grant_type:'',
        username: '',
        password: ''
      },
    }
  }
}
</script>

<style lang="scss" scoped>
.login-container {
  height: 100%;
  width: 100%;
  background-image: linear-gradient(to top, #3584A7 0%, #473B7B 100%);
  .login-form {
    width: 520px;
    padding: 160px 35px 0;
    margin: 0 auto;
    .login-btn {
        width:100%;
    }
  }
  .svg-container {
    padding: 6px 5px 6px 15px;
    color: #889aa4;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }
}

/* reset element-ui css */
.login-container ::v-deep .el-input {
    display: inline-block;
    height: 47px;
    width: 85%;
    input {
      background: transparent;
      border: 0px;
      -webkit-appearance: none;
      border-radius: 0px;
      padding: 12px 5px 12px 15px;
      height: 47px;
      caret-color: #fff;
    }
  }

.login-container ::v-deep .el-form-item {
    border: 1px solid rgba(255, 255, 255, 0.1);
    color: #454545;
}

.login-container ::v-deep .el-tabs__item {
  color: white;
  font-size: 18px;
}

.login-container ::v-deep .is-active {
  color:#13C2C2;
}

</style>