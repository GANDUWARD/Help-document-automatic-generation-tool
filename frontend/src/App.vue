<template>
  <router-view></router-view>
</template>
<script>
import HelloWorld from './components/program.vue'
import { createRouter, createWebHistory } from 'vue-router'
const routes = [
    {path: '/', component:HelloWorld},
]

const router = createRouter({
    history: createWebHistory(),
    routes, 
})
export default {
  name: 'App',
  components: {

  },
  router,
}
</script>

<style>
#app {
  height: 100vh;
  width: 100vw;
  overflow: hidden; /* 防止滚动条出现，根据需要添加或删除此行 */
}
/* 如果需要让子组件占据100%高度，也需要设置子组件的高度为100% */
.router-view {
  height: 100%;
}
</style>
