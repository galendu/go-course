# 前端开发准备

## 安装NodeJs

到[NodeJs官网](https://nodejs.org/)下载最新的稳定版, 并安装
1. NodeJs 安装  
```sh
# NodeJs版本
> node -v
v16.15.0
# npm包管理工具版本
> npm -v 
8.5.5
```

## 安装Yarn
你可以认为Yarn是npm的增强版, 具体对比可以参考: [Yarn vs npm](https://www.cnblogs.com/ypppt/p/13050845.html)

```sh
# 安装Yarn
> npm install --global yarn
# 查看当前安装的版本
> yarn -v
1.22.18
```

## Yarn 源的管理

默认Yarn使用的是国外的源, 这对于国内开放者而言的体验是很差的(由于网速经常拉去不下来包), 因此我们需要切换源, 而yrm 就是专门用于管理yarn源配置的工具, YARN registry manager(yrm):
```sh
# 安装yrm
> npm install -g yrm

# 查看yrm的版本
> yrm -V    
1.0.6
```

处理这样查看我们可以通过npm来查看当前系统上已经安装的全局工具:
```sh
> npm -g ls
/usr/local/lib
├── corepack@0.10.0
├── npm@8.5.5
├── yarn@1.22.18
└── yrm@1.0.6
```

查看当前有哪些可用的源
```sh
> yrm ls
* npm ---- https://registry.npmjs.org/
  cnpm --- http://r.cnpmjs.org/
  taobao - https://registry.npm.taobao.org/
  nj ----- https://registry.nodejitsu.com/
  rednpm - http://registry.mirror.cqupt.edu.cn/
  npmMirror  https://skimdb.npmjs.com/registry/
  edunpm - http://registry.enpmjs.org/
  yarn --- https://registry.yarnpkg.com
```

最后我们通过yrm来设置我们的源:
```sh
# 使用淘宝的源
> yrm use taobao
   YARN Registry has been set to: https://registry.npm.taobao.org/
   NPM Registry has been set to: https://registry.npm.taobao.org/

# 测试下淘宝源当前下载速度
> yrm test taobao
    * taobao - 273ms
```

## npx安装

npm 从5.2版开始，增加了 npx 命令, 如果没有安装请手动安装:
```sh
# 查看当前npx版本
> npx -v
8.5.5
# 如果没有手动安装到全局
> npm install -g npx
```

## IDE插件安装

以下介绍使用vscode开发vue会使用的到的常见插件:

首先是一些前端开发必备插件:
+ Auto Rename Tag: 自动完成另一侧标签的同步修改

![](./images/auto_rename.png)

+ Beautify: 格式化 html ,js,css

![](./images/beautify.png)

+ ESLint: 代码质量提醒

![](./images/eslint.png)

+ open in browser: 鼠标右键快速在浏览器中打开html文件

![](./images/open-in-browser.png)

+ Code Runner: 快速运行

![](./images/code%20runner.png)


vue开发增强插件:
+ Vue Language Features (Volar): vue3语法支持

![](./images/volar.png)

+ Vue VSCode Snippets: 代码片段

![](./images/vue-snippets.png)


# 参考

+ [npx 使用教程](https://www.ruanyifeng.com/blog/2019/02/npx.html)
+ [node-pre-gyp官方介绍](https://www.npmjs.com/package/@mapbox/node-pre-gyp)
+ [20+前端常用的vscode插件](https://www.php.cn/tool/vscode/475531.html)
+ [用Prettier和ESLint自动格式化修复JavaScript代码](https://juejin.cn/post/6971635051998117924)
+ [2021 VSCode前端插件推荐](https://juejin.cn/post/7014300784649043981)
