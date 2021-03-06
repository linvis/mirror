import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: 'Dashboard',
      component: () => import('@/views/dashboard/index'),
      meta: { title: 'Dashboard', icon: 'dashboard' }
    }]
  },

  // health
  {
    path: '/health',
    component: Layout,
    redirect: '/health/show',
    name: 'health',
    meta: { title: 'Health', icon: 'example' },
    children: [
      {
        path: 'sleep',
        name: 'sleep',
        component: () => import('@/views/health/sleep'),
        meta: { title: 'Sleep', icon: 'table' }
      },
      {
        path: 'weight',
        name: 'weight',
        component: () => import('@/views/health/weight'),
        meta: { title: 'weight', icon: 'tree' }
      }
    ]
  },

  {
    path: '/evt',
    component: Layout,
    redirect: '/evt',
    name: 'evt',
    meta: { title: 'Daily Event', icon: 'example' },
    children: [
      {
        path: 'evt',
        name: 'evt',
        component: () => import('@/views/evt/evt'),
        meta: { title: 'Daily Event', icon: 'table' }
      },
      {
        path: 'program',
        name: 'program',
        component: () => import('@/views/evt/program'),
        meta: { title: 'Program', icon: 'table' }
      }
    ]
  },

  // finance
  {
    path: '/finance',
    component: Layout,
    redirect: '/finance/show',
    meta: { title: 'Finance Event', icon: 'example' }
  },

  // library
  {
    path: '/wiki',
    component: Layout,
    redirect: '/wiki/show',
    meta: { title: 'Wiki', icon: 'example' },
    children: [{
      path: 'wiki',
      name: 'wiki',
      component: () => import('@/views/wiki/index'),
      meta: { title: 'Wiki', icon: 'example' }
    }]
  },
  // settings
  {
    path: '/settings',
    component: Layout,
    redirect: '/settings/show',
    meta: { title: 'Settings', icon: 'example' },
    children: [{
      path: 'settings',
      name: 'Settingss',
      component: () => import('@/views/settings/index'),
      meta: { title: 'Settings', icon: 'example' }
    }]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
