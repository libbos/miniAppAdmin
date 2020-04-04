import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

// 解决两次访问相同路由地址报错
const originalPush = Router.prototype.push
Router.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

const router =  new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/main',
      name: 'index',
      component: () => import('./views/index.vue'),
      children: [
        { // 首页
          path: '/main/shouye',
          name: 'shouye',
          component: () => import('./views/shouye/shouye.vue')
        },
        { // 菜品
          path: '/main/menu_list',
          name: 'menu',
          component: () => import('./views/menu/menu_list.vue')
        },
        { // 菜系
          path: '/main/category_list',
          name: 'category',
          component: () => import('./views/menu/category_list.vue')
        },
        { // 门店
          path: '/main/shop_list',
          name: 'shop',
          component: () => import('./views/shop/shop_list.vue')
        },
        { // 餐桌
          path: '/main/table_list',
          name: 'canzhuo',
          component: () => import('./views/shop/canzhuo_list.vue')
        },
        { // 点餐记录
          path: '/main/order_list',
          name: 'orderRecords',
          component: () => import('./views/order/records_list.vue')
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('./views/login.vue')
    }
  ]
})

router.beforeEach((to, from, next)=>{
  // console.log(to)
  // console.log(from)
  next()
})

export default router