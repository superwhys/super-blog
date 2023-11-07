<template>
  <div class="postPage">
    <base-page>
      <template v-slot:blogLeftMainShow>
        {{year}}-{{month}}-{{day}}-{{name}}
      </template>
    </base-page>
  </div>
</template>

<script>
import BasePage from "@/components/basePage.vue";

export default {
  name: "blogPage",
  components: {BasePage},
  data() {
    return {
      year: "",
      month:"",
      day:"",
      blogName: ""
    }
  },
  created() {
    console.log(this.$route.params.name)
    this.year = this.$route.params.year
    this.month = this.$route.params.month
    this.day = this.$route.params.day
    this.name = this.$route.params.name
  }
}
</script>

<style scoped>

</style>