<template>
  <div class="tagsModule">
    <base-main-right-module to-end-point="/tag">
      <template v-slot:rightModuleTitle>
        FEATURED TAGS
      </template>
      <template v-slot:rightModuleMain>
        <div class="tagItems">
          <tag-item v-for="item in tags" :key="item.key" :text="item.tag" :to-end-point="item.toEndPoint"></tag-item>
        </div>
      </template>
    </base-main-right-module>
  </div>
</template>

<script>
import TagItem from "@/components/tagItem.vue";
import BaseMainRightModule from "@/components/baseMainRightModule.vue";

export default {
  name: "tagsModule",
  components: {BaseMainRightModule, TagItem},
  props: {
    tags: {
      type: Array,
      default: [{tag: "", toEndPoint: "#"}]
    }
  }
}
</script>

<style scoped>
.tagItems {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
</style>