<template>
  <div id="gitalk-container"></div>
</template>

<script>
export default {
  name: "blogComments",
  mounted() {
    const gitalk = new Gitalk({
      clientID: 'f10cc349c493295d54b8',
      clientSecret: 'c3512c2c2de6f4041cd394dfcc88292c7c673eec',
      repo: 'blog-posts',
      owner: 'superwhys',
      admin: ['superwhys'],
      id: decodeURIComponent(location.pathname),
      title: decodeURIComponent(location.pathname),
      distractionFreeMode: false,
      proxy: "https://cors-server-ecru.vercel.app/github_access_token",
    })

    gitalk.render('gitalk-container')
  }
}
</script>

<style scoped>

</style>