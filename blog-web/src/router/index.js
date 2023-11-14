import Vue from 'vue'
import VueRouter from 'vue-router'

const basePath = window.location.pathname.split('/')[1];

Vue.use(VueRouter)

const homePage = () => import('@/views/homePage.vue')
const aboutPage = () => import('@/views/aboutPage.vue')
const tagsPage = () => import('@/views/tagsPage.vue')
const blogPage = () => import('@/views/blogPage.vue')


const routes = [
    {
        path: '/',
        name: 'home',
        component: homePage,
        meta: {keepAlive: true}
    },
    {
        path: '/about',
        name: 'about',
        component: aboutPage,
        meta: {keepAlive: true}
    },
    {
        path: '/tag/:tag?',
        name: 'tag',
        component: tagsPage,
        meta: {keepAlive: true}
    },
    {
        path: '/post/:year/:month/:day/:name',
        name: 'blog',
        component: blogPage,
        meta: {keepAlive: false}
    }
]

const router = new VueRouter({
    mode: 'history',
    base: `/${basePath}`,
    // base: process.env.VUE_APP_BASE_URL,
    routes
})

export default router
