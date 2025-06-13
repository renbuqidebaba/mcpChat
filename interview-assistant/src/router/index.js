import { createRouter, createWebHistory } from 'vue-router'
import Login from '../components/Login.vue'
import MainLayout from '../components/MainLayout.vue'
import UserCenter from '../components/UserCenter.vue'
import ChatInterface from '../components/ChatInterface.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: Login
  },
  {
    path: '/main',
    name: 'Main',
    component: MainLayout,
    meta: { requiresAuth: true }
  },
  {
    path: '/chat',
    name: 'Chat',
    component: ChatInterface,
    meta: { requiresAuth: true }
  },
  {
    path: '/user',
    name: 'UserCenter',
    component: UserCenter,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

// 路由守卫 - 检查登录状态
router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const currentUser = localStorage.getItem('currentUser')
  
  if (requiresAuth && (!currentUser || currentUser === 'null' || currentUser === 'undefined')) {
    // 需要登录但未登录，跳转到登录页
    next('/')
  } else if (to.path === '/' && currentUser && currentUser !== 'null' && currentUser !== 'undefined') {
    // 已登录用户访问登录页，跳转到主页
    next('/main')
  } else {
    next()
  }
})

export default router