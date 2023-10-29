<template>
  <div class="aboutPage">
    <base-page>
      <template v-slot:blogHeaderImgInnerShow>
        <div class="title">About</div>
        <div class="subTitle">Hey, this is SuperYong</div>
      </template>
    </base-page>
  </div>
</template>

<script>
import BasePage from "@/components/basePage.vue";

export default {
  name: "aboutPage",
  components: {BasePage}
}
</script>

<style scoped>
.title {
  font-size: 50px;
  font-weight: bold;
}

.subTitle {
  font-size: 20px;
}
</style>