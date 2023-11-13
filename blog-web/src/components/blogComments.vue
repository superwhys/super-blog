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
      id: this.parseString(decodeURIComponent(location.pathname)),
      title: this.parseString(decodeURIComponent(location.pathname)),
      distractionFreeMode: false,
    })

    gitalk.render('gitalk-container')
  },
  methods: {
    parseString(input) {
      const parts = input.split('/');

      const year = parts[3];
      const month = parts[4];
      const day = parts[5];
      const title = parts[6];

      return `${year}-${month}-${day}-${title}`;
    }
  }
}
</script>

<style scoped>

</style>