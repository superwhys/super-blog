import Vue from 'vue'
import Vuex from 'vuex'
import {BlogItem} from "@/models/blogItem";

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        posts: {}
    },
    getters: {
        getAllPosts: state => {
            return state.posts;
        },
        getPost: (state) => (toEndPoint) => {
            return state.posts[toEndPoint];
        }
    },
    mutations: {
        setPost(state, post = new BlogItem({})) {
            Vue.set(state.posts, post.toEndPoint, post)
        }
    },
    actions: {},
    modules: {}
})
