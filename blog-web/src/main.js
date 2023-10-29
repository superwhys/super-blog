import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false

import {
    Container, Header, Main, Footer,
    Menu, MenuItem
} from "element-ui";

Vue.use(Container)
Vue.use(Header)
Vue.use(Main)
Vue.use(Footer)
Vue.use(Menu)
Vue.use(MenuItem)

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
