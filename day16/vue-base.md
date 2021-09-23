# Vue入门基础

Vue 是一套用于构建用户界面的渐进式框架。与其它大型框架不同的是，Vue 被设计为可以自底向上逐层应用

Vue 的核心库只关注视图层，不仅易于上手，还便于与第三方库或既有项目整合, 比如实现拖拽: vue + sortable.js

vue借鉴了很有框架优秀的部分进行了整合:
+ 借鉴 angular 的模板和数据绑定技术
+ 借鉴 react 的组件化和虚拟 DOM 技术

所有框架的逻辑 都是通过js封装，上层概念比如MVVM, 方便快速开发，因此学习任何框架前都比较具体Web基础:
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

请参考 [vue-devtools离线安装](https://www.jianshu.com/p/63f09651724c)

### vscode 插件

+ Beautify: js, css, html 语法高亮差距
+ ESLint: js eslint语法风格检查
+ Auto Rename Tag: tag rename
+ Veter: vue语法高亮插架


## Hello World

到我们的demo工程下面, 使用脚手架初始化vue 前端工程ui
```sh
$ vue create demo
```

通过vue-cli搭建一个vue项目，会自动生成一系列文件，而这些文件具体是怎样的结构、文件对应起什么作用，可以看看下面的解释
```
├── dist/                      # 项目构建后的产物
├── node_module/               #项目中安装的依赖模块
├── public/                    # 纯静态资源, 入口文件也在里面
|── src/
│   ├── main.js                 # 程序入口文件
│   ├── App.vue                 # 程序入口vue组件, 大写字母开头,后缀.vue
│   ├── components/             # 组件
│   │   └── ...
│   └── assets/                 # 资源文件夹，一般放一些静态资源文件, 比如CSS/字体/图片
│       └── ...
├── babel.config.js             # babel 配置文件, es6语法转换
├── .gitignore                  # 用来过滤一些版本控制的文件，比如node_modules文件夹 
└── package.json                # 项目文件，记载着一些命令和依赖还有简要的项目描述信息 
└── README.md                   #介绍自己这个项目的，可参照github上star多的项目。
```

然后运行项目
```sh
$ cd demo 
$ npm run serve
```

我们可以通过dev-tools查看当前页面组建构成

![](./images/vue-tools.jpg)

如何部署:
```sh
$ npm run build ## 会在项目的dist目录下生成html文件, 使用这个静态文件部署即可

## 比如我们使用python快速搭建一个http静态站点, 如果是nginx copy到 对应的Doc Root位置
$ cd dist
$ python3 -m http.server
```

## MVVM如何诞生

现在主流的前端框架都是MVVM模型, MVVM分为三个部分：
+ M（Model，模型层 ）: 模型层，主要负责业务数据相关, 对应vue中的 data部分
+ V（View，视图层）: 视图层，顾名思义，负责视图相关，细分下来就是html+css层, 对应于vue中的模版部分
+ VM（ViewModel, 控制器）: V与M沟通的桥梁，负责监听M或者V的修改，是实现MVVM双向绑定的要点, 对应vue中双向绑定

Vue就是这种思想下的产物, 但是要讲清楚这个东西，我们不妨来看看web技术的进化史

### CGI时代

最早的HTML页面是完全静态的网页，它们是预先编写好的存放在Web服务器上的html文件, 浏览器请求某个URL时，Web服务器把对应的html文件扔给浏览器，就可以显示html文件的内容了

如果要针对不同的用户显示不同的页面，显然不可能给成千上万的用户准备好成千上万的不同的html文件，所以，服务器就需要针对不同的用户，动态生成不同的html文件。一个最直接的想法就是利用C、C++这些编程语言，直接向浏览器输出拼接后的字符串。这种技术被称为CGI：Common Gateway Interface

下面是一个python的cgi样例:

![](./images/python-cgi.jpeg)


### 后端模版时代

很显然，像新浪首页这样的复杂的HTML是不可能通过拼字符串得到的, 于是，人们又发现，其实拼字符串的时候，大多数字符串都是HTML片段，是不变的，变化的只有少数和用户相关的数据, 所以我们做一个模版出来，把不变的部分写死, 变化的部分动态生成, 其实就是一套模版渲染系统, 其中最典型的就是:
+ ASP: 微软, C#体系
+ JSP: SUN, Java体系
+ PHP: 开源社区

下面是一段PHP样例:

![](./images/php.jpg)

但是，一旦浏览器显示了一个HTML页面，要更新页面内容，唯一的方法就是重新向服务器获取一份新的HTML内容。如果浏览器想要自己修改HTML页面的内容，怎么办？那就需要等到1995年年底，JavaScript被引入到浏览器

有了JavaScript后，浏览器就可以运行JavaScript，然后，对页面进行一些修改。JavaScript还可以通过修改HTML的DOM结构和CSS来实现一些动画效果，而这些功能没法通过服务器完成，必须在浏览器实现

### JavaScript原生时代

```html
<p id="userInfo">
姓名:<span id="name">Gloria</span>
性别:<span id="sex">男</span>
职业:<span id="job">前端工程师</span>
</p>
```

有以上html片段，想将其中个人信息替换为alice的，我们的做法
```js
// 通过ajax向后端请求, 然后利用js动态修改展示页面
document.getElementById('name').innerHTML = alice.name;
document.getElementById('sex').innerHTML = alice.sex;
document.getElementById('job').innerHTML = alice.job;
```

jQuery在这个时代脱颖而出

```html
<div id="name" style="color:#fff">前端你别闹</div> <div id="age">3</div>
<script>
$('#name').text('好帅').css('color', '#000000'); $('#age').text('666').css('color', '#fff');
/* 最终页面被修改为 <div id="name" style="color:#fff">好帅</div> <div id="age">666</div> */
</script>
```

在此情况下可以下 前后端算是分开了, 后端提供数据, 前端负责展示, 只是现在 前端里面的数据和展示并有分开，不易于维护


### 前端模版时代

在架构上前端终于走上后端的老路: 模版系统, 有引擎就是ViewModel动态完成渲染

```html
<script id="userInfoTemplate">
姓名:<span>{name}</span>
性别:<span>{sex}</span>
职业:<span>{job}</span>
</script>
```

```js
var userInfo = document.getElementById('userInfo');
var userInfoTemplate = document.getElementById('userInfoTemplate').innerHTML;
userInfo.innerHTML = templateEngine.render(userInfoTemplate, users.alice);
```

### 虚拟DOM技术

用我们传统的开发模式，原生JS或JQ操作DOM时，浏览器会从构建DOM树开始从头到尾执行一遍流程, 操作DOM的代价仍旧是昂贵的，频繁操作还是会出现页面卡顿，影响用户体验

如何才能对Dom树做局部更新，而不是全局更新喃? 答案就是由js动态来生产这个树, 修改时动态更新, 这就是虚拟Dom

这是一个真实dom

![](./images/raw-dom.png)

这是动态生成的Dom是不是和CGI很像

![](./images/vdom.png)

### 组件化时代

MVVM最早由微软提出来，它借鉴了桌面应用程序的MVC思想，在前端页面中，把Model用纯JavaScript对象表示，View负责显示，两者做到了最大限度的分离。

![](./images/vue-mvvm.png)

结合虚拟Dom技术, 我们就可以动态生成view, 在集合mvvm的思想, 前端终于迎来了组件化时代

+ 页面由多个组建构成
+ 每个组件都有自己的 MVVM

![](./images/vue-components.png)


## Vue与MVVM

+ Model: vue中用于标识model的数据是 data对象, data 对象中的所有的 property 加入到 Vue 的响应式系统中, 有vue监听变化
+ View: vue使用模版来实现展示, 但是渲染时要结合 vdom技术
+ ViewModle: vue的核心, 负责视图的响应, 也就是数据双向绑定
    + 监听view中的数据,  如果数据有变化, 动态同步到 data中
    + 监听data中的数据,  如果数据有变化, 通过vdom动态渲视图


比如在我们的demo中, 修改HelloWorld, 在model中添加一个name属性
```js
<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      name: '老喻'
    }
  },
  props: {
    msg: String
  }
}
</script>
```

然后在模版中添加个输入框来修改他, 给他h2展示name属性
```html
<h2>{{ name }}</h2>
<input v-model="name" type="text">
```

![](./images/binding-example.jpg)

在安装了Vue Devtools的时候，我们可以在console里看到我们的虚拟dom

![](./images/vm-console.png)


## Vue实例




## 参考

+ [VUE2官方文档](https://cn.vuejs.org/v2/guide)
+ [那些前端MVVM框架是如何诞生的](https://zhuanlan.zhihu.com/p/36453279)
+ [MVVM设计模式](https://zhuanlan.zhihu.com/p/36141662)
+ [vue核心之虚拟DOM(vdom)](https://www.jianshu.com/p/af0b398602bc)