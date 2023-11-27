<template>
  <div class="header-menu" :style="{backgroundColor: backgroundColor}">
    <div class="header-menu-left color-trans" :style="{color: navTextColor}">
      SuperYong
    </div>
    <div class="menu-icon" :style="{color: navTextColor}" @click="toggleNav" v-if="isMobile">
      <i class="el-icon-menu"></i>
    </div>
    <transition name="vertical-nav">
      <div class="vertical-nav" v-show="showNav">
        <div class="nav-header">
          <div style="font-size: 18px; font-weight: bold">SuperYong</div>
          <i class="el-icon-close" style="{color: #d6d6d6}" @click="toggleNav"></i>
        </div>
        <div class="spacer"></div>
        <router-link
            v-for="menu in menus"
            :key="menu.name"
            :to="menu.to"
            @click.native="toggleNav"
            :class="{isCurrentMenu: currentMenu === menu.to}"
            class="navMenu"
        >{{ menu.name }}</router-link>
        <div class="spacer"></div>
      </div>
    </transition>
    <div class="header-menu-right" v-if="!isMobile">
      <router-link
          v-for="menu in menus"
          :key="menu.name"
          :to="menu.to"
          :style="{color: navTextColor}"
          class="color-trans"
      >{{ menu.name}}</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: "navBarModule",
  data() {
    return {
      navTextColor: "white",
      backgroundColor: "transparent",
      mobileWidth: 576,
      isMobile: window.innerWidth < this.mobileWidth,
      showMobileNav: false,
      showNav: false,
      currentMenu: this.$route.path,
      menus: [
        {to: '/', name: 'HOME'},
        {to: '/about', name: 'ABOUT'},
        {to: '/tag', name: 'TAGS'},
      ],
    }
  },
  watch: {
    $route(to, from) {
      this.currentMenu = to.path;
    }
  },
  methods: {
    changeColor() {
      if (window.scrollY > 400) {
        this.backgroundColor = 'rgb(255,255,255, 0.8)';
        this.navTextColor = 'black'
      } else {
        this.backgroundColor = 'transparent';
        this.navTextColor = 'white'
      }
    },
    handleResize() {
      this.isMobile = window.innerWidth < this.mobileWidth;
      console.log(this.isMobile)
    },
    toggleNav() {
      this.showNav = !this.showNav;

      if (this.showNav) {
        document.body.classList.add('no-scroll');
      } else {
        document.body.classList.remove('no-scroll');
      }
    },
  },
  mounted() {
    window.addEventListener('scroll', this.changeColor);
    window.addEventListener('resize', this.handleResize);
    this.handleResize();

    if (window.innerWidth < this.mobileWidth) {
      document.querySelectorAll('.header-menu-right').forEach(function (el) {
        el.classList.add('no-animation');
      });
    }
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.changeColor);
    window.removeEventListener('resize', this.handleResize);
  }
}
</script>

<style scoped>
.header-menu {
  display: flex;
  justify-content: space-between;
  width: 100%;
  height: 60px;
  position: fixed;
  z-index: 1000;
  top: 0;
  left: 0;
  right: 0;
  transition: background-color 0.5s;
}

.header-menu-left {
  color: white;
  margin: 10px 0 0 20px;
  font-size: 30px;
  font-weight: bold;
  user-select: none;
  display: flex;
}

.color-trans {
  transition: color 0.5s;
}

.header-menu-right {
  display: flex;
  flex-direction: row;
}

@keyframes slideRightToHide {
  0% {
    clip-path: inset(0 0 0 0);
  }
  100% {
    clip-path: inset(0 0 0 100%);
  }
}

@keyframes slideLeftToAppear {
  from {
    clip-path: inset(0 0 0 100%);
  }
  to {
    clip-path: inset(0 0 0 0);
  }
}

.header-menu-right a {
  text-decoration: none;
  color: white;
  font-size: 12px;
  font-weight: 900;
  margin: 20px 20px 20px 20px;
}

.header-menu-right a:hover {
  color: #d2cdcd;
}

@media (max-width: 576px) {
  .header-menu-right {
    animation: slideRightToHide 0.8s ease forwards;
  }

  .no-animation {
    animation: none;
    clip-path: inset(0 100% 0 0);
  }
}

@media (min-width: 576px) {
  .header-menu-right {
    animation: slideLeftToAppear 0.8s ease forwards;
  }
}

.menu-icon {
  display: none;
}

.vertical-nav {
  display: none;
}

@media (max-width: 576px) {
  .menu-icon {
    display: flex;
    align-items: center;
    margin: 10px 20px 0 0;
  }

  .vertical-nav {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: center;
    position: fixed;
    top: 0;
    right: 0;
    width: 60%;
    height: 100vh;
    background: #fff;
    z-index: 1001;
  }
}

.spacer {
  flex-grow: 1;
}

.nav-header {
  width: 80%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  margin: 20px 20px 0 20px;
}

.navMenu {
  text-decoration: none;
  width: 60%;
  color: black;
  font-size: 13px;
  font-weight: 900;
  padding: 20px 20px 20px 20px;
}

.isCurrentMenu {
  border: 1px solid var(--solidLineColor);
  border-radius: 50px;
}

@keyframes slideInFromRight {
  from {
    transform: translateX(100%);
  }
  to {
    transform: translateX(0);
  }
}

@keyframes slideOutToRight {
  from {
    transform: translateX(0);
  }
  to {
    transform: translateX(100%);
  }
}

.vertical-nav-enter-active, .vertical-nav-leave-active {
  animation-duration: 0.5s;
}

.vertical-nav-enter, .vertical-nav-leave-to
{
  transform: translateX(100%);
}

.vertical-nav-enter-active {
  animation-name: slideInFromRight;
}

.vertical-nav-leave-active {
  animation-name: slideOutToRight;
}

</style>