import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

/* Layout */
import Layout from "@/layout";

const routes = [
  {
    path: "/login",
    component: () => import("@/views/login/index"),
    hidden: true
  },
  {
    path: "/404",
    component: () => import("@/views/404"),
    hidden: true
  },
  {
    path: "/",
    name: "root",
    redirect: "/home",
    component: Layout,
    hidden: true
  },
  {
    path: "/home",
    name: "Home",
    component: Layout,
    children: [
      {
        path: "sumary",
        name: "Sumary",
        component: () => import("@/views/home/index"),
        meta: { text: "Home", icon: "mdi-home" }
      }
    ]
  },
  {
    path: "/healthy",
    name: "Healthy",
    component: Layout,
    meta: { text: "Healthy", icon: "mdi-hand-heart" },
    children: [
      {
        path: "sleep",
        name: "Sleep",
        component: () => import("@/views/healthy/sleep"),
        meta: { text: "Sleep", icon: "mdi-power-sleep" }
      },
      {
        path: "weight",
        name: "Weight",
        component: () => import("@/views/home/index"),
        meta: { text: "Weight", icon: "mdi-weight-kilogram" }
      }
    ]
  },
  {
    path: "/wiki",
    name: "Wiki",
    component: Layout,
    meta: { text: "Wiki", icon: "mdi-wikipedia" },
    children: [
      {
        path: "page",
        name: "Wiki",
        component: () => import("@/views/wiki/index"),
        meta: { text: "Wiki", icon: "mdi-wikipedia" }
      }
    ]
  },
  {
    path: "/program",
    name: "Program",
    component: Layout,
    meta: { text: "Program", icon: "mdi-wikipedia" },
    children: [
      {
        path: "sumary",
        name: "sumary",
        component: () => import("@/views/program/index"),
        meta: { text: "Program", icon: "mdi-wikipedia" }
      }
    ]
  },
  {
    path: "/settings",
    name: "Settings",
    component: Layout,
    meta: { text: "Settings", icon: "mdi-wikipedia" },
    children: [
      {
        path: "",
        name: "",
        component: () => import("@/views/settings/index"),
        meta: { text: "settings", icon: "mdi-wikipedia" }
      }
    ]
  },
  { path: "*", redirect: "/404", hidden: true }
];

const router = new VueRouter({
  scrollBehavior: () => ({ y: 0 }),
  routes: routes
});

export default router;
