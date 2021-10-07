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
  {
    path: '/404',
    component: () => import('@/views/common/error-page/404'),
    hidden: true
  },
  // 如果前面所有路径都没有匹配就跳转到404页面
  { path: '*', redirect: '/404', hidden: true }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach(beforeEach)
router.afterEach(afterEach)

export default router;
