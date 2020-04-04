<template>
<el-row>
<el-col :span="12" class="mt width">
<el-form size="small" label-width="80px">
  <el-form-item label="菜名">
    <el-input v-model="info.name"></el-input>
  </el-form-item>
  <el-form-item label="分类">
    <el-select v-model="info.category_id" placeholder="请选择名菜分类">
      <el-option v-for="c in categoryList" :key="c.id" :label="c.name" :value="c.id"></el-option>
    </el-select>
  </el-form-item>
  <el-form-item label="价格">
    <el-input v-model="info.price"></el-input>
  </el-form-item>
  <el-form-item label="成本">
    <el-input v-model="info.cost"></el-input>
  </el-form-item>
  <el-form-item label="描述">
    <el-input type="textarea" v-model="info.remarks"></el-input>
  </el-form-item>
  <el-form-item label="状态">
    <el-radio-group v-model="info.putaway">
      <el-radio :label="true">上架</el-radio>
      <el-radio :label="false">下架</el-radio>
    </el-radio-group>
  </el-form-item>

  <el-form-item>
    <el-button size="small" type="danger" icon="el-icon-close" @click="close">取消</el-button>
    <el-button size="small" type="primary" icon="el-icon-check" @click="submit">提交</el-button>
  </el-form-item>
</el-form>
</el-col>
<el-col :span="12" class="mt width">
<el-form size="small" label-width="80px">
  <el-form-item label="图片">
    <el-upload
      class="avatar-uploader"
      :action="uploadPaht"
      :show-file-list="false"
      :on-success="handleAvatarSuccess"
      :before-upload="beforeAvatarUpload">
      <img v-if="shopImg" :src="shopImg" class="avatar">
      <i v-else class="el-icon-plus avatar-uploader-icon"></i>
    </el-upload>
  </el-form-item>
</el-form>
</el-col>
</el-row>
</template>

<script>
import L from '@/assets/js/L.js'
export default {
  data(){
    return {
      uploadPaht: '',
      shopImg: '',
      categoryList: [{name: "特色菜", id: 1}, {name: "必点菜", id: 2}],
      mid: 0,
      info: {
        name: "",
        category_id: null,
        price: 1,
        cost: 1,
        remarks: "",
        putaway: true,
        picture: '',
      },
    }
  },
  props: {
    row: {
      type: Object,
      default: ()=>{
        return {}
      }
    }
  },
  methods: {
    handleAvatarSuccess(res, file) {
      this.shopImg = URL.createObjectURL(file.raw);
      this.info.picture = res.data;
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === 'image/jpeg';
      const isLt2M = file.size / 1024 / 1024 < 2;

      if (!isJPG) {
        this.$message.error('上传头像图片只能是 JPG 格式!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isJPG && isLt2M;
    },
    close(){
      L.closeIframe(this)
    },
    submit(){
      this.info.price = parseFloat(this.info.price)
      this.info.cost = parseFloat(this.info.cost)
      if (this.info.name == "") {
        L.topMsg("未输入美食名称", "error")
        return false;
      } else if (!this.info.price > 0) {
        L.topMsg("价格不正确", "error")
        return false;
      } else if (!this.info.cost > 0) {
        L.topMsg("成本不正确", "error");
        return false;
      } else if (!this.info.category_id) {
        L.topMsg("未选择菜品分类", "error")
        return false;
      }
      L.post("menu/addEdit?id="+this.mid, this.info, res=>{
        L.topMsg(res, "success")
        // 子组件调父组件方法
        this.$parent.getList()
        this.close()
      })
    },
    getCategory(){
      L.post("category/list", null, res=>{
        this.categoryList = res
      })
    }
  },
  mounted(){
    this.mid = this.row.id;
    this.info.name = this.row.name;
    this.info.category_id = this.row.category_id;
    this.info.price = this.row.price;
    this.info.cost = this.row.cost;
    this.info.remarks = this.row.remarks;
    this.info.putaway = this.row.putaway;
    this.info.picture = this.row.picture;
  },
  created(){
    this.uploadPaht = L.REMOTE + "/uploadImage"
    this.getCategory()
  }
}
</script>