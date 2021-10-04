# 前端框架搭建

这里我们直接使用vue cli 初始化我们的骨架, 然后在此基础上进行修改
```sh
vue create devcloud
Vue CLI v4.5.13
? Please pick a preset: Manually select features
? Check the features needed for your project: Choose Vue version, Babel, Router, Vuex, CSS Pre-processors, Linter
? Choose a version of Vue.js that you want to start the project with 2.x
? Use history mode for router? (Requires proper server setup for index fallback in production) Yes
? Pick a CSS pre-processor (PostCSS, Autoprefixer and CSS Modules are supported by default): Sass/SCSS (with node-sass)
? Pick a linter / formatter config: Prettier
? Pick additional lint features: Lint on save
? Where do you prefer placing config for Babel, ESLint, etc.? In package.json
? Save this as a preset for future projects? Yes
? Save preset as: devcloud
```
确认下我们的项目依赖是否和预期一样
```js
"core-js": "^3.6.5",
"vue": "^2.6.11",
"vue-router": "^3.2.0",
"vuex": "^3.4.0"
```

## 项目npm配置

项目中经常会遇到 有些依赖报拉取不下来, npm除了可以配置国内淘宝源，其实还有很多可以配置的选项: [The npm config files](https://www.npmjs.cn/files/npmrc/)

我们也可以通过npm config edit来查看有哪些选项可以配置
```
npm config edit
```

由于我们后面需要用到node-sass, 由于这个是个二进制的css预处理器(编译器), 而默认的url又是国外下载地址(不在npm源中), 针对这种类型的依赖，我们往往需要单独配置下载的url, 因此我们使用sass_binary_site将它配置到npm的配置文件中

本项目npm配置文件: .npmrc
```
sass_binary_site=https://npm.taobao.org/mirrors/node-sass/
registry=https://registry.npm.taobao.org
```


## 项目vue配置

vue.config.js 是一个可选的配置文件，如果项目的 (和 package.json 同级的) 根目录中存在这个文件，那么它会被 @vue/cli-service 自动加载, 关于具体的配置选项可以参考: [配置说明](https://cli.vuejs.org/zh/config/)

我们创建一个js的模块, 用于导出配置给cli使用, 这里采用commonJS模块导出语法, 关于模块相关语法还有不清楚的请查看之前课件: [JavaScript入口](../day15/javascript.md)

我们在项目根下面创建: vue.config.js
```js
// All configuration item explanations can be find in https://cli.vuejs.org/config/
module.exports = {
}
```

### 基础配置

```js
  /**
   * You will need to set publicPath if you plan to deploy your site under a sub path,
   * for example GitHub Pages. If you plan to deploy your site to https://foo.github.io/bar/,
   * then publicPath should be set to "/bar/".
   * In most cases please use '/' !!!
   * Detail: https://cli.vuejs.org/config/#publicpath
   */
  /*
    部署应用包时的基本 URL, '/'表示部署于跟路径,
    如果不放置于根, 比如: /my-app/, 访问路径就是： https://www.my-app.com/my-app/
  */
  publicPath: '/',
  /*
  当运行 vue-cli-service build 时生成的生产环境构建文件的目录
  */
  outputDir: 'dist',
  /*
  放置生成的静态资源 (js、css、img、fonts) 的 (相对于 outputDir 的) 目录
  */
  assetsDir: 'static',
  /*
  是否在开发环境下通过 eslint-loader 在每次保存时 lint 代码。这个值会在 @vue/cli-plugin-eslint 被安装之后生效,
  这里配置在开发环境生效

  注意: process.env 是当前进程的环境变量对象, 通过它可以访问到当前进程的所有环境变量
  */
  lintOnSave: process.env.NODE_ENV === 'development',
  /*
  是否需要生产环境的 source map, 可以将其设置为 false 以加速生产环境构建, 
  */
  productionSourceMap: false,
```

### DevServer

```js
    // If your port is set to 80,
    // use administrator privileges to execute the command line.
    // For example, Mac: sudo npm run
    // You can change the port by the following method:
    // port = 9527 npm run dev OR npm run dev --port = 9527
    const port = process.env.port || process.env.npm_config_port || 9527 // dev port

  /*
    所有 webpack-dev-server 的选项都支持, 具体见: https://webpack.js.org/configuration/dev-server/
  */
  devServer: {
    /* 通过环保变量获取当前开放服务器需要监听的端口*/
    port: port,
    /* 启动后浏览器默认打开的页面, 比如 open: ['/my-page', '/another-page'], true表示默认浏览器的publicPath */
    open: true,
    /* 当有编译报错时, 直接显示在页面上, 下面配置errors显示, warnings不显示 */
    overlay: {
      warnings: false,
      errors: true
    }
    /* 配置服务端代理, 开发时临时解决跨域问题 */
    // proxy: {
    //   '/workflow/api': {
    //     target: 'http://keyauth.nbtuan.vip',
    //     ws: true,
    //     secure: false,
    //     changeOrigin: true
    //   },
    // } 
    },
```


### webpack基础配置

```js
const name = defaultSettings.title || '极乐研发云' // page title
/*
webpack 相关配置
*/
configureWebpack: {
    // provide the app's title in webpack's name field, so that
    // it can be accessed in index.html to inject the correct title.
    name: name,
    // webpack 插件配置, 具体见: https://webpack.js.org/configuration/plugins/#plugins
    plugins: [],
    // 使用import的路由别名, 比如'@/components/Tips' 会别解析成: src/components/Tips
    // 更多resovle相关配置请查看: https://webpack.js.org/configuration/resolve/#resolve
    resolve: {
        alias: {
            '@': resolve('src')
        }
    }
},
```

### webpack插件配置

Vue CLI 内部的 webpack 配置是通过 webpack-chain 维护的。
这个库提供了一个 webpack 原始配置的上层抽象，使其可以定义具名的 loader 规则和具名插件，
并有机会在后期进入这些规则并对它们的选项进行修改

#### 静态资加载

静态资源的加载对页面性能起着至关重要的作用, 浏览器提供的两个资源指令-preload/prefetch，它们能够辅助浏览器优化资源加载的顺序和时机，提升页面性能
其中rel="prefetch"被称为Resource-Hints（资源提示），也就是辅助浏览器进行资源优化的指令
类似的指令还有rel="preload"

##### 预提取预prefetch

prefetch通常翻译为预提取, 其利用浏览器空闲时间来下载或预取用户在不久的将来可能访问的文档,网页向浏览器提供一组预取提示，并在浏览器完成当前页面的加载后开始静默地拉取指定的文档并将其存储在缓存中。

当用户访问其中一个预取文档时，便可以快速的从浏览器缓存中得到, 比如: 
```html
<link rel="prefetch" href="static/img/ticket_bg.a5bb7c33.png">
```

一个Vue CLI应用会为所有作为async chunk生成的JavaScript文件(通过动态import()按需code splitting的产物)自动生成prefetch提示。

这些提示会被@vue/preload-webpack-plugin注入，并且可以通过chainWebpack的config.plugin('prefetch')进行修改和删除

when there are many pages, it will cause too many meaningless requests

```js
config.plugins.delete('prefetch')
```

##### 加载(preload)

preload则翻译为预加载, 对于页面即刻需要的资源，你可能希望在页面加载的生命周期的早期阶段就开始获取，在浏览器的主渲染机制介入前就进行预加载简单来说，就是通过标签显式声明一个高优先级资源，强制浏览器提前请求资源
```html
<link rel="preload" href="xxx" as="xx">
```

一个Vue CLI应用会为所有初始化渲染需要的文件自动生成preload提示, 这些提示会被@vue/preload-webpack-plugin注入，并且可以通过chainWebpack的config.plugin('preload')进行修改和删除

it can improve the speed of the first screen, it is recommended to turn on preload

```js
config.plugin('preload').tap(() => [
{
    rel: 'preload',
    // to ignore runtime.js
    // https://github.com/vuejs/vue-cli/blob/dev/packages/@vue/cli-service/lib/config/app.js#L171
    fileBlacklist: [/\.map$/, /hot-update\.js$/, /runtime\..*\.js$/],
    include: 'initial'
}
])
```

#### svg-sprite-loader

```js
// set svg-sprite-loader
config.module
.rule('svg')
.exclude.add(resolve('src/icons'))
.end()

config.module
.rule('icons')
.test(/\.svg$/)
.include.add(resolve('src/icons'))
.end()
.use('svg-sprite-loader')
.loader('svg-sprite-loader')
.options({
    symbolId: 'icon-[name]'
})
.end()
```

#### vue-loader


#### svg图标处理

无论我们使用那个UI组件, 总会遇到icon不够用的时候, 这时候为了保证icon放大不失真，我们需要使用svg icon, 最常用的svg icon库就是:
[阿里巴巴矢量图标库](https://www.iconfont.cn/search/index?searchType=icon&q=gitee&page=1&fromCollection=-1&fills=&tag=)



```js
npm i -D svg-sprite-loade svgo-loader
```


## 参考 

+ [VUE CLI 全局配置](https://cli.vuejs.org/zh/config/#vue-config-js)
+ [webpack dev-server配置](https://webpack.js.org/configuration/dev-server/)
+ [使用 Preload&Prefetch 优化前端页面的资源加载](https://zhuanlan.zhihu.com/p/273298222)
+ [svg-sprite-loader 使用教程](https://www.jianshu.com/p/70f9c9268c83)
+ [JetBrains svg-sprite-loader](https://github.com/JetBrains/svg-sprite-loader)
+ [使用 svg-sprite-loader、svgo-loader 优化项目中的 Icon](https://juejin.cn/post/6854573215646875655)