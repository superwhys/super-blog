<template>
  <div class="allTags">
    <tag-item text="All" to-end-point="/tag"></tag-item>
    <tag-item v-for="item in tags.tags" :key="item.tag" :text="item.tag" :to-end-point="item.toEndPoint" style="margin-left: 5px"></tag-item>
  </div>
</template>

<script>
import {TagList} from "@/models/tagItem";
import {getTagsList} from "@/network/tagsRequest";
import TagItem from "@/components/tagItem.vue";
export default {
  name: "allTagsModule",
  components: {TagItem},
  data() {
    return {
      tags: new TagList({tags: []}),
    }
  },
  created() {
    getTagsList().then(resp => {
      console.debug(resp)
      if (resp.data) {
        this.tags = new TagList(resp.data)
      }
    })
  },
}
</script>

<style scoped>

.allTags {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  flex-wrap: wrap;
}
</style>