<template>
  <div class="login-form">
    <a-form
      ref="loginForm"
      :model="form"
      :style="{ width: '400px', height: '100%', justifyContent: 'center' }"
      @submit="handleSubmit"
    >
      <a-form-item>
        <div class="title">登录博客管理后台</div>
      </a-form-item>
      <a-form-item
        field="username"
        label=""
        :rules="[{ required: true, message: '请输入用户名' }]"
        hide-asterisk
      >
        <a-input v-model="form.username" placeholder="请输入用户名">
          <template #prefix>
            <icon-user />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item
        field="password"
        label=""
        :rules="[{ required: true, message: '请输入密码' }]"
        hide-asterisk
      >
        <a-input
          type="password"
          v-model="form.password"
          placeholder="请输入密码"
        >
          <template #prefix>
            <icon-lock />
          </template>
        </a-input>
      </a-form-item>
      <a-form-item>
        <a-button style="width: 100%" type="primary" html-type="submit"
          >登录</a-button
        >
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { reactive } from "vue";
import { Message } from "@arco-design/web-vue";
import { useRouter } from "vue-router";

const router = useRouter();

const form = reactive({
  username: "",
  password: "",
});
const handleSubmit = (data) => {
  if (data.errors === undefined) {
    let form = data.values;
    if (form.username === "admin" && form.password === "123456") {
      // 保存登录状态
      localStorage.setItem("username", form.username);
      localStorage.setItem("password", form.password);

      // 登录成功直接跳转到后台页
      const { redirect, ...othersQuery } = router.currentRoute.value.query;
      router.push({
        name: redirect || "BlogList",
        query: {
          ...othersQuery,
        },
      });
    } else {
      Message.error("用户名或者密码不正确");
    }
  }
};
</script>

<style scoped>
.login-form {
  height: 100%;
  display: flex;
  align-content: center;
  justify-content: center;
  align-items: center;
}

.title {
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: center;
  width: 100%;
  font-weight: 500;
}
</style>
