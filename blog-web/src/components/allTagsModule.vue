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
    let allTags = this.$store.getters.getTags;
    let tags = []
    for (let key in allTags) {
      tags.push(allTags[key])
    }
    if (tags.length !== 0) {
      this.tags.tags = tags
      console.log(this.tags)
      return
    }

    getTagsList().then(resp => {
      console.debug(resp)
      if (resp.data) {
        this.tags = new TagList(resp.data)
        for (let tag of this.tags.tags) {
          this.$store.commit('setTag', tag)
        }
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