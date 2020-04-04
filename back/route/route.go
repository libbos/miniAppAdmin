package route

import (
	"OrderMeal/ctrl/greens"
	"OrderMeal/ctrl/shop"
	"OrderMeal/ctrl/common"
	"OrderMeal/pub"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine{
	r := gin.Default()
	// 跨域中间件
	r.Use(pub.Cors())
	v1 := r.Group("/v1")
	{
		// 菜品 greens
		v1.POST("/menu/addEdit", greens.AddEdit)
		v1.POST("/menu/list", greens.List)
		v1.GET("/menu/del", greens.Del)

		// 菜系 greens
		v1.POST("/category/addEdit", greens.AddCategory)
		v1.POST("/category/list", greens.ListCategory)
		v1.GET("/category/del", greens.DelCategory)

		// 门店 shop
		v1.POST("/shop/addEdit", shop.Add)

		// 餐桌 table
		v1.POST("/table/addEdit", shop.AddTable)
		v1.POST("/table/list", shop.ListTable)
		v1.GET("/table/del", shop.DelTable)

		// 公共方法
		v1.POST("/uploadImage", common.UploadFile)
	}

	return r
}
