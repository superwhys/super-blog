import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

Vue.config.productionTip = false
import VMdPreview from '@kangc/v-md-editor/lib/preview';
import '@kangc/v-md-editor/lib/style/preview.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';
import hljs from 'highlight.js';

VMdPreview.use(githubTheme, {
    Hljs: hljs,
});

Vue.use(VMdPreview);

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
