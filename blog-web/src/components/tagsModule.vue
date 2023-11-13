<template>
  <div class="tagsModule">
    <base-main-right-module to-end-point="/tag">
      <template v-slot:rightModuleTitle>
        FEATURED TAGS
      </template>
      <template v-slot:rightModuleMain>
        <all-tags-module></all-tags-module>
      </template>
    </base-main-right-module>
  </div>
</template>

<script>
import TagItem from "@/components/tagItem.vue";
import BaseMainRightModule from "@/components/baseMainRightModule.vue";
import AllTagsModule from "@/components/allTagsModule.vue";

export default {
  name: "tagsModule",
  components: {AllTagsModule, BaseMainRightModule, TagItem},
}
</script>

<style scoped>
</style>