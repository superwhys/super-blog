<template>
  <div class="basePage">
    <el-container style="min-height: 100vh">
      <el-header style="padding: 0" height="420px">
        <blog-header style="width: 100%; height: 100%">
          <template v-slot:headerImgInnerShow>
            <slot name="blogHeaderImgInnerShow"></slot>
          </template>
        </blog-header>
      </el-header>
      <el-main style="padding: 50px 0; min-height: calc(100vh - 420px - 5vh);">
        <div class="mainContainer">
          <div class="leftContainer">
            <slot name="blogLeftMainShow"></slot>
          </div>
          <div class="rightContainer">
            <slot name="blogRightMainShow"></slot>
          </div>
        </div>
      </el-main>
      <el-footer style="height: 5vh">
        <slot name="blogFooterShow"></slot>
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import BlogHeader from "@/components/blogHeader.vue";

export default {
  name: "basePage",
  components: {BlogHeader}
}
</script>

<style scoped>
.basePage {
  width: 100%;
  height: 100%;
}

.mainContainer {
  width: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: row;
}

@media (min-width: 768px) {
  .mainContainer {
    width: 750px;
  }
}

@media (min-width: 992px) {
  .mainContainer {
    width: 970px;
  }
}

@media (min-width: 1200px) {
  .mainContainer {
    width: 1170px;
  }
}

.leftContainer {
  flex: 8;
  display: flex;
  margin-right: 20px;
}

.rightContainer {
  flex: 2;
  display: flex;
  flex-direction: column;
}

@media (max-width: 900px) {
  .rightContainer {
    display: none;
  }
}
</style>