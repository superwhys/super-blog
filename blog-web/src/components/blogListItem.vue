<template>
  <div class="blogItem">
    <a :href="toEndPoint">
      <h2 class="postTitle">{{ title }}</h2>
      <h3 class="postSubtitle">{{ subTitle }}</h3>
      <div class="postContentPreview">{{ Content }}</div>
    </a>
    <p class="postMeta">{{ Meta }}</p>
  </div>
</template>

<script>
export default {
  name: "blogListItem",
  props: {
    toEndPoint: {
      type: String,
      default: "#"
    },
    title: String,
    subTitle: String,
    Meta: String,
    Content: String,
  },
  data() {
    return {}
  },
}
</script>

<style scoped>
.blogItem {
  display: flex;
  flex-direction: column;
  padding: 20px 0;
  border-bottom: 1px solid var(--solidLineColor);
}

.blogItem a {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  text-decoration: none;
  color: gray;
}

.postTitle {
  margin: 0;
  color: black;
}

.postTitle:hover, .postTitle:hover ~ .postSubtitle, .postSubtitle:hover, .postSubtitle:hover ~ .postTitle {
  color: var(--hightlight);
}

.postSubtitle {
  margin: 10px 0;
  font-size: 16px;
  font-weight: normal;
}

.postContentPreview {
  text-align: left;
  font-size: 15px;
  font-style: italic;
}

.postMeta {
  display: flex;
  color: grey;
  font-size: 18px;
  margin: 10px 0;
  justify-content: flex-start;
}
</style>