<template>
  <div class="postPage">
    <base-page>
      <template v-slot:blogLeftMainShow>
        {{ postItem.postContent }}
      </template>
    </base-page>
  </div>
</template>

<script>
import BasePage from "@/components/basePage.vue";
import {getPost} from "@/network/postRequest";
import {BlogItem} from "@/models/blogItem";

export default {
  name: "blogPage",
  components: {BasePage},
  data() {
    return {
      year: "",
      month:"",
      day:"",
      name: "",
      postItem: new BlogItem({})
    }
  },
  created() {
    console.log(this.$route.params.name)
    this.year = this.$route.params.year
    this.month = this.$route.params.month
    this.day = this.$route.params.day
    this.name = this.$route.params.name

    getPost(this.year, this.month, this.day, this.name).then(resp => {
      console.debug(resp)
      this.postItem = new BlogItem(resp.data)
    })
  }
}
</script>

<style scoped>

</style>