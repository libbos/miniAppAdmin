<template>
  <el-row>
    <el-col :span="24" class="ov main-row-col-head">
      <el-button type="primary" size="small" icon="el-icon-plus" class="fl" @click="add">添加</el-button>
      <el-button size="small" icon="el-icon-refresh" class="fr" @click="getList">刷新</el-button>
    </el-col>
    <el-col :span="24">
      <el-button-group class="mr mb" v-for="c in list" :key="c.id">
        <el-button type="success" plain @click="edit(c.id)">{{c.name}}</el-button>
        <el-button type="danger" plain icon="el-icon-delete" @click="del(c.id)"></el-button>
      </el-button-group>
    </el-col>
  </el-row>
</template>

<script>
import L from '@/assets/js/L.js'
export default {
  data(){
    return {
      list: [],

    }
  },
  methods: {
    add(){
      let that = this;
      L.prompt("添加", "请输入菜系名称", (val) => {
        L.post("category/addEdit", {name: val}, (res)=>{
          L.topMsg(res, "success")
          that.getList()
        })
      })
    },
    edit(id){
      let that = this;
      L.prompt("编辑", "请输入菜系名称", (val) => {
        L.post("category/addEdit", {name: val, id: id}, (res)=>{
          L.topMsg(res, "success")
          that.getList()
        })
      })
    },
    del(id){
      let that = this;
      L.confirm("删除", "确定删除当前菜系，且当前菜系下所有菜品都将被删除。是否继续？", (val)=>{
        L.get("category/del", {id: id}, (res)=>{
          L.topMsg(res, "success")
          that.getList()
        })
      })
    },
    getList(){
      L.post("category/list", null, res=>{
        this.list = res
      })
    },
  },
  mounted(){
    this.getList()
  }
}
</script>