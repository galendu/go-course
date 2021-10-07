import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: '/login',
    name: "Login",
    component: () => import('../views/keyauth/login/index'),
  },
  {
    path: '/',
    name: "Home",
    component: () => import('../views/dashboard/index'),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
