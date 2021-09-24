<template>
  <div id="app">
    <span>父组件的msg: {{ msg1 }}</span>
    <img alt="Vue logo" src="./assets/logo.png">
    <hello-world v-model="msg1" />
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
  data() {
    return {
      msg1: 'Welcome to Your Vue.js App'
    }
  },
  created() {
    this.$root.$data.b = 2
  },
  components: {
    HelloWorld
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
