<template>
  <div id="app">
    <p>

    </p>
  </div>
</template>

<script>

import {getCaptcha} from '@/api/base'

export default {
  name: 'App',
  // data(){
  //   return{
  //     tt : process.env.VUE_APP_PORT
  //   }
  // },
  created() {
    getCaptcha();
  },
  components: {

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
