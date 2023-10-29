<template>
  <div class="header-container">
    <div class="header-content">
      <div class="header-menu">
        <div class="header-menu-left">
          SuperYong
        </div>
        <div class="header-menu-right">
          <el-menu router mode="horizontal" text-color="#303133" active-text-color="#fff">
            <el-menu-item index="/home">Home</el-menu-item>
            <el-menu-item index="/about">About</el-menu-item>
            <el-menu-item index="/tags">Contact</el-menu-item>
          </el-menu>
        </div>
      </div>
      <div class="header-text">
        <slot name="headerImgInnerShow">
          <h1>Welcome to My Blog</h1>
        </slot>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "blogHeader",
  data() {
    return {
      headerImgUrl: require("@/assets/post-bg-desk.jpg")
    }
  }
}
</script>

<style scoped>
.header-container {
  background-image: url('@/assets/post-bg-desk.jpg');
  background-size: cover;
  background-position: center;
  color: #333;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
}

.el-menu {
  background-color: transparent !important;
  border: 0 !important;
}

.el-menu-item {
  font-weight: bold;
}

.header-menu-left .el-menu-item {
  font-size: 23px;
}

.el-menu-item:hover {
  background-color: transparent !important;
}

.el-menu-item.is-active {
  background-color: transparent !important;
}

.header-menu {
  display: flex;
  justify-content: space-between;
  width: 100%;
  height: 15%;
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
}

.header-content {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  width: 100%;
  height: 100%;
}

.header-text {
  color: #fff;
  height: 85%;
  font-size: 32px;
  user-select: none;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
}
</style>