# Vue路由与状态管理


## 页面路由

我们看看当前的vue配置
```js
// Root Vue实例
new Vue({
  render: h => h(App),
}).$mount('#app')
```

你会发现无论我们访问的url path怎么变化 访问到的都是这个页面, 因为我们在渲染的时候并没有 按照path 来进行不同视图进行渲染的路由逻辑

### 简单路由


[自己处理路由](https://cn.vuejs.org/v2/guide/routing.html)

### vue-router




## 页面状态



### 共享内存

```js
var vmA = new Vue({
  data: sourceOfTruth
})
```

### 本地存储

浏览器提供一些存储功能可供我们存储客户端临时信息:

![](./images/web-storage.jpg)

+ cookie
+ sessionStorage
+ localStorage



#### cookie

cookie是有可以设置过期时间的, 同一个域下的页面都可以访问

cookie在没有设置过期时间时，系统默认浏览器关闭时失效，只有设置了没到期的保存日期时，浏览器才会把cookie作为文件保存在本地上。当expire到期时，cookie不会自动删除，仅在下次启动浏览器或者刷新浏览器时，浏览器会检测cookie过期时间，如已过期浏览器则会删除过期cookie

注意:
  + 数据存放大小: 4k, 因为每次http请求都会携带cookie
  + 浏览器关闭时， cookie会失效
  + 注意cookie可以支持httpOnly, 这个时候前端js是修改不了的

```js
// 读取cookie, 注意读取出来的cookie是个字符串
document.cookie
'language=zh; Sub-System=develop; sidebarStatus=1; Current-Namespace=c16mhsddrei91m4ri0jg; Refresh-Token=paBuyTIfsX3BeKrXrCmD8khUla6x8y1g'
// 需要自己处理
document.cookie.split('; ')

// 直接赋值就添加了一个key-value
document.cookie = 'cookieKey=cookieValue'
'language=zh; Sub-System=develop; sidebarStatus=1; Current-Namespace=c16mhsddrei91m4ri0jg; Refresh-Token=paBuyTIfsX3BeKrXrCmD8khUla6x8y1g; cookieKey=cookieValue'

// 当然cookie还有很多选项可以设置, 通过;隔开比如
document.cookie="username=John Doe; expires=Thu, 18 Dec 2043 12:00:00 GMT; path=/";

// 修改cookie和设置cookie一样, 保证key相同就可以
document.cookie = 'cookieKey=cookieValue2'
document.cookie
'language=zh; Sub-System=develop; sidebarStatus=1; Current-Namespace=c16mhsddrei91m4ri0jg; Refresh-Token=paBuyTIfsX3BeKrXrCmD8khUla6x8y1g; cookieKey=cookieValue2'

// 删除cookie时，把expires 设置到过期的时间即可, 比如设置个2019年的时间
document.cookie = `cookieKey=;expires=Mon, 26 Aug 2019 12:00:00 UTC`
document.cookie
'language=zh; Sub-System=develop; sidebarStatus=1; Current-Namespace=c16mhsddrei91m4ri0jg; Refresh-Token=paBuyTIfsX3BeKrXrCmD8khUla6x8y1g'
```



#### sessionStorage

存储的数据只有在同一个会话中的页面才能访问并且当会话结束后数据也随之销毁, 因此sessionStorage不是一种持久化的本地存储，仅仅是会话级别的存储

那么，到底什么是一个会话？多个标签页之间的数据是否会共享呢？

我们可以验证下: 开启2个窗口, 直接通过浏览器修改sessionStorage 进行验证

通过验证我们可以知道 一个标签页 就表示一个回话, 当标签页关闭, 回话就清除,  不通标签页之间不共享数据


```js
```



#### localStorage

localStorage生命周期是永久, 除非主动删除数据，否则数据是永远不会过期的

相同浏览器的不同页面间可以共享相同的 localStorage（页面属于相同域名和端口）

我们可以验证下: 开启2个窗口, 直接通过浏览器修改localStorage 进行验证

```js
```


### vuex




### vuex-persist

Vuex的状态存储并不能持久化，存储在 Vuex 中的 store 里的数据，只要一刷新页面，数据就丢失了

具体使用说明请参考[vuex-persist Github](https://github.com/championswimmer/vuex-persist)


## 参考

+ [Vue Router文档](https://next.router.vuejs.org/zh/introduction.html)
+ [Vuex 文档](https://vuex.vuejs.org/zh/)
+ [cookies、sessionStorage和localStorage解释及区别](https://www.cnblogs.com/pengc/p/8714475.html)
+ [JavaScript Cookie](https://www.runoob.com/js/js-cookies.html)
+ [JavaScript创建、读取和删除cookie](https://www.jb51.net/article/169117.htm)