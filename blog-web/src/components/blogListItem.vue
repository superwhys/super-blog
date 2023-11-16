<template>
  <div class="blogItem">
    <router-link :to="'/post'+toEndPoint">
      <h2 class="postTitle oneLine">{{ title }}</h2>
      <h3 class="postSubtitle oneLine">{{ subTitle }}</h3>
      <div class="postContentPreview">{{ postContent }}</div>
    </router-link>
    <p class="postMeta">{{ metaData }}</p>
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
    metaData: String,
    postContent: String,
  },
  data() {
    return {}
  },
}
</script>

<style scoped>
.blogItem {
  width: 100%;
  display: flex;
  flex-direction: column;
  padding: 20px 0;
  border-bottom: 1px solid var(--solidLineColor);
}

.blogItem a {
  display: flex;
  flex-direction: column;
  text-decoration: none;
  color: gray;
}

.postTitle {
  margin: 0;
  color: black;
  text-align: left;
}

.oneLine {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  width: 100%;
}

.postTitle:hover, .postTitle:hover ~ .postSubtitle, .postSubtitle:hover, .postSubtitle:hover ~ .postTitle {
  color: var(--hightlight);
}

.postSubtitle {
  margin: 10px 0;
  font-size: 16px;
  font-weight: normal;
  text-align: left;
}

.postContentPreview {
  text-align: left;
  font-size: 15px;
  font-style: italic;

  max-height: 150px;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 3;
  overflow: hidden;
}

.postMeta {
  display: flex;
  color: grey;
  font-size: 18px;
  margin: 10px 0;
  justify-content: flex-start;
}
</style>