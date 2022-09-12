import { createRouter, createWebHistory } from "vue-router";
import FrontendLayout from "../layout/FrontendLayout.vue";
import BackendLayout from "../layout/BackendLayout.vue";
import BlogView from "../views/frontend/BlogView.vue";
import BlogDetail from "../views/frontend/BlogDetail.vue";
import BlogList from "../views/backend/BlogList.vue";
import TagList from "../views/backend/TagList.vue";

import { beforeEachHandler, afterEachHandler } from "./permession";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      redirect: "/frontend",
    },
    {
      path: "/frontend",
      name: "frontend",
      component: FrontendLayout,
      children: [
        {
          path: "",
          name: "frontend",
          component: BlogView,
        },
        {
          path: "blogs/:id",
          name: "BlogDetail",
          component: BlogDetail,
        },
      ],
    },
    {
      path: "/backend",
      name: "backend",
      component: BackendLayout,
      children: [
        {
          path: "blogs",
          name: "BlogList",
          component: BlogList,
        },
        {
          path: "tags",
          name: "TagList",
          component: TagList,
        },
      ],
    },
    {
      path: "/login",
      name: "LoginPage",
      component: () => import("@/views/login/LoginPage.vue"),
    },
    {
      path: "/errors/403",
      name: "PermissionDeny",
      component: () => import("@/views/errors/PermissionDeny.vue"),
    },
    {
      path: "/:pathMatch(.*)*",
      name: "NotFound",
      component: () => import("@/views/errors/NotFound.vue"),
    },
  ],
});

// 补充导航守卫
router.beforeEach(beforeEachHandler);
router.afterEach(afterEachHandler);

export default router;
