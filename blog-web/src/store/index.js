import Vue from 'vue'
import Vuex from 'vuex'
import {BlogItem} from "@/models/blogItem";
import {TagGroup, TagItem, TagList} from "@/models/tagItem";

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        posts: {},
        tags: {},
        tagGroups: {},
    },
    getters: {
        getAllPosts: state => {
            return state.posts;
        },
        getPost: (state) => (toEndPoint) => {
            return state.posts[toEndPoint];
        },
        getTags: state => {
            return state.tags
        },
        getTagGroup: (state) => (tag) => {
            if (tag === "") {
                return state.tagGroups["all"]
            }
            return state.tagGroups[tag]
        },
    },
    mutations: {
        setPost(state, post = new BlogItem({})) {
            Vue.set(state.posts, post.toEndPoint, post)
        },
        setTag(state, tag = new TagItem({})) {
            Vue.set(state.tags, tag.tag, tag)
        },
        setTagGroup(state, {tag="", tagGrp = new TagGroup({})}) {
            if (tag === "") {
                tag = "all"
            }
            Vue.set(state.tagGroups, tag, tagGrp)
        }
    },
    actions: {},
    modules: {}
})
