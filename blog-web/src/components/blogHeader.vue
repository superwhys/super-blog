<template>
  <div class="header-container">
    <div class="header-content">
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
    }
  }
}
</script>

<style scoped>
.header-container {
  overflow: visible;
  background-size: cover;
  background-position: center;
  color: #333;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
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