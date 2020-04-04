<template>
  <div id="app">
    <el-container>
      <el-header>
        <img class="fl header-left-logo" src="@/assets/img/sys-logo.png">
        <div class="header-right-info fr">
          <img class="header-default-user fr" src="@/assets/img/default-user.jpg">
        </div>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <el-menu default-active="1" @select="pointMenu">
            <!-- 门店管理 -->
            <el-submenu index="1">
              <template slot="title">
                <i class="el-icon-s-home"></i>
                <span>门店管理</span>
              </template>
              <el-menu-item index="shop"><i class="el-icon-s-shop"></i>门店列表</el-menu-item>
            </el-submenu>
            <!-- 店面装修 -->
            <el-submenu index="2">
              <template slot="title">
                <i class="el-icon-dish"></i>
                <span>店面装修</span>
              </template>
              <el-menu-item index="canzhuo"><i class="el-icon-fork-spoon"></i>餐桌设置</el-menu-item>
            </el-submenu>
            <!-- 菜单管理 -->
            <el-submenu index="3">
              <template slot="title">
                <i class="el-icon-fork-spoon"></i>
                <span>美食系列</span>
              </template>
              <el-menu-item index="menu"><i class="el-icon-tableware"></i>菜单列表</el-menu-item>
              <el-menu-item index="category"><i class="el-icon-food"></i>菜系列表</el-menu-item>
            </el-submenu>
            <!-- 点餐管理 -->
            <el-submenu index="4">
              <template slot="title">
                <i class="el-icon-s-grid"></i>
                <span>点餐管理</span>
              </template>
              <el-menu-item index="orderRecords"><i class="el-icon-tickets"></i>点餐记录</el-menu-item>
            </el-submenu>

          </el-menu>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import L from '@/assets/js/L.js'
export default {
  data(){
    return {}
  },
  methods: {
    pointMenu(index){
      this.$router.push({name: index})
    }
  },
  created(){
    this.$router.push({name: "shouye"})
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
#app .el-header{
  background: #2f333e;
  text-align: left;
}
#app .el-aside{
  background: #e2e9ff;
}
#app .el-main{
  padding: 10px;
}
#app .el-container{
  height: 100vh;
}
.el-menu{
  border-right: 0px !important;
  background: none !important;
}
.header-left-logo{
  display: block;
  height: 30px;
  margin-top: 15px;
}
.header-right-info{
  width: 40px;
}
.header-default-user{
  display: block;
  width: 40px;
  height: 40px;
  margin-top: 10px;
  border-radius: 50%;
}

</style>
