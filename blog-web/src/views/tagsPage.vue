<template>
  <div class="tagsPage">
    <base-page :show-right="false">
      <template v-slot:blogHeaderImgInnerShow>
        <div class="flyInFromTop">
          {{ tag }}
        </div>
      </template>
      <template v-slot:blogLeftMainShow>
        <!-- all tags list -->
        <all-tags-module class="flyInFromLeft" style="margin: 20px"></all-tags-module>
        <div class="tagsGroup flyInFromBottom" v-for="(groupTagPosts, groupTag) in tagInfoList.tagGroup" :key="groupTag">
          <div class="tagsGroupKey el-icon-paperclip">
            {{ groupTag }}
          </div>
          <ul class="tagsGroupPosts">
            <router-link class="tagsGroupPost" v-for="post in groupTagPosts" :to="'/post'+post.info.toEndPoint" :key="post.info.title" tag="li">
              <div class="tagsGroupPostTitle oneLine">
                {{ post.info.title }}
              </div>
              <div class="tagsGroupPostSubTitle oneLine">
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
    let tag = this.$route.params.tag === undefined ? "" : this.$route.params.tag
    this.tag = tag

    let tagGrp = this.$store.getters.getTagGroup(tag)
    if (tagGrp !== undefined) {
      this.tagInfoList = tagGrp
      return
    }

    getTagsInfo(this.tag).then(resp => {
      console.debug(resp)
      if (resp.data) {
        this.tagInfoList = new TagGroup(resp.data)
        this.$store.commit('setTagGroup', {tag: tag, tagGrp: this.tagInfoList})
      }
    })

    this.tag = this.tag === "" ? "Tags" : `Tag: ${this.tag}`
  },
  watch: {
    $route(to, from) {
      if (to.name !== "tag") {
        return
      }
      let tag = this.$route.params.tag === undefined ? "" : this.$route.params.tag
      this.tag = tag

      let tagGrp = this.$store.getters.getTagGroup(tag)
      if (tagGrp !== undefined) {
        this.tagInfoList = tagGrp
        return
      }

      getTagsInfo(this.tag).then(resp => {
        console.debug(resp)
        if (resp.data) {
          this.tagInfoList = new TagGroup(resp.data)
          this.$store.commit('setTagGroup', {tag: tag, tagGrp: this.tagInfoList})
        }
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
  width: 100%;
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
  width: 100%;
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

.oneLine {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  width: 100%;
  text-align: left;
}
</style>