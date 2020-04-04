<template>
  <el-row>
    <el-col :span="24" class="ov main-row-col-head">
      <el-button type="primary" size="small" icon="el-icon-plus" class="fl" @click="add">置办餐桌</el-button>
      <el-button size="small" icon="el-icon-refresh" class="fr" @click="getList">刷新</el-button>
    </el-col>
    <el-col :span="24">
      <el-card class="card-table fl mr ml mt" v-for="t in list" :key="t.id">
        <img src="@/assets/img/table.png" class="w b">
        <div class="ov mt">
          <el-tooltip :content="t.name" placement="top-start">
          <el-button type="text" class="fl w-txt ov" @click="edit(t)">{{t.name}}</el-button>
          </el-tooltip>
          <el-button type="text" class="fr red" icon="el-icon-delete" @click="del(t.id)"></el-button>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
import L from "@/assets/js/L.js"
export default {
  data(){
    return {
      list: [],
      tname: ""
    }
  },
  methods: {
    add(){
      let that = this;
      L.prompt("添加", "请输入餐桌编号或名称", (val)=>{
        L.post("table/addEdit", {name: val}, (res)=>{
          L.topMsg(res, "success")
          that.getList()
        })
      })
    },
    edit(row){
      let that = this;
      L.prompt("编辑", "请输入餐桌编号或名称", (val)=>{
        L.post("table/addEdit", {name: val, id: row.id}, (res)=>{
          L.topMsg(res, "success")
          that.getList()
        })
      })
    },
    del(id){
      let that = this;
      L.confirm("删除", "删除当前已置办餐桌，是否继续？", ()=>{
        L.get("table/del", {id: id}, (res)=>{
          L.topMsg(res, "success")
          that.getList()
        })
      })
    },
    getList(){
      L.post("table/list", null, (res)=>{
        this.list = res;
      })
    }
  },
  mounted(){
    this.getList()
  }
}
</script>

<style scope>
.card-table{
  width: 140px;
}
.w-txt{
  min-width: 50px;
  max-width: 90px;
}
.card-table .el-card__body{
  padding: 10px;
}
</style>