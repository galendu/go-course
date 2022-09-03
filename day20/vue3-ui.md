# UI组件调研

我调研市面上对vue3支持的ui插件:
+ [Element Plus](https://element-plus.org/zh-CN/guide/design.html): Element开源UI库
+ [Ant Design Vue](https://www.antdv.com/docs/vue/introduce-cn): 阿里开源UI库
+ [Vuetify](https://vuetifyjs.com/zh-Hans/): Material 样式的 Vue UI 组件库
+ [TDesign](https://tdesign.tencent.com/vue-next/overview): 腾讯开源UI库
+ [Arco Design](https://arco.design/): 字节跳动出品的企业级设计系统

## Element Plus

通过插件的方式安装UI组件: plugins/elementPlus.ts
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

## Arco Design

1. 安装UI库
```sh
# npm
npm install --save-dev @arco-design/web-vue
# yarn
yarn add --dev @arco-design/web-vue
```

2. vue加载UI库
修改: nuxt.config.ts, 依赖compute-scroll-into-view，需要使用Babel处理下
```ts
// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
    build: {
        transpile: ['compute-scroll-into-view'],
    },
})
```

通过插件的方式安装UI组件: plugins/arcoDesign.ts
```ts
// 引入组件库
import ArcoVue from "@arco-design/web-vue";
// Arco图标是一个独立的库，需要额外引入并注册使用
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
// 加载样式
import "@arco-design/web-vue/dist/arco.css";

export default defineNuxtPlugin(nuxtApp => {
  // Doing something with nuxtApp
  nuxtApp.vueApp.use(ArcoVue)
  nuxtApp.vueApp.use(ArcoVueIcon)
})
```

3. 引入一个Menu组件进行测试, 修改pages/app.vue:
```vue
<template>
  <div class="menu-demo">
    <a-menu
      :style="{ width: '200px', height: '100%' }"
      :default-open-keys="['0']"
      :default-selected-keys="['0_2']"
      show-collapse-button
      breakpoint="xl"
      @collapse="onCollapse"
    >
      <a-sub-menu key="0">
        <template #icon><icon-apps></icon-apps></template>
        <template #title>Navigation 1</template>
        <a-menu-item key="0_0">Menu 1</a-menu-item>
        <a-menu-item key="0_1">Menu 2</a-menu-item>
        <a-menu-item key="0_2">Menu 3</a-menu-item>
        <a-menu-item key="0_3">Menu 4</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="1">
        <template #icon><icon-bug></icon-bug></template>
        <template #title>Navigation 2</template>
        <a-menu-item key="1_0">Menu 1</a-menu-item>
        <a-menu-item key="1_1">Menu 2</a-menu-item>
        <a-menu-item key="1_2">Menu 3</a-menu-item>
      </a-sub-menu>
      <a-sub-menu key="2">
        <template #icon><icon-bulb></icon-bulb></template>
        <template #title>Navigation 3</template>
        <a-menu-item key="2_0">Menu 1</a-menu-item>
        <a-menu-item key="2_1">Menu 2</a-menu-item>
        <a-sub-menu key="2_2" title="Navigation 4">
          <a-menu-item key="2_2_0">Menu 1</a-menu-item>
          <a-menu-item key="2_2_1">Menu 2</a-menu-item>
        </a-sub-menu>
      </a-sub-menu>
    </a-menu>
  </div>
</template>
<script lang="ts" setup>
import { Message } from '@arco-design/web-vue';

const onCollapse = (val: String, type: String) => {
  const content = type === 'responsive' ? '触发响应式收缩' : '点击触发收缩';
  Message.info({
    content,
    duration: 2000,
  });
}
</script>
<style scoped>
.menu-demo {
  box-sizing: border-box;
  width: 100%;
  height: 600px;
  padding: 40px;
  background-color: var(--color-neutral-2);
}
</style>
```
