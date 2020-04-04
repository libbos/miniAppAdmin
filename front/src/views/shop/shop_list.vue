<template>
  <el-row>
    <el-col :span="24" class="ov main-row-col-head">
      <el-button type="primary" size="small" icon="el-icon-plus" class="fl" @click="add">添加</el-button>
      <div class="fr">
        <el-input
          class="w2 mr"
          size="small"
          placeholder="请输入搜索内容"
          v-model="order.sear"
          @keyup.enter="getList">
          <el-button slot="append" icon="el-icon-search" @click="getList"></el-button>
        </el-input>
        <el-button size="small" icon="el-icon-refresh" @click="getList">刷新</el-button>
      </div>
    </el-col>
    <el-col :span="24">
      <el-table :data="list" stripe class="w table-h ov-y table-loading">
        <el-table-column prop="id" width="60" label="编号"></el-table-column>
        <el-table-column prop="name" label="菜名"></el-table-column>
        <el-table-column prop="price" label="价格"></el-table-column>
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
    <!-- 添加、编辑 -->
    <el-col :span="24">
<el-dialog title="门店操作" width="600px" :close-on-click-modal="false" :visible.sync="showForm">
  <el-form ref="form" v-model="info" label-width="80px">
    <el-form-item label="门店名称">
      <el-input size="medium" placeholder="请输入门店名称" v-model="info.name"></el-input>
    </el-form-item>
    <el-form-item label="联系电话">
      <el-input size="medium" placeholder="座机或手机号" v-model="info.phone"></el-input>
    </el-form-item>
    <el-form-item label="所在地址">
      <el-input size="medium" type="textarea" v-model="info.address"></el-input>
    </el-form-item>
    <el-form-item label="门店描述">
      <el-input size="medium" type="textarea" v-model="info.remarks"></el-input>
    </el-form-item>
    <el-form-item label="有效期至">
      <el-date-picker size="medium" v-model="info.valid" type="date" placeholder="有效期截至到">
      </el-date-picker>
    </el-form-item>
  </el-form>



  <div slot="footer">
    <el-button size="small" type="danger" icon="el-icon-close" @click="showForm = false">取消</el-button>
    <el-button size="small" type="primary" icon="el-icon-check" @click="addSubmit">提交</el-button>
  </div>
</el-dialog>
    </el-col>
  </el-row>
</template>

<script>
import L from "@/assets/js/L.js"
export default {
  data(){
    return {
      showForm: false,
      list: [],
      count: 0,
      order: {
        page: 1,
        size: 10,
        sear: ""
      },
      info: {
        name: '',
        address: '',
        phone: '',
        remarks: '',
        picture: [],
        logo: '',
        valid: ''
      }
    }
  },
  methods: {
    add(){
      this.showForm = true
    },
    addSubmit(){

    },
    getList(){

    },
    changeCurrent(val){
      this.order.page = val
      this.getList()
    },
    ftime(tm){
      return L.ftime(tm, true)
    }
  },
  mounted(){

  }
}
</script>

