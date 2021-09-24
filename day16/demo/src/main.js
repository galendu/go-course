import Vue from 'vue'
import App from './App.vue'
import router from './router'

// vue实例的配置
Vue.config.productionTip = false

// 添加全局过滤器
Vue.filter('parseTime', function (value) {
  let date = new Date(value)
  return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()} ${date.getHours()}:${date.getMinutes()}`
})

// 注册一个全局自定义指令 `v-focus`
Vue.directive('focus', {
  // 当被绑定的元素插入到 DOM 中时……
  inserted: function (el) {
    // 聚焦元素
    el.focus()
  }
})


// Root Vue实例
new Vue({
  el:'#app',
  router,
  render: h => h(App)
})
