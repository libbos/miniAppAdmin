<template>
    <el-row>
      <el-col :span="24" class="ov main-row-col-head">
        <el-button type="primary" size="small" icon="el-icon-plus" class="fl" @click="add">添加</el-button>
        <div class="fr">
          <el-input class="w2 mr" size="small" placeholder="请输入搜索内容" v-model="order.sear"
            @keyup.enter="getList">
            <el-button slot="append" icon="el-icon-search" @click="getList"></el-button>
          </el-input>
          <el-button size="small" icon="el-icon-refresh" @click="getList">刷新</el-button>
        </div>
      </el-col>
      <el-col :span="24">
        <el-table :data="list" stripe class="w table-h ov-y table-loading">
          <el-table-column prop="name" label="菜名"></el-table-column>
          <el-table-column prop="price" label="价格"></el-table-column>
          <el-table-column prop="cost" label="成本"></el-table-column>
          <el-table-column prop="cname" label="菜系"></el-table-column>
          <el-table-column prop="remarks" width="200" label="描述" show-overflow-tooltip></el-table-column>
          <el-table-column prop="putaway" label="上线">
            <template slot-scope="scope">
              <el-tag type="success" size="small" v-if="scope.row.putaway">是</el-tag>
              <el-tag type="warning" size="small" v-else>否</el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="created_at" width="160" label="添加时间">
            <template slot-scope="scope">
              {{fmtTime(scope.row.created_at)}}
            </template>
          </el-table-column>
          <el-table-column prop="updated_at" width="160" label="最近更新">
            <template slot-scope="scope">
              {{fmtTime(scope.row.updated_at)}}
            </template>
          </el-table-column>
          <el-table-column width="200" label="操作">
            <template slot-scope="scope">
              <el-button-group>
                <el-button size="mini" type="warning" icon="el-icon-edit" @click="edit(scope.row)">编辑</el-button>
                <el-button size="mini" type="danger" icon="el-icon-delete" @click="dele(scope.row.id)">删除</el-button>
              </el-button-group>
            </template>
          </el-table-column>
        </el-table>
        <div class="main-table-bottom">
          <el-pagination
            background
            @current-change="changeCurrent"
            :current-page.sync="order.page"
            layout="total, prev, pager, next"
            :total="count">
          </el-pagination>
        </div>
      </el-col>
    </el-row>
</template>

<script>
import L from "@/assets/js/L.js"
import tmp from "./menu_add.vue"
export default {
  data(){
    return {
      list: [],
      count: 0,
      order: {
        page: 1,
        size: 10,
        sear: ""
      }
    }
  },
  methods: {
    add(){
      let row = {
        putaway: true
      }
      L.openIframe(this, tmp, '菜单添加', {row: row}, 600, 500)
    },
    edit(row){
      L.openIframe(this, tmp, '菜单编辑', {row: row}, 600, 500)
    },
    dele(id){
      let that = this;
      L.confirm("删除", "确认删除当前菜单？", ()=>{
        L.get("menu/del", {id: id}, res=>{
          L.topMsg(res, "success")
          that.getList()
        })
      })
      
    },
    changeCurrent(val){
      this.order.page = val
      this.getList()
    },
    getList(){
      L.post("menu/list", this.order, res=>{
        this.list = res.list;
        this.count = res.count;
      })
    },
    fmtTime(tm){
      return L.ftime(tm, true)
    }
  },
  mounted(){
    this.getList()
  }
}
</script>

