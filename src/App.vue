<template>
  <div id="app">
    <img src="./assets/logo.png">
    <Menu/>
    <router-view/>
  </div>
</template>

<script>
import Menu from './components/Menu.vue'
import store from './store/index.js'

export default {
  name: 'app',
  components: {
    Menu,
  },
  store,
  created() {
    this.$store.dispatch('getPosts')
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
