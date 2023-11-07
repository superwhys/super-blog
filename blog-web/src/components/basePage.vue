<template>
  <div class="basePage">
    <el-container style="min-height: 100vh;">
      <div class="header-menu" :style="{backgroundColor: backgroundColor}">
        <div class="header-menu-left" :style="{color: navTextColor}">
          SuperYong
        </div>
        <div class="header-menu-right">
          <router-link to="/home" :style="{color: navTextColor}">HOME</router-link>
          <router-link to="/about" :style="{color: navTextColor}">ABOUT</router-link>
          <router-link to="/tags" :style="{color: navTextColor}">TAGS</router-link>
        </div>
      </div>
      <el-header :style="backgroundImageStyle">
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
  components: {BlogHeader},
  props: {
    headerImgUrl: {
      type: String,
      default: require("@/assets/post-bg-desk.jpg")
    }
  },
  data() {
    return {
      navTextColor: "white",
      backgroundColor: "transparent"
    }
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
  mounted() {
    window.addEventListener('scroll', this.changeColor);
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.changeColor);
  },
  methods: {
    changeColor() {
      if (window.scrollY > 400) {
        this.backgroundColor = 'rgb(255,255,255, 0.8)';
        this.navTextColor = 'black'
      } else {
        this.backgroundColor = 'transparent';
        this.navTextColor = 'white'
      }
    },
  },
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


.header-menu {
  display: flex;
  justify-content: space-between;
  width: 100%;
  height: 60px;
  position: fixed;
  z-index: 1000;
  top: 0;
  left: 0;
  right: 0;
}

.header-menu-left {
  color: white;
  margin: 10px 0 0 20px;
  font-size: 30px;
  font-weight: bold;
  user-select: none;
  display: flex;
}

.header-menu-right {
  display: flex;
  flex-direction: row;
}

.header-menu-right a {
  text-decoration: none;
  color: white;
  font-size: 12px;
  font-weight: 900;
  margin: 20px 20px 20px 20px;
}

.header-menu-right a:hover {
  color: #d2cdcd;
}
</style>