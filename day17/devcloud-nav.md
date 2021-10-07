# 导航页面

我们登录后的主页面 就是一张白纸, 显然这不是我们想要的, 现在就来完善我们的主页面, 大致效果如下:

![](./images/nav-page.jpg)

## Layout布局

我们整体采用上-左-右的布局, 具体构成:
+ 最上边: 顶部导航
+ 左边: 侧边栏导航
+ 右边: 内容页面

按照这个划分, 我们把这些部分独立成一个一个的Layout组件, 创建目录:layout/components

+ Navbar.vue
+ Sidebar.vue
+ AppMain.vue

先写框架, 留空页面

Navbar.vue 组件:
```html
<template>
  <div class="navbar">
      navbar
  </div>
</template>

<script>
export default {
  name: 'Navbar',
}
</script>
```

Sidebar.vue 组件:
```html
<template>
  <div class="sidebar">
      sidebar
  </div>
</template>

<script>
export default {
  name: 'Sidebar',
}
</script>
```


AppMain.vue 组件:
```html
<template>
  <section class="app-main">
    <transition name="fade-transform" mode="out-in">
      <router-view :key="key" />
    </transition>
  </section>
</template>

<script>
export default {
  name: 'AppMain',
  computed: {
    key() {
      return this.$route.path
    }
  }
}
</script>
```

为了方便使用components内的组件, 我们创建一个index.js 把这些组件导出来: layout/components/index.js
```js
export { default as AppMain } from './AppMain'
export { default as Navbar } from './Navbar'
export { default as Sidebar } from './Sidebar'
```

最后我们把这些组件组合起来, 就是我们的Layout组件了: layout/index.vue

```html
<template>
  <div>
    <!-- 顶部导航栏 -->
    <div class="navbar-container">
        <navbar />
    </div>

    <!-- 主内容区 -->
    <div class="app-wrapper">
        <!-- 侧边栏导航 -->
        <div class="sidebar-container">
            <sidebar />
        </div>
        <!-- 内容页面区 -->
        <div class="main-container">
            <app-main />
        </div>
    </div>
  </div>
</template>

<script>
import { AppMain, Navbar, Sidebar } from './components'

export default {
  name: 'Layout',
  components: {
    AppMain,
    Navbar,
    Sidebar
  },
}
</script>
```

最后修改我们的Home路由, 采用Layout布局: router/index.js
```js
{
  path: '/',
  component: Layout,
  redirect: '/dashboard',
  children: [
    {
      path: 'dashboard',
      component: () => import('@/views/dashboard/index'),
      name: 'Dashboard',
    }
  ]
},
```

最终我们的Home页面就是这样的:
![](./images/nav-fw.jpg)

> 试着修改Home页面的内容, 看看页面是否正常显示

由于现在没有任何样式, 显得很Low B, 接下里就为其填充样式

## Layout样式

由于layout是全局样式, 我们新增一个全局样式css文件: styles/layout.scss
```scss
#app {}
```

然后通过styles下index.js 导入, 加载到全局
```js
@import './element-ui.scss';
@import './layout.scss';

...
```

然后我们来为这3个组件补充基础样式
```scss
#app {
    .navbar-container {
        width:100vw;
        height:60px;
        background-color: #444681;
    }
}
```