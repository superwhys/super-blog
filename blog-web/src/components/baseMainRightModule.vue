<template>
  <div class="baseRight">
    <router-link class="baseRightTitle" :to="toEndPoint">
      <slot name="rightModuleTitle"></slot>
    </router-link>
    <div class="baseRightMain">
      <slot name="rightModuleMain"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "baseMainRightModule",
  props: {
    toEndPoint: {
      type: String,
      default: "#"
    }
  }
}
</script>

<style scoped>
.baseRight {
  color: gray;
  border-top: 1px solid var(--solidLineColor);
  padding: 10px 0 20px 0;
  display: flex;
  flex-direction: column;
}

.baseRightTitle {
  color: gray;
  font-size: 14px;
  text-decoration: none;
  font-weight: bold;
  margin: 10px 0;
  display: flex;
  justify-content: flex-start;
}

.baseRightTitle:hover {
  color: var(--hightlight);
}

.baseRightMain {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}
</style>