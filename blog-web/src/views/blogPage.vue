<template>
  <div class="postPage">
    <base-page>
      <template v-slot:blogHeaderImgInnerShow>
        <div class="imgTitle">
          {{ postItem.title }}
        </div>
        <div class="imgTag">
          <tag-item
              v-for="item in postItem.metaData.tags"
              :key="item"
              :text="item"
              :to-end-point="'/tag/'+ item"
              style="margin-left: 5px; color: white"
          ></tag-item>
        </div>
      </template>
      <template v-slot:blogLeftMainShow>
        <v-md-preview :text="postItem.postContent" style="text-align: left;" ref="preview"></v-md-preview>
      </template>
      <template v-slot:blogRightMainShow>
        <div class="anchor">
          <div style="text-align: left; font-size: 20px; font-weight: bold; padding-bottom: 15px">
            - CATALOG
          </div>
          <div
              class="anchorItem"
              v-for="anchor in titles"
              :style="{ padding: `10px 0 10px ${anchor.indent * 20}px`}"
              @click="handleAnchorClick(anchor)"
          >
            <a style="cursor: pointer">{{ anchor.title }}</a>
          </div>
        </div>
      </template>
    </base-page>
  </div>
</template>

<script>
import BasePage from "@/components/basePage.vue";
import {getPost} from "@/network/postRequest";
import {BlogItem} from "@/models/blogItem";
import TagItem from "@/components/tagItem.vue";

export default {
  name: "blogPage",
  components: {TagItem, BasePage},
  data() {
    return {
      year: "",
      month: "",
      day: "",
      name: "",
      titles: [],
      mdToc: [],
      postItem: new BlogItem({})
    }
  },
  created() {
    this.year = this.$route.params.year
    this.month = this.$route.params.month
    this.day = this.$route.params.day
    this.name = this.$route.params.name

    getPost(this.year, this.month, this.day, this.name).then(resp => {
      console.debug(resp)
      if (resp.data) {
        this.postItem = new BlogItem(resp.data)

        this.$nextTick(() => {
          this.findTitles();
        });
      }
    })
  },
  methods: {
    findTitles() {
      this.$nextTick(() => {
        const anchors = this.$refs.preview.$el.querySelectorAll('h1,h2,h3,h4,h5,h6');
        const titles = Array.from(anchors).filter((title) => !!title.innerText.trim());
        if (!titles.length) {
          this.titles = [];
          return;
        }

        const hTags = Array.from(new Set(titles.map((title) => title.tagName))).sort();

        this.titles = titles.map((el) => ({
          title: el.innerText,
          lineIndex: el.getAttribute('data-v-md-line'),
          indent: hTags.indexOf(el.tagName),
        }));
      })
    },
    handleAnchorClick(anchor) {
      const {lineIndex} = anchor;
      const heading = this.$refs.preview.$el.querySelector(`[data-v-md-line="${lineIndex}"]`);

      if (heading) {
        window.scrollTo({
          top: heading.offsetTop - 60,
          behavior: 'smooth'
        });
      }
    },
  }
}
</script>

<style scoped>
.postPage {
  height: 100%;
}

.anchor {
  position: sticky;
  top: 100px;
}

.anchorItem {
  font-size: 14px;
  text-align: left;
}
</style>