import ui from "element-ui"
import Vue from "vue"

class L{
  constructor(){
    this.REMOTE = "http://localhost:8888/v1/"
    this.iframeID = []
  }
  // POST请求
  async post(uri, data, cb){
    let tk = this.getLocalStory('tk')
    if (!tk) {
      this.topMsg("Token不存在", "error")
      return false
    }
    let options = {
      lock: true,
      text: "数据加载中 ...",
      spinner: "el-icon-loading",
      background: "rgba(255,255,255,0)"
    }
    let loadingInstance = ui.Loading.service(options);
    let headers = {
      'Content-Type': 'application/json',
      'Token': tk
    }
    return await new Promise((res, rej)=>{
      Vue.axios.post(this.REMOTE + uri, data, {headers: headers})
      .then(e => e.data)
      .then(e => {
        if (e.code != 0){
          this.topMsg(e.msg, 'error')
        }else{
          cb(e.data)
        }
      })
      .catch(e => rej(e))
      loadingInstance.close();
    })
  }
  // GET请求
  async get(uri, data, cb){
    return await new Promise((res, rej)=>{
      Vue.axios.get(this.REMOTE +  uri, {params: data})
      .then(e => e.data)
      .then(e => {
        if (e.code != 0){
          this.topMsg(e.msg, 'error')
        }else{
          cb(e.data)
        }
      })
    })
  }

  // 顶部提示信息
  topMsg(msg, tp){
    if (tp){
      ui.Message({
        message: msg,
        type: tp
      })
    }else{
      ui.Message(msg)
    }
  }

  // 确认提示框
  confirm(title, msg, cb){
    ui.MessageBox.confirm(msg, title, {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    }).then(() => {cb()}).catch(e => {})
  }

  // 提交弹框
  prompt(title, msg, cb){
    let self
    ui.MessageBox.prompt(msg, title, {
      confirmButtonText: '确定',
      cancelButtonText: '取消'
    }).then(({ value }) => {
      cb(value)
    }).catch(() => {
      ui.Message({message: '已取消'});       
    });
  }

  // 格式化时间戳
  ftime(tm, tp){
    let t = new Date(tm*1000)
    let yy = t.getFullYear()
    let mm = t.getMonth() < 9 ? "0"+(t.getMonth()+1) : t.getMonth()+1;
    let dd = t.getDate() < 10 ? "0"+t.getDate() : t.getDate();
    let hh = t.getHours() < 10 ? "0"+t.getHours() : t.getHours();
    let ii = t.getMinutes() < 10 ? "0"+t.getMinutes() : t.getMinutes();
    let ss = t.getSeconds() < 10 ? "0"+t.getSeconds() : t.getSeconds();
    if (tp) {
      return yy+"-"+mm+"-"+dd+" "+hh+":"+ii+":"+ss;
    }
    return yy+"-"+mm+"-"+dd
  }

  // layer弹层
  // params：vue，组件，标题，宽，高，props参数，回调
  openIframe(self, component, title, data, width, height){
    let iframeID = self.$layer.iframe({
      content: {content: component, parent: self, data: data || null},
      area: [width+'px', height+'px'],
      title: title,
      scrollbar: false,
      shadeClose: false
    })
    this.iframeID.push(iframeID)
  }

  // 关闭layer弹层
  closeIframe(self){
    if (this.iframeID.length) {
      let ifr = this.iframeID.pop()
      self.$layer.close(ifr)
    } else {
      console.log('弹窗不存在')
    }
  }

  // 临时获取本地缓存
  getLocalStory(key){
    let item = localStorage.getItem(key)
    return item
  }

  // 临时设置本地缓存
  setLocalStory(key, val){
    if (typeof(val) == 'string'){
      localStorage.setItem(key, val)
    } else {
      localStorage.setItem(key, JSON.stringify(val))
    }
  }


}

export default new L()