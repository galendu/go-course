# 前端框架搭建

这里我们直接使用vue-router和vuex
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