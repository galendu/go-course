# Vue入门基础

学习Vue需要提前掌握Web基础:
+ HTML
+ CSS
+ Javascript

## 安装

我们需要安装的清单如下:
+ node
+ vue
+ vue/cli
+ Vue Devtools, 
+ Vuter

### 安装node

NodeJS[下载地址](https://nodejs.org/zh-cn/), 安装最新的版本

下面是我当前环境:
```sh
$ node -v
v14.17.1
$ npm -v
6.14.13
```

配置使用国内源

我们可以通过 `npm config get registry` 查看当前配置的源是那个
```sh
$ npm config get registry
https://registry.npmjs.org/
```

npm官方源下载依赖的速度很慢, 所以往往我们都需要更新成国内的源, 下面是使用淘宝的源

```sh
$ npm config set registry http://registry.npm.taobao.org/
$ npm config get registry
http://registry.npm.taobao.org/
```

到处我们就更换成了国内源的地址

### 安装Vue和脚手架工具

直接使用node进行全局安装

```sh
# 最新稳定版
$ npm install -g vue
```

官方同时提供了1个cli工具, 用于快速初始化一个vue工程, 官方文档: [vue-cli](https://cli.vuejs.org/)

安装项目脚手架工具:
```
sudo npm install -g @vue/cli
```

本教程版本要求:
+ vue: 2.6.x
+ vue-cli: 4.5.x

安装完成后使用 npm list -g 查看当前node环境下 该模块是否安装成功, 版本是否正确
```sh
$ npm list -g | grep vue
├─┬ @vue/cli@4.5.13 ## vue cli安装成功, 版本4.5.13
│ ├─┬ @vue/cli-shared-utils@4.5.13
│ ├─┬ @vue/cli-ui@4.5.13
│ │ ├── @vue/cli-shared-utils@4.5.13 deduped
│ ├── @vue/cli-ui-addon-webpack@4.5.13
│ ├── @vue/cli-ui-addon-widgets@4.5.13
│ ├── vue@2.6.14  ## vue 安装成功, 版本2.6.14
│ ├─┬ vue-codemod@0.0.5
│ │ ├─┬ @vue/compiler-core@3.2.6
│ │ │ ├── @vue/shared@3.2.6
│ │ ├─┬ @vue/compiler-dom@3.2.6
│ │ │ ├── @vue/compiler-core@3.2.6 deduped
│ │ │ └── @vue/shared@3.2.6 deduped
```

### Vue Devtools

在使用 Vue 时，我们推荐在你的浏览器上安装 Vue Devtools。它允许你在一个更友好的界面中审查和调试 Vue 应用

#### chrome商店安装

vue-devtools可以从chrome商店直接下载安装，非常简单， 具体请参考: [Chrome Setting](https://devtools.vuejs.org/guide/installation.html#settings) 这里就不过多介绍了。不过要注意的一点就是，需要翻墙才能下载


#### 离线安装


### vscode 插件

+ Beautify: js, css, html 语法高亮差距
+ ESLint: js eslint语法风格检查
+ Auto Rename Tag: tag rename
+ Veter: vue语法高亮插架


## 参考


+ [VUE2官方文档](https://cn.vuejs.org/v2/guide)