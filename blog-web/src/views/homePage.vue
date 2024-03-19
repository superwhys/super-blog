<template>
  <div class="homePage">
    <base-page>
      <template v-slot:blogHeaderImgInnerShow>
        <div class="imgHeading flyInFromTop">
          My Blog
        </div>
        <span class="imgSubHeading flyInFromTop">
          Thinking will not overcome fear but action will.
        </span>
      </template>
      <template v-slot:blogLeftMainShow>
        <div class="blogItems">
          <blog-list-item
              v-for="item in postList.items"
              class="flyInFromBottom"
              :key="item.metaData.title"
              :title="item.metaData.title"
              :sub-title="item.metaData.subTitle"
              :post-content="item.postContent"
              :meta-data="item.postedTime"
              :to-end-point="item.toEndPoint"
          >
          </blog-list-item>
          <div class="pagination">
            <div class="paginationItems" @click="newerClick" v-show="page !== 1">
              < Newer Posts
            </div>
            <div class="paginationItems" @click="olderClick" v-show="page * pageSize < postList.total">
              Older Posts >
            </div>
          </div>
        </div>
      </template>
      <template v-slot:blogRightMainShow>
        <tags-module class="flyInFromBottom"></tags-module>
        <about-me-module class="flyInFromBottom"></about-me-module>
      </template>
    </base-page>
  </div>
</template>

<script>
import BasePage from "@/components/basePage.vue";
import BlogListItem from "@/components/blogListItem.vue";
import TagsModule from "@/components/tagsModule.vue";
import AboutMeModule from "@/components/aboutMeModule.vue";
import {getBlogItemList} from "@/network/blogItemListRequest";
import {BlogList} from "@/models/blogItem";
import {getTagsList} from "@/network/tagsRequest";
import {TagList} from "@/models/tagItem";

export default {
  name: "homePage",
  components: {AboutMeModule, TagsModule, BlogListItem, BasePage},
  data() {
    return {
      page: 1,
      pageSize: 7,
      postList: new BlogList({items: [], total: 0})
    }
  },
  created() {
    let posts = this.$store.getters.getAllPosts;
    let postList = []
    if (Object.keys(posts).length !== 0) {
      Object.entries(posts).forEach(([key, value]) => {
        postList.push(value)
      });
      this.postList.items = postList
      console.debug(`reload post list`)
      console.debug(this.postList)
    } else {
      getBlogItemList(this.page, this.pageSize).then(resp => {
        console.debug(resp)
        if (resp.data) {
          this.postList = new BlogList(resp.data)
          for (let post of this.postList.items) {
            this.$store.commit('setPost', post);
          }
        }
      });
    }
  },
  methods: {
    newerClick() {
      console.debug(`newerClick`)
      if (this.page > 1) {
        this.page -= 1
        getBlogItemList(this.page, this.pageSize).then(resp => {
          console.debug(resp)
          if (resp.data) {
            this.postList = new BlogList(resp.data)
            for (let post of this.postList.items) {
              this.$store.commit('setPost', post);
            }
          }
        });
      }
    },
    olderClick() {
      console.debug(`olderClick`)
      if (this.page * this.pageSize < this.postList.total) {
        this.page += 1
        getBlogItemList(this.page, this.pageSize).then(resp => {
          console.debug(resp)
          if (resp.data) {
            this.postList = new BlogList(resp.data)
            for (let post of this.postList.items) {
              this.$store.commit('setPost', post);
            }
          }
        });
      }
    }
  }
}
</script>

<style scoped>
.homePage {
  height: 100%;
}

.imgHeading {
  font-size: 80px;
  font-weight: bold;
}

.imgSubHeading {
  font-size: 18px;
}

.blogItems {
  width: 100%;
  display: flex;
  flex-direction: column;
}

.pagination {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}

.paginationItems {
  cursor: pointer;
  font-size: 18px;
  font-weight: bold;
  padding: 17px 25px;
  border: 1px solid var(--solidLineColor);
}

.paginationItems:hover {
  color: #fff;
  background-color: #0085a1;
}
</style>
