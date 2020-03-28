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
    meta: { text: "Home", icon: "mdi-home" },
    children: [
      {
        path: "/",
        name: "Home",
        component: () => import("@/views/home/index"),
        meta: { text: "Home", icon: "mdi-home" }
      }
    ]
  },
  {
    path: "/healthy",
    name: "Healthy",
    component: Layout,
    meta: { text: "Healthy", icon: "mdi-spa" },
    redirt
    children: [
      {
        path: "/healthy/sleep",
        name: "Sleep",
        component: () => import("@/views/About")
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
