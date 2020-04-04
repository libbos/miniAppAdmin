<template>
  <div class="login">
  <el-tabs stretch v-model="tabs" type="card"  class="login-card" @tab-click="handleClick">
    <el-tab-pane label="帐号登录" name="1" class="p">
      <el-input
        class="mb"
        placeholder="请输入登录帐号"
        prefix-icon="el-icon-user-solid"
        v-model="username">
      </el-input>

      <el-input
        class="mt mb"
        placeholder="请输入登录密码"
        type="password"
        prefix-icon="el-icon-s-opportunity"
        v-model="password">
      </el-input>

      <el-checkbox v-model="savePassword" class="mt mb b">记住密码</el-checkbox>

<el-button-group class="w mt">
  <el-button class="ws" icon="el-icon-refresh-left">重置</el-button>
  <el-button class="ws" icon="el-icon-s-promotion" type="primary" @click="login">登录</el-button>
</el-button-group>

    </el-tab-pane>
    <el-tab-pane label="微信登录" name="2">
      <img :src="wxImg" id="wxImg"/>
      <!-- <el-image :src="wxImg" id="wxImg"></el-image> -->
    </el-tab-pane>
  </el-tabs>
  </div>
</template>

<script>
export default {
  data(){
    return {
      username: "",
      password: "",
      tabs: '1',
      savePassword: true,
      wxImg: "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1578308409744&di=447fa4c72d99072e18eb0fa6baa06c35&imgtype=0&src=http%3A%2F%2Fs9.rr.itc.cn%2Fr%2FwapChange%2F20174_7_16%2Fa92hvd87222756647405.jpg"
    }
  },
  methods: {
    login(){
      this.$parent.goIndex()
    },
    handleClick(tab, env){
      console.log(tab)
      console.log(env)
    }
  }
}
</script>

<style scoped>
.login{
  width: 100vw;
  height: 100vh;
  background: url("../assets/img/bg-login.jpg") no-repeat;
  background-size: 100% 100%;
  position: absolute;
}
.login .login-card{
  width: 350px;
  height: 380px;
  position: relative;
  left: 50%;
  top: 20%;
  margin-left: -200px;
  background: rgba(250, 250, 250, .8);
  /* box-shadow: 2px -2px 2px #eee; */
}
#wxImg{
  width: 300px;
  height: 300px;
  position: relative;
  left: 50%;
  margin-left: -150px;
}
</style>