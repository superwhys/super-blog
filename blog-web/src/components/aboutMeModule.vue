<template>
  <div class="aboutMeModule">
    <base-main-right-module to-end-point="/about">
      <template v-slot:rightModuleTitle>
        ABOUT ME
      </template>
      <template v-slot:rightModuleMain>
        <a href="/about">
          <img width="100" height="100" :src=personLogo alt="">
        </a>
        <p>Talk is cheap, show me the code!</p>
        <p>✉️ superyong4869@163.com</p>
      </template>
    </base-main-right-module>
  </div>
</template>

<script>
import BaseMainRightModule from "@/components/baseMainRightModule.vue";

export default {
  name: "aboutMeModule",
  components: {BaseMainRightModule},
  data() {
    return {
      personLogo: require("@/assets/person.png")
    }
  }
}
</script>

<style scoped>
.aboutMeModule p {
  text-align: left;
  font-size: 13px;
}
</style>