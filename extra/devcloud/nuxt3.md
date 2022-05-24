# NuxtJs工程

![](./images/nuxt3.png)


## 项目初始化


### 初始化工程
使用nuxi 初始化工程: devcloud
```sh
> npx nuxi init devcloud
Nuxt CLI v3.0.0-rc.3 
ℹ cloned nuxt/starter#v3 to /Users/yumaojun/Workspace/Nodejs/devcloud
 ✨ Your legendary Nuxt project is just created! Next steps:
 📁  cd devcloud
 💿  Install dependencies with npm install or yarn install or pnpm install --shamefully-hoist
 🚀  Start development server with npm run dev or yarn dev or pnpm run dev 
```

### 下载工程依赖
```sh
> yarn install
yarn install v1.22.18
info No lockfile found.
[1/4] 🔍  Resolving packages...
warning nuxt > nitropack > @vercel/nft > node-pre-gyp@0.13.0: Please upgrade to @mapbox/node-pre-gyp: the non-scoped node-pre-gyp package is deprecated and only the @mapbox scoped package will recieve updates in the future
[2/4] 🚚  Fetching packages...
warning vscode-languageclient@7.0.0: The engine "vscode" appears to be invalid.
[3/4] 🔗  Linking dependencies...
[4/4] 🔨  Building fresh packages...
success Saved lockfile.
✨  Done in 20.06s.
```

### 解决warning问题

1. 解决node-pre-gyp版本过低问题
```sh
> yarn upgrade @mapbox/node-pre-gyp
# 可以看到gyp的版本已经升级上去了
> yarn list | grep gyp
├─ @mapbox/node-pre-gyp@1.0.9
│  ├─ @mapbox/node-pre-gyp@^1.0.5
│  ├─ node-gyp-build@^4.2.2
│  ├─ node-pre-gyp@^0.13.0
├─ node-gyp-build@4.4.0
├─ node-pre-gyp@0.13.0
```

第二个问题等待nuxtjs官方升级, 展示对项目没影响

### 启动工程 

```sh
> yarn dev -o
```

启动完成后我们会看到这样一个页面:

![](./images/start-up.png)

接下来了解Nuxt这个脚手架，并编写Vue代码页面

## NuxtJs工程介绍

Nuxt的工程结构如下:

![](./images/directory-structure.png)

### 入口文件

首先我们需要找到工程的入口文件:
```
The app.vue file is the main component in your Nuxt 3 applications.
```

修改app.vue文件
```vue
<template>
  <div>
    <!-- <NuxtWelcome /> -->
    <h1>Hello Nuxt3!</h1>
  </div>
</template>
```

### 页面与路由

我们不可能把所有的页面逻辑都写在入口文件里面, 因此Nuxt为我们准备了一个pages目录, 放在该目录下的vue文件, nuxt会根据文件路径自动为我们创建路由映射, 比如:
```
pages/index.vue --->   /
pages/detail.vue --->  /detail
```

+ pages/index.vue
```vue
<template>
    <div>
        <h1>Index page</h1>
    </div>
</template>
```

+ pages/detail.vue
```vue
<template>
    <div>
        <h1>Detail Page</h1>
    </div>
</template>
```

+ app.vue
```vue
<template>
  <div>
    <h1>hello, nuxt3</h1>
    <!-- 添加页面路由出口 -->
     <NuxtPage />
  </div>
</template>
```

然后我们切换页面访问路径: / --> /detail 也试图就会改变

#### 动态路由

同一个页面 可能由于访问的用户不同展示出来的页面内容的数据也所有差异, 如果解决这个问题喃? 有如下2种思路:
+ 固定路径 + 路径参数, 比如 /detail?id=xxx
+ 动态路由参数, 比如 /detail/xxx, 

1. 固定路由

修改pages/detail.vue页面:
```
<template>
    <div>
        <h1>Detail Page</h1>
        <!-- $route保存了当前路由信息 -->
        <p>{{ $route }}</p>
    </div>
</template>
```

访问页面: /detail?id=xxx, 就能看到当前路由页面的路由信息
```json
{
    "fullPath":"/detail?id=xxx",
    "hash":"",
    "query":{"id":"xxx"},
    "name":"detail",
    "path":"/detail",
    "params":{},
    "matched":[ ... ],
    "meta":{},
    "href":"/detail?id=xxx"
}
```

那我们在编程就可以根据id向后端请求不同的数据:
```js
getDataById($route.query.id)
```

2. 动态路由

为了避免之前路径的影响，先删除之前的detail.vue页面，然后创建一个pages/detail/[id].vue的页面, 这里使用[id], 就是路径参数变量的表示
```vue
<template>
    <div>
        <h1>Detail Page</h1>
        <!-- $route保存了当前路由信息, 通过params获取路径参数的所有变量 -->
        <p>{{ $route.params }}</p>
    </div>
</template>
```

访问页面: /detail/xxx, 就能看到当前路由页面的路由信息

#### 自定义404页面

Nuxt3默认有404页面，如果想要自定义404页面, 只需要添加pages/404.vue页面,比如: 
```vue
<template>
    <div>
        <h1>custom 404</h1>
    </div>
</template>
```

再次访问就能访问到看看404页面是否生效

#### 路由嵌套

很多场景下 我们一个页面可能很复杂，有很多个组件构成, 我们可以选择以组件的方式构建页面，比如:
```vue
<template>
    <div>
        <h1>dashboard main page</h1>
        <component1 />
        <component2 />
        <component3 />
    </div>
</template>
```

我们也可以选择路由嵌套来实现, 比如下面就是一个dashboard嵌套页面的目录:
```sh
# 注意: dashboard.vue 如果是是嵌套, 则需要创建一个名称为:dashboard的目录, 里面存放具体嵌套的页面
-| pages/
---| dashboard/
------| index.vue
------| component1.vue
------| component2.vue
------| component3.vue
---| dashboard.vue
```

上面目录结构 映射为 下面这样的 嵌套路由:
```json
[
  {
    path: '/dashboard',
    component: '~/pages/dashboard.vue',
    name: 'dashboard',
    children: [
      {
        path: '/',
        component: '~/pages/dashboard/index.vue',
        name: 'dashboard-index'
      },
      {
        path: '/component1',
        component: '~/pages/dashboard/component1.vue',
        name: 'dashboard-component1'
      },
      {
        path: '/component2',
        component: '~/pages/dashboard/component2.vue',
        name: 'dashboard-component2'
      },
      {
        path: '/component3',
        component: '~/pages/dashboard/component3.vue',
        name: 'dashboard-component3'
      }
    ]
  }
]
```

这里需要注意的是, 页面嵌套时, 主页面需要添加 NuxtPage 才能路由的字页面, 比如pages/dashboard.vue
```vue
<template>
    <div>
        <h1>dashboard main page</h1>
        <!-- 子页面出口 -->
        <NuxtPage />
    </div>
</template>
```


思考: 你想把当前的页面(日程)通过url分享给另外一个人, 如何确保分享点开页面是显示的是日程这个标签页

![](./images/tabs.png)

#### 路由跳转


#### 页面元数据


#### 编程式路由



#### Nuxt与Vue Router



### 安装UI组件


#### UI组调研

+ [Element Plus](https://element-plus.org/zh-CN/guide/design.html): Element开源UI库
+ [Ant Design Vue](https://www.antdv.com/docs/vue/introduce-cn): 阿里开源UI库
+ [Vuetify](https://vuetifyjs.com/zh-Hans/): Material 样式的 Vue UI 组件库
+ [TDesign](https://tdesign.tencent.com/vue-next/overview): 腾讯开源UI库
+ [Arco Design](https://arco.design/): 字节跳动出品的企业级设计系统

#### 安装Element Plus

通过插件的方式安装UI组件: plugins/element-plus.ts
```ts
import ElementPlus from 'element-plus'

export default defineNuxtPlugin(nuxtApp => {
    nuxtApp.vueApp.use(ElementPlus)
})
```

修改Nuxt配置, 添加全局样式表

nuxt.config.ts
```ts
import { defineNuxtConfig } from 'nuxt'

// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    // css
    css: ['~/assets/style/index.css'],
})
```

#### 安装Arco Design

1. 安装UI库
```sh
# npm
npm install --save-dev @arco-design/web-vue
# yarn
yarn add --dev @arco-design/web-vue
```

2. vue加载UI库
```vue
yarn add -D less
yarn add -D less-loader
yarn add -D unplugin-auto-import
yarn add -D unplugin-icons
yarn add -D unplugin-vue-components
```


### 页面布局



## 参考

+ [vue3官方文档](https://vuejs.org/guide/introduction.html)
+ [nuxtjs官网](https://v3.nuxtjs.org/getting-started/quick-start)
+ [nuxt项目启动时跳过Are you interested in participation](http://www.flydream.cc/article/nuxt-bootstrap-skip-participation/)
+ [element-plus-nuxt-starter](https://github.com/element-plus/element-plus-nuxt-starter)
+ [Arco Design Plans for Nuxt 3 support?](https://github.com/arco-design/arco-design-vue/issues/24)