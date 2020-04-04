package shop

import (
	table "OrderMeal/model"
	model "OrderMeal/model/shop"
	"OrderMeal/pub"
	"OrderMeal/data/log"
	"net/http"
	// "OrderMeal/data/qiniu"
	// "strconv"
	"github.com/gin-gonic/gin"
)

// 门店添加
func Add(c *gin.Context) {
	var json table.Shop
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, pub.Fail("参数错误"))
		return
	}
	af := model.Add(json)
	if af == 0 {
		c.JSON(http.StatusBadRequest, pub.Fail("添加失败"))
		return
	}
	c.JSON(http.StatusOK, pub.Success("添加成功"))
}


// 删除餐桌
func DelTable(c *gin.Context) {
	id := pub.StrToNum(c.Query("id"))
	result := model.DelTable(id)
	if result {
		c.JSON(http.StatusOK, pub.Success("删除成功"))
		return
	}
	c.JSON(http.StatusBadRequest, pub.Fail("删除失败"))
}

// 餐桌添加、编辑
func AddTable(c *gin.Context) {
	var json table.Table
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Error("餐桌添加错误： ", err)
		c.JSON(http.StatusBadRequest, pub.Fail("参数错误"))
		return
	}
	json.UpdatedAt = pub.StrToNum(pub.GetSecond())
	var result bool
	if json.Id == 0 {
		json.ShopId = pub.GetRedisShopId(c)
		json.CreatedAt = pub.StrToNum(pub.GetSecond())
		result = model.AddTable(&json)
		if result {
			c.JSON(http.StatusOK, pub.Success("添加成功"))
			return
		}
		c.JSON(http.StatusBadRequest, pub.Fail("添加失败"))
		return
	}
	result = model.EditTable(&json)
	if result {
		c.JSON(http.StatusOK, pub.Success("编辑成功"))
		return
	}
	c.JSON(http.StatusBadRequest, pub.Fail("编辑失败"))
}

// 餐桌列表
func ListTable(c *gin.Context) {
	ps := new(pub.PageSize)
	ps.ShopId = pub.GetRedisShopId(c)
	result := make([]table.Table, 0)
	result = model.ListTable(ps)
	c.JSON(http.StatusOK, pub.Success(result))
}