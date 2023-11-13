<template>
  <div class="tagsPage">
    <base-page>
      <template v-slot:blogHeaderImgInnerShow>
        {{ tag }}
      </template>
      <template v-slot:blogLeftMainShow>
        <!-- all tags list -->
        <all-tags-module style="margin: 20px"></all-tags-module>
        <div class="tagsGroup" v-for="(groupTagPosts, groupTag) in tagInfoList.tagGroup" :key="groupTag">
          <div class="tagsGroupKey el-icon-paperclip">
            {{ groupTag }}
          </div>
          <ul class="tagsGroupPosts">
            <router-link class="tagsGroupPost" v-for="post in groupTagPosts" :to="'/post'+post.info.toEndPoint" :key="post.info.title" tag="li">
              <div class="tagsGroupPostTitle">
                {{ post.info.title }}
              </div>
              <div class="tagsGroupPostSubTitle">
                {{ post.info.subTitle }}
              </div>
            </router-link>
          </ul>
        </div>
      </template>
      <template v-slot:blogRightMainShow>
        <div style="display: none"></div>
      </template>
    </base-page>
  </div>
</template>

<script>
import BasePage from "@/components/basePage.vue";
import {getTagsInfo, getTagsList} from "@/network/tagsRequest";
import {TagGroup, TagList} from "@/models/tagItem";
import TagItem from "@/components/tagItem.vue";
import AllTagsModule from "@/components/allTagsModule.vue";

export default {
  name: "tagsPage",
  components: {AllTagsModule, TagItem, BasePage},
  data() {
    return {
      tag: "",
      tagInfoList: new TagGroup({}),
    }
  },
  created() {
    this.tag = this.$route.params.tag === undefined ? "" : this.$route.params.tag
    getTagsInfo(this.tag).then(resp => {
      console.debug(resp)
      if (resp.data) {
        this.tagInfoList = new TagGroup(resp.data)
      }
    })

    this.tag = this.tag === "" ? "Tags" : `Tag: ${this.tag}`
  },
  watch: {
    $route(to, from) {
      this.tag = this.$route.params.tag === undefined ? "" : this.$route.params.tag
      getTagsInfo(this.tag).then(resp => {
        console.debug(resp)
        this.tagInfoList = new TagGroup(resp.data)
      })
      this.tag = this.tag === "" ? "Tags" : `Tag: ${this.tag}`
    }
  }
}
</script>

<style scoped>
.tagsGroupKey {
  font-size: 24px;
  color: var(--hightlight);
}

.tagsGroup {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  margin: 25px 0 0 0;
}

.tagsGroupPosts {
  list-style: none;
  display: flex;
  flex-direction: column;
  padding: 0;
  margin: 10px 0;
}

.tagsGroupPost {
  height: 100px;
  width: 100%;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  flex-direction: column;
  user-select: none;
  cursor: pointer;
  border-bottom: 1px solid var(--solidLineColor);
}

.tagsGroupPost:hover {
  color: var(--hightlight);
}

.tagsGroupPostTitle {
  font-size: 20px;
  font-weight: bold;
  padding-bottom: 10px;
}
</style>