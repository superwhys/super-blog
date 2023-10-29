<template>
  <a class="tagItem" :href="toEndPoint">
    {{ text }}
  </a>
</template>

<script>
export default {
  name: "tagItem",
  props: {
    toEndPoint: {
      type: String,
      default: "#"
    },
    text: String
  }
}
</script>

<style scoped>
.tagItem {
  color: grey;
  font-size: 12px;
  padding: 7px;
  margin: 2px;
  line-height: 10px;
  border: 1px solid darkgrey;
  border-radius: 12px;
  text-decoration: none;
}

.tagItem:hover {
  color: var(--hightlight);
  border: 1px solid var(--hightlight);
}
</style>