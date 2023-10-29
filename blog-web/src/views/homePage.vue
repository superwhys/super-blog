<template>
  <div class="homePage">
    <base-page>
      <template v-slot:blogHeaderImgInnerShow>this is template</template>
      <template v-slot:blogLeftMainShow>
        <div class="blogItems">
          <blog-list-item
              v-for="item in postList"
              :title="item.title"
              :sub-title="item.subTitle"
              :content="item.Content"
              :meta="item.Meta"
          >
          </blog-list-item>
        </div>
      </template>
      <template v-slot:blogRightMainShow>
        <tags-module :tags="tags"></tags-module>
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

export default {
  name: "homePage",
  components: {AboutMeModule, TagsModule, BlogListItem, BasePage},
  data() {
    return {
      tags: [
        {text: "golang", toEndPoint: "/tags/golang"},
        {text: "python", toEndPoint: "/tags/python"},
        {text: "java", toEndPoint: "/tags/java"},
        {text: "c", toEndPoint: "/tags/c"},
        {text: "linux", toEndPoint: "/tags/linux"},
      ],
      postList: [
        {
          title: "Golang详解",
          subTitle: "Golang 底层原理详解",
          Content: "前沿 这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿",
          Meta: "Posted by SuperYong on March 1, 2023"
        },
        {
          title: "Golang详解",
          subTitle: "Golang 底层原理详解",
          Content: "前沿 这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿",
          Meta: "Posted by SuperYong on March 1, 2023"
        },
        {
          title: "Golang详解",
          subTitle: "Golang 底层原理详解",
          Content: "前沿 这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿",
          Meta: "Posted by SuperYong on March 1, 2023"
        },
        {
          title: "Golang详解",
          subTitle: "Golang 底层原理详解",
          Content: "前沿 这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿这是前沿",
          Meta: "Posted by SuperYong on March 1, 2023"
        },
      ]
    }
  },
}
</script>

<style scoped>
.homePage {
  /*height: 100%;*/
}

.blogItems {
  display: flex;
  flex-direction: column;
}
</style>