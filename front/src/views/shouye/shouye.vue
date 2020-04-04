<template>
  <div class="shouye">
    <el-calendar v-model="dateValue"></el-calendar>
  </div>
</template>

<script>
export default {
  data(){
    return {
      dateValue: new Date()
    }
  }
}
</script>