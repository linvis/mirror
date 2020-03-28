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
    meta: { text: "Healthy", icon: "mdi-spa" },
    children: [
      {
        path: "sleep",
        name: "Sleep",
        component: () => import("@/views/home/index"),
        meta: { text: "Sleep", icon: "mdi-alarm" }
      },
      {
        path: "weight",
        name: "Weight",
        component: () => import("@/views/home/index"),
        meta: { text: "Weight", icon: "mdi-spa" }
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
