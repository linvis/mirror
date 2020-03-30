import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

/* Layout */
import Layout from "@/layout";

const routes = [
  {
    path: "/",
    name: "root",
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
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
