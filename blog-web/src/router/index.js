import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const homePage = () => import('@/views/homePage.vue')
const aboutPage = () => import('@/views/aboutPage.vue')
const tagsPage = () => import('@/views/tagsPage.vue')


const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/home',
    name: 'home',
    component: homePage
  },
  {
    path: '/about',
    name: 'about',
    component: aboutPage
  },
  {
    path: '/tags',
    name: 'tags',
    component: tagsPage
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
