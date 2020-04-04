<template>
  <el-row>
    <el-col :span="24" class="ov main-row-col-head">
      <el-date-picker
        class="fl"
        size="small"
        v-model="stime"
        type="date"
        placeholder="选择查看日期"
        format="yyyy 年 MM 月 dd 日"
        value-format="timestamp">
      </el-date-picker>
      <el-button size="small" type="primary" icon="el-icon-search" class="fl ml">高级搜索</el-button>

      <el-button size="small" icon="el-icon-refresh" class="fr" @click="getList">刷新</el-button>
    </el-col>
    <el-col :span="24">
      <el-table :data="list" stripe class="w table-h ov-y table-loading">
        <el-table-column prop="name" label="菜名"></el-table-column>
        <el-table-column prop="price" label="价格"></el-table-column>
      </el-table>
      <div class="main-table-bottom">
        <el-pagination
          background
          @size-change="changeSize"
          @current-change="changeCurrent"
          :current-page.sync="order.page"
          :page-sizes="[10, 20, 50, 100]"
          :page-size="order.size"
          layout="total, sizes, prev, pager, next"
          :total="count">
        </el-pagination>
      </div>
    </el-col>
  </el-row>
</template>

<script>
import L from "@/assets/js/L.js"
export default {
  data(){
    return {
      stime: '',
      list: [],
      count: 0,
      order: {
        page: 1,
        size: 20,
        sear: ""
      }
    }
  },
  methods: {
    getList(){

    },
    changeCurrent(val){
      this.order.page = val
      this.getList()
    },
    changeSize(val){
      this.order.size = val;
      this.getList()
    },
    fmtTime(tm){
      return L.ftime(tm, true)
    }
  },
  mounted(){
    
  }
}
</script>