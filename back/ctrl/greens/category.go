package greens

import (
	table "OrderMeal/model"
	model "OrderMeal/model/greens"
	"OrderMeal/pub"
	"OrderMeal/data/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* 菜系分类 --- controller */
// 添加、编辑
func AddCategory(c *gin.Context) {
	var json table.Category
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, pub.Fail("参数错误"))
		return
	}
	// 若json.Id == 0 则添加，否则修改
	json.UpdatedAt = pub.StrToNum(pub.GetSecond())
	var affect int64
	if json.Id == 0 {
		json.CreatedAt = pub.StrToNum(pub.GetSecond())
		json.ShopId = pub.GetRedisShopId(c)
		affect = model.AddCategory(&json)
		if affect == 0 {
			log.Debug("菜系添加错误")
			c.JSON(http.StatusBadRequest, pub.Fail("添加失败"))
			return
		}
		c.JSON(http.StatusOK, pub.Success("添加成功"))
		return
	}
	affect = model.EditCategory(&json)
	if affect == 0 {
		log.Debug("菜系编辑错误")
		c.JSON(http.StatusBadRequest, pub.Fail("编辑失败"))
		return
	}
	c.JSON(http.StatusOK, pub.Success("编辑成功"))
}

// 菜系列表
func ListCategory(c *gin.Context) {
	ps := new(pub.PageSize)
	ps.ShopId = pub.GetRedisShopId(c)
	result := model.ListCategory(ps)
	c.JSON(http.StatusOK, pub.Success(result))
}

// 删除菜系，且当前菜系下所有菜品也将被删除
func DelCategory(c *gin.Context) {
	id := pub.StrToNum(c.Query("id"))
	result := model.DelCategory(id)
	if result {
		c.JSON(http.StatusOK, pub.Success("删除成功"))
		return
	}
	c.JSON(http.StatusOK, pub.Fail("删除失败"))
}