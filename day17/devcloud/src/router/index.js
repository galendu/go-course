import Vue from "vue";
import VueRouter from "vue-router";
import {beforeEach, afterEach} from './permission'

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

router.beforeEach(beforeEach)
router.afterEach(afterEach)

export default router;
