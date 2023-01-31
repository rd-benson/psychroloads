<script setup lang="ts">
import router from '@/router/router';
import { RouterView } from 'vue-router';
import Item from './components/Item.vue';

document.body.addEventListener("click", function (event) {
  event.preventDefault();
});

</script>

<script lang="ts">

const routes: string[] = router.getRoutes().map((route) => {
  return route.name as string
})

</script>

<template>
  <!-- Menu -->
  <div class="menu-container">
    <div class="menu">
      <Item v-for="route in routes" :name="route" />
    </div>
  </div>
  <!-- Page -->
  <div class="view">
    <router-view />
  </div>
</template>

<style lang="scss">
@import url("@/assets/css/reset.css");
@import url("@/assets/css/font.css");
@import url("@/assets/css/colour.css");
@import url("@/assets/css/app.scss");

body {
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 0;
  font-family: "Cormorant";
  background-color: transparent;
}

#app {
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: var(--light1);
  overflow: hidden;
}

.menu-container {
  position: absolute;
  margin: 0;
  width: var(--menu-width);
  height: 100%;
}
.menu {
  position: absolute;
  top: 0;
  left: 0;
  min-width: var(--menu-width);
  height: 100%;
  padding: 10px 10px;
  background-color: var(--light);
  box-shadow: 0px 0px 5px var(--shadow);
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  a {
    color: var(--dark2);
    text-align: center;
    text-decoration: none;
    font-size: 12pt;
    font-family: 'Cormorant Italic';
    white-space: nowrap;

    &.router-link-exact-active {
        font-family: 'Cormorant Bold Italic';
    }
}
}

.view {
  position: absolute;
  margin: 0;
  margin-left: var(--menu-width);
  width: calc(100vw - var(--menu-width));
}
</style>
