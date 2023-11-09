<template>
  <div class="homePage">
    <base-page>
      <template v-slot:blogHeaderImgInnerShow>this is template</template>
      <template v-slot:blogLeftMainShow>
        <div class="blogItems">
          <blog-list-item
              v-for="item in postList.items"
              :key="item.title"
              :title="item.title"
              :sub-title="item.subTitle"
              :post-content="item.postContent"
              :meta-data="item.postedTime"
              :to-end-point="item.toEndPoint"
          >
          </blog-list-item>
        </div>
      </template>
      <template v-slot:blogRightMainShow>
        <tags-module :tags="tags.tags"></tags-module>
        <about-me-module></about-me-module>
      </template>
      <template v-slot:blogFooterShow>Footer</template>
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
      tags: new TagList({tags: []}),
      postList: new BlogList({items: []})
    }
  },
  created() {
    getBlogItemList().then(resp => {
      console.debug(resp)
      this.postList = new BlogList(resp.data)
    });
    getTagsList().then(resp => {
      console.debug(resp)
      this.tags = new TagList(resp.data)
    })
  }
}
</script>

<style scoped>
.homePage {
  /*height: 100%;*/
}

.blogItems {
  width: 100%;
  display: flex;
  flex-direction: column;
}
</style>