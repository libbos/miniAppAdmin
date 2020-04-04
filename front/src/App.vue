<template>
  <router-view></router-view>
</template>

<script>
import L from '@/assets/js/L.js'
export default {
  data(){
    return {}
  },
  methods: {
    goIndex(){
      L.setLocalStory('tk', "token")
      this.$router.push({name: 'index'})
    }
  },
  created(){
    let tk = L.getLocalStory("tk")
    if (!tk){
      this.$router.push({name: 'login'})
    } else {
      this.$router.push({name: 'index'})
    }
  }
}
</script>
