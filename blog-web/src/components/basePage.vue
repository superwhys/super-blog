<template>
  <div class="basePage">
    <el-container style="min-height: 100vh;">
      <el-header :style="backgroundImageStyle">
        <nav-bar-module></nav-bar-module>
        <blog-header style="width: 100%; height: 100%">
          <template v-slot:headerImgInnerShow>
            <slot name="blogHeaderImgInnerShow"></slot>
          </template>
        </blog-header>
      </el-header>
      <el-main style="overflow: visible; height: 100%; padding: 50px 0; min-height: calc(100vh - 420px - 5vh);">
        <div class="mainContainer">
          <div class="leftContainer">
            <slot name="blogLeftMainShow"></slot>
          </div>
          <div class="rightContainer" :style="{display: showRight===true?'block':'none'}">
            <slot name="blogRightMainShow"></slot>
          </div>
        </div>
      </el-main>
      <el-footer class="blogFooter" style="height: 100%">
        <div>
          <slot name="blogFooterShow"></slot>
          <span style="color: grey;">Copyright © SuperYong Blog 2023</span>
        </div>
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import BlogHeader from "@/components/blogHeader.vue";
import NavBarModule from "@/components/navBarModule.vue";

export default {
  name: "basePage",
  components: {NavBarModule, BlogHeader},
  props: {
    showRight: {
      type: Boolean,
      default: true,
    },
    headerImgUrl: {
      type: String,
      default: require("@/assets/post-bg-desk.jpg")
    }
  },
  data() {
    return {}
  },
  computed: {
    backgroundImageStyle() {
      return {
        height: "420px",
        padding: "0",
        backgroundImage: `url(${this.headerImgUrl})`
      }
    }
  },
}
</script>

<style scoped>
.basePage {
  width: 100%;
  height: 100%;
}

.mainContainer {
  height: 100%;
  margin: 0 auto;
  display: flex;
  flex-direction: row;
}

@media (max-width: 576px) {
  .mainContainer {
    width: 95%;
  }

  .leftContainer {
    margin: 0 20px;
  }
}

@media (min-width: 576px) and (max-width: 768px) {
  .mainContainer {
    width: 90%;
  }

  .leftContainer {
    margin: 0 20px;
  }
}

@media (min-width: 768px) and (max-width: 992px) {
  .mainContainer {
    width: 90%;
  }

  .leftContainer {
    margin-right: 40px;
  }
}

@media (min-width: 992px) and (max-width: 1200px) {
  .mainContainer {
    width: 80%;
  }

  .leftContainer {
    margin-right: 40px;
  }
}

@media (min-width: 1200px) {
  .mainContainer {
    width: 75%;
  }

  .leftContainer {
    margin-right: 40px;
  }
}

.leftContainer {
  flex: 8;
  width: 70%;
  display: flex;
  flex-direction: column;
}

.rightContainer {
  flex: 2;
  width: 30%;
  display: flex;
  flex-direction: column;
}

@media (max-width: 768px) {
  .rightContainer {
    display: none !important;
  }
}


.blogFooter {
  height: 200px;
  font-size: 14px;
  text-align: center;
  padding-bottom: 30px;
}

</style>