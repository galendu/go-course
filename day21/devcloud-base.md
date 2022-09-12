# 前端框架搭建

## 项目初始化

### 生成项目模版

这里我们直接使用vue cli 初始化我们的骨架, 然后在此基础上进行修改
```sh
npm init vue@latest
Need to install the following packages:
  create-vue@3.3.4
Ok to proceed? (y) y

Vue.js - The Progressive JavaScript Framework

✔ Project name: … vblog
✔ Add TypeScript? … No / Yes
✔ Add JSX Support? … No / Yes
✔ Add Vue Router for Single Page Application development? … No / Yes
✔ Add Pinia for state management? … No / Yes
✔ Add Vitest for Unit Testing? … No / Yes
✔ Add Cypress for both Unit and End-to-End testing? … No / Yes
✔ Add ESLint for code quality? … No / Yes
✔ Add Prettier for code formatting? … No / Yes

Scaffolding project in /Users/yumaojun/Workspace/Golang/go-course/extra/vblog...

Done. Now run:

  cd vblog
  npm install
  npm run lint
  npm run dev
```

### 设置项目源

检查源设置:
```sh
yrm ls

  npm ---- https://registry.npmjs.org/
  cnpm --- http://r.cnpmjs.org/
* taobao - https://registry.npm.taobao.org/
  nj ----- https://registry.nodejitsu.com/
  rednpm - http://registry.mirror.cqupt.edu.cn/
  npmMirror  https://skimdb.npmjs.com/registry/
  edunpm - http://registry.enpmjs.org/
  yarn --- https://registry.yarnpkg.com
```

如果不是，使用yrm use来进行切换
```sh
# 使用淘宝的源
> yrm use taobao
   YARN Registry has been set to: https://registry.npm.taobao.org/
   NPM Registry has been set to: https://registry.npm.taobao.org/
```

### 清理模版页面

在做Home页面之前，先清理掉脚手架为我们生成的页面

1. 清理App.vue, 只保留router试图部分, 其他部分删除掉
```vue
<script setup>
import { RouterView } from "vue-router";
</script>

<template>
  <RouterView />
</template>

<style scoped>
</style>
```

2. 只保留了HomeView, 内容留白:
```vue
<script setup></script>

<template>
  <main>Home</main>
</template>
```

3. 删除其他页面和无用的组件:
+ AboutView 删除
+ components 目录下的所有组件

4. 清理样式

base.css设置全局样式
```css
* {
  box-sizing: border-box;
}

html,
body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-size: 14px;
  background-color: var(--color-bg-1);
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
}

#app {
  width: 100%;
  height: 100%;
}
```

main.css 只保留基础样式
```css
@import "./base.css";
```

## 安装UI插件

接下来就需要编写我们的业务页面, 这我们选择的UI组件为: Arco Design


### 安装 Arco Design

```sh
# npm
npm install --save-dev @arco-design/web-vue
# yarn
yarn add --dev @arco-design/web-vue
```

在vue安装该插件
```js
import { createApp } from 'vue'
import ArcoVue from '@arco-design/web-vue';
import App from './App.vue';
import '@arco-design/web-vue/dist/arco.css';

const app = createApp(App);
app.use(ArcoVue);
app.mount('#app');
```

### 验证插件

修改HomeView, 引入buttom UI组件进行测试
```vue
<script setup></script>

<template>
  <main>
    <a-space>
      <a-button type="primary">Primary</a-button>
      <a-button>Secondary</a-button>
      <a-button type="dashed">Dashed</a-button>
      <a-button type="outline">Outline</a-button>
      <a-button type="text">Text</a-button>
    </a-space>
  </main>
</template>
```

## 错误页面

这里我们需要补充2种异常页面:
+ 404页面: 当用户输入的URL并没有匹配页面时
+ 403页面: 当用户未登陆就访问管理页面时

我们使用[HTTP状态码 404](https://arco.design/vue/component/result)组件进行封装

### 404页面

errors/NotFound.vue:
```vue
<template>
  <div class="content">
    <a-result class="result" status="404" :subtitle="'not found'">
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
```

这里我们使用到了less, 一种css编译器(css扩展), 因此需要安装less的编译器
```sh
npm install --save-dev less
```

最后我们在路由上补充上404路由
```js
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/:pathMatch(.*)*",
      name: "notFound",
      component: () => import("@/views/errors/NotFound.vue"),
    },
  ],
});
```

### 403页面

添加页面: errors/PermissionDeny.vue
```vue
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
```

补充路由:
```js
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/errors/403",
      name: "PermissionDeny",
      component: () => import("@/views/errors/PermissionDeny.vue"),
    },
    {
      path: "/:pathMatch(.*)*",
      name: "NotFound",
      component: () => import("@/views/errors/NotFound.vue"),
    },
  ],
});
```

## Layout布局

界面分为前台和管理后台

博客前台:

![](./images/vblog_frontend.png)

博客管理后台:

![](./images/vblog_backend.png)

### 博客前台

因为前台与后台布局样式不一样, 因此分别使用独立布局模版.

#### Layout

前台布局模版: FrontendLayout.vue
```vue
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
          <a-button size="mini" type="text">后台管理</a-button>
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
```

#### blog占位页

补充一个前台展位页面: frontend/BlogView.vue
```vue
<script setup></script>

<template>
  <main>
    博客页面
  </main>
</template>
```

#### Blog路由

```js
    {
      path: "/",
      name: "home",
      redirect: "/frontend",
    },
    {
      path: "/frontend",
      name: "frontend",
      component: FrontendLayout,
      children: [
        {
          path: "",
          name: "frontend",
          component: BlogView,
        },
      ],
    },
```


### 博客后台

#### Loayout

博客后台使用的布局模版: BackendLayout.vue

这里我们需要使用到侧边栏导航:[Arco Design菜单 Menu](https://arco.design/vue/component/menu)

```vue
<script setup>
import { RouterView, useRouter } from "vue-router";

const router = useRouter();
const clickMenu = (key) => {
  router.push(key);
};
</script>

<template>
  <div>
    <div class="header">
      <div class="logo">我的博客</div>
      <div class="right-header">
        <div>
          <a-button size="mini" type="text">前台</a-button>
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
```

#### blog占位页

添加 backend/BlogList.vue

```vue
<script setup></script>

<template>
  <main>博客列表</main>
</template>
```

添加 backend/TagList.vue

```vue
<script setup></script>

<template>
  <main>标签列表</main>
</template>
```

#### blog路由

```js
    {
      path: "/backend",
      name: "backend",
      component: BackendLayout,
      children: [
        {
          path: "blogs",
          name: "BlogList",
          component: BlogList,
        },
        {
          path: "tags",
          name: "TagList",
          component: TagList,
        },
      ],
    },
```

### 切换到前台

无效任务 后台可以直接切换到前台:

layout/BackendLayout.vue
```js
const jumpToFrontend = () => {
  router.push("/frontend");
};
```

### 切换到后台

但是前台切换到后台，是需要认证的, 因此需要先做登录页面

#### 登录页面

这里我们将要使用到: [表单 Form](https://arco.design/vue/component/form)

```vue
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
      <a-form-item field="remember">
        <a-checkbox v-model="form.remember"> 记住密码 </a-checkbox>
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
  remember: false,
});
const handleSubmit = (data) => {
  if (data.errors === undefined) {
    let form = data.values;
    if (form.username === "admin" && form.password === "123456") {
      // 记住密码
      if (form.remember) {
        localStorage.setItem("username", form.username);
        localStorage.setItem("password", form.password);
      }
      // 登录成功直接跳转到后台页面
      router.push("/backend/blogs");
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
```

添加路由:
```js
    {
      path: "/login",
      name: "LoginPage",
      component: () => import("@/views/login/LoginPage.vue"),
    },
```

#### 切换到后台

前台页面添加登录跳转: fronend/FrontendLayout.vue:
```vue
<script setup>
import { RouterView, useRouter } from "vue-router";

const router = useRouter();

const jumpToBackend = () => {
  const username = localStorage.getItem("username");
  const password = localStorage.getItem("password");
  if (
    username !== null &&
    password !== null &&
    username !== "" &&
    password !== ""
  ) {
    // 直接跳转后台管理页面
    router.push("/backend/blogs");
  } else {
    // 跳转去登录页面
    router.push("/login");
  }
};
</script>
```



 





## 参考 

