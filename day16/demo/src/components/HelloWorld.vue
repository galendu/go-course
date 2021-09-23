<template>
  <div class="hello">
    <h1>{{ value }}</h1>
    <h2>{{ reverseName }}</h2>
   
    <input :value="value" type="text" @input="$emit('input', $event.target.value)">
    <br>
    <input v-focus v-model="name" type="text" @keyup.enter="pressEnter(name)">
    <ul>
      <li v-for="(item, index) in items" :key="item.message">
        {{ item.message }} - {{ index}}
        <br>
        <span v-for="(value, key) in item" :key="key"> {{ value }} {{ key }} <br></span>
      </li>
    </ul>
    <br>
    {{ ts | parseTime }}
    <br>
    <button :disabled="isButtomDisabled" v-on:click="clickButtom" >Button</button>
    <p>
      For a guide and recipes on how to configure / customize this project,<br>
      check out the
      <a href="https://cli.vuejs.org" target="_blank" rel="noopener">vue-cli documentation</a>.
    </p>
    <h3>Installed CLI Plugins</h3>
    <ul>
      <li><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-babel" target="_blank" rel="noopener">babel</a></li>
      <li><a href="https://github.com/vuejs/vue-cli/tree/dev/packages/%40vue/cli-plugin-eslint" target="_blank" rel="noopener">eslint</a></li>
    </ul>
    <h3>Essential Links</h3>
    <ul>
      <li><a href="https://vuejs.org" target="_blank" rel="noopener">Core Docs</a></li>
      <li><a href="https://forum.vuejs.org" target="_blank" rel="noopener">Forum</a></li>
      <li><a href="https://chat.vuejs.org" target="_blank" rel="noopener">Community Chat</a></li>
      <li><a href="https://twitter.com/vuejs" target="_blank" rel="noopener">Twitter</a></li>
      <li><a href="https://news.vuejs.org" target="_blank" rel="noopener">News</a></li>
    </ul>
    <h3>Ecosystem</h3>
    <ul>
      <li><a href="https://router.vuejs.org" target="_blank" rel="noopener">vue-router</a></li>
      <li><a href="https://vuex.vuejs.org" target="_blank" rel="noopener">vuex</a></li>
      <li><a href="https://github.com/vuejs/vue-devtools#vue-devtools" target="_blank" rel="noopener">vue-devtools</a></li>
      <li><a href="https://vue-loader.vuejs.org" target="_blank" rel="noopener">vue-loader</a></li>
      <li><a href="https://github.com/vuejs/awesome-vue" target="_blank" rel="noopener">awesome-vue</a></li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data() {
    return {
      name: '老喻',
      isButtomDisabled: false,
      items: [
        { message: 'Foo', level: 'info' },
        { message: 'Bar', level: 'error'}
      ],
      urlHash: '',
      ts: Date.now(),
    }
  },
  mounted() {
    let that = this
    window.onhashchange = function () {
      console.log('URL发生变化了', window.location.hash);
      that.urlHash = window.location.hash
    };
  },
  watch: {
    urlHash: function(newURL, oldURL) {
      console.log(newURL, oldURL)
    }
  },
  methods: {
    clickButtom() {
      alert("别点我")  
    },
    pressEnter(name) {
      alert(`${name}点击了回车键`)
    },
    reverseData(data) {
      return data.split('').reverse().join('')
    },
    changeProps() {
      console.log(this.value)
      this.$emit('input', this.value)
    }
  },
  computed: {
    reverseName: {
      get() {
        return this.name.split('').reverse().join('')
      },
      set(value) {
        this.name = value.split('').reverse().join('')
      }
    }
  },
  props: {
    value: String,
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
