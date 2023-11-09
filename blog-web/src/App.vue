<template>
  <div id="app">
    <keep-alive exclude="blogPage">
      <router-view></router-view>
    </keep-alive>
  </div>
</template>

<script setup>
</script>

<style>
:root {
  --hightlight: cornflowerblue;
  --solidLineColor: #d6d6d6;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  height: 100%;
}
</style>