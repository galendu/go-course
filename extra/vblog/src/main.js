import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";
import router from "./router";
import "./assets/main.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);

// 安装UI插件
import ArcoVue from "@arco-design/web-vue";
import "@arco-design/web-vue/dist/arco.css";
app.use(ArcoVue);
// 额外引入图标库
import ArcoVueIcon from "@arco-design/web-vue/es/icon";
app.use(ArcoVueIcon);

app.mount("#app");
