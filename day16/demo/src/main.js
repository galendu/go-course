import Vue from 'vue'
import App from './App.vue'

// vue实例的配置
Vue.config.productionTip = false


// Root Vue实例
new Vue({
  render: h => h(App),
}).$mount('#app')
