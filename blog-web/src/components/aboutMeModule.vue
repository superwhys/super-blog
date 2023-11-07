<template>
  <div class="aboutMeModule">
    <base-main-right-module to-end-point="/about">
      <template v-slot:rightModuleTitle>
        ABOUT ME
      </template>
      <template v-slot:rightModuleMain>
        <router-link to="/about">
          <img class="personLogo" width="200" height="200" :src=personLogo alt="">
        </router-link>
        <p class="aboutModuleText">Talk is cheap, show me the code!</p>
        <p class="aboutModuleText">✉️ superyong4869@163.com</p>
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

.personLogo {
  border-radius: 10px;
}

.aboutModuleText {
  margin: 5px 0;
}
</style>