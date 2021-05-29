# 环境搭建

## 在线体验

在线演练地址: https://play.golang.org/

## 安装Go

参考官方文档: https://golang.org/doc/install

安装包下载地址:
  + https://golang.org/dl/
  + https://golang.google.cn/dl/


## 安装Git
先安装git, 准备好git shell 具体查看官方文档: https://git-scm.com/


## 开发工具

+ Visual Studio code
+ Sublime Text
+ Atom
+ Eclipse+GoClipse
+ LiteIDE
+ Goland
+ Vim
+ Emacs
+ ……

## vscode 介绍

### Go开发环境安装
```sh
// 设置代理
go env -w GOPROXY=https://goproxy.cn
// 下载依赖工具
go get -v golang.org/x/tools/gopls
go get -v honnef.co/go/tools/cmd/staticcheck
go get -v github.com/go-delve/delve/cmd/dlv@master
go get -v github.com/go-delve/delve/cmd/dlv
go get -v github.com/haya14busa/goplay/cmd/goplay
go get -v github.com/josharian/impl
go get -v github.com/fatih/gomodifytags
go get -v github.com/cweill/gotests/gotests
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs
```

### 设置终端使用Git Shell

The easiest way now (at least from Visual Studio Code 1.22 on) is to type Shift + Ctrl + P to open the Command Palette and type:

Select Default Shell

### 插件安装

+ 快捷运行代码的插件: Code Runner
+ 最好用Git工具没有之一: Gitlens

## 作业
 
+ 使用vscode debug hello world程序, 完整了解程序的执行流程