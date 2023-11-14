<template>
  <div class="aboutPage">
    <base-page>
      <template v-slot:blogHeaderImgInnerShow>
        <div class="flyInFromTop">
          <div class="title">About</div>
          <div class="subTitle">Hey, this is SuperYong</div>
        </div>
      </template>
      <template v-slot:blogLeftMainShow>
        <div class="flyInFromBottom">
          <div class="aphorism">
            <p>猝然临之而不惊</p>
            <p>无故加之而不怒</p>
          </div>
          <div class="content">
            <div class="greeting">
              <p>👋 Hi, 我是 <a href="https://github.com/superwhys">@superwhys</a>,  一只在深圳工作的 <span>Gopher</span></p>
            </div>
            <div>
              <p>工作之余, 还是一个台球🎱, 篮球🏀 爱好者.... 以及睡懒觉爱好者 (¦3[▓▓] </p>
            </div>
            <div>
              <p>这是我使用 Golang Packr2 和 Vue2 搭建的个人博客. (前后端代码并未开源), 欢迎大家前来深入探讨各种问题~</p>
            </div>
            <el-tooltip class="item" effect="dark" content="😇 It'll help the luck stick 😇" placement="top">
              <p>😇 我相信 -> 好运即将发生! 😇 </p>
            </el-tooltip>
            <h3>下面是我的一些常用技术栈:</h3>
            <p class="languagesTools">
              <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://vuejs.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/vuejs/vuejs-original-wordmark.svg" alt="vuejs" width="40" height="40"/> </a> <a href="https://www.python.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/python/python-original.svg" alt="python" width="40" height="40"/> </a> <a href="https://www.djangoproject.com/" target="_blank" rel="noreferrer"> <img src="https://cdn.worldvectorlogo.com/logos/django.svg" alt="django" width="40" height="40"/> </a> <a href="https://www.docker.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original-wordmark.svg" alt="docker" width="40" height="40"/> </a> <a href="https://www.elastic.co" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/elastic/elastic-icon.svg" alt="elasticsearch" width="40" height="40"/> </a> <a href="https://flask.palletsprojects.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/pocoo_flask/pocoo_flask-icon.svg" alt="flask" width="40" height="40"/> </a> <a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a> <a href="https://www.w3.org/html/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/html5/html5-original-wordmark.svg" alt="html5" width="40" height="40"/> </a> <a href="https://www.jenkins.io" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/jenkins/jenkins-icon.svg" alt="jenkins" width="40" height="40"/> </a> <a href="https://kafka.apache.org/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/apache_kafka/apache_kafka-icon.svg" alt="kafka" width="40" height="40"/> </a> <a href="https://kubernetes.io" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/kubernetes/kubernetes-icon.svg" alt="kubernetes" width="40" height="40"/> </a> <a href="https://www.linux.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" alt="linux" width="40" height="40"/> </a> <a href="https://www.mongodb.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mongodb/mongodb-original-wordmark.svg" alt="mongodb" width="40" height="40"/> </a> <a href="https://www.mysql.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/mysql/mysql-original-wordmark.svg" alt="mysql" width="40" height="40"/> </a> <a href="https://www.nginx.com" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/nginx/nginx-original.svg" alt="nginx" width="40" height="40"/> </a> <a href="https://nodejs.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/nodejs/nodejs-original-wordmark.svg" alt="nodejs" width="40" height="40"/> </a> <a href="https://pandas.pydata.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/2ae2a900d2f041da66e950e4d48052658d850630/icons/pandas/pandas-original.svg" alt="pandas" width="40" height="40"/> </a> <a href="https://postman.com" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/getpostman/getpostman-icon.svg" alt="postman" width="40" height="40"/> </a> <a href="https://redis.io" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/redis/redis-original-wordmark.svg" alt="redis" width="40" height="40"/> </a> <a href="https://www.tensorflow.org" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/tensorflow/tensorflow-icon.svg" alt="tensorflow" width="40" height="40"/> </a> <a href="https://webpack.js.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/d00d0969292a6569d45b06d3f350f463a0107b0d/icons/webpack/webpack-original-wordmark.svg" alt="webpack" width="40" height="40"/> </a>
            </p>
          </div>
        </div>
      </template>
      <template v-slot:blogRightMainShow>
        <tags-module class="flyInFromBottom"></tags-module>
        <about-me-module class="flyInFromBottom"></about-me-module>
      </template>
    </base-page>
  </div>
</template>

<script>
import BasePage from "@/components/basePage.vue";
import TagsModule from "@/components/tagsModule.vue";
import AboutMeModule from "@/components/aboutMeModule.vue";

export default {
  name: "aboutPage",
  components: {AboutMeModule, TagsModule, BasePage}
}
</script>

<style scoped>
.title {
  font-size: 50px;
  font-weight: bold;
}

.subTitle {
  font-size: 20px;
}

.aphorism {
  border-left: 5px solid #d6d6d6;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 10px 0;
}

.aphorism p {
  margin: 0 0 5px 15px;
  font-size: 17px;
  color: #727272;
}

.content {
  display: flex;
  margin-top: 20px;
  flex-direction: column;
  align-items: flex-start;
}

.greeting {
  text-align: start;
}

.greeting a, .greeting span{
  text-decoration: none;
  color: #333333;
  font-weight: bold;
  font-size: 19px;
}

.languagesTools {
  width: 80%;
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  flex-wrap: wrap;
}
</style>