import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

/* Layout */
import Layout from "@/layout";

const routes = [
  {
    path: "/",
    name: "root",
    component: Layout
  },
  {
    path: "/home",
    name: "Home",
    component: Layout,
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
    path: "/about",
    name: "About",
    component: Layout,
    children: [
      {
        path: "/",
        name: "About",
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
