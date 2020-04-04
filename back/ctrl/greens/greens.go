package greens

import (
	table "OrderMeal/model"
	model "OrderMeal/model/greens"
	"OrderMeal/pub"
	"OrderMeal/data/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 菜单添加，更新
func AddEdit(c *gin.Context) {
	var json table.Greens
	id := pub.StrToNum(c.Query("id"))
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, pub.Fail("参数错误"))
		return
	}

	// 若Id==0则添加
	json.UpdatedAt = pub.StrToNum(pub.GetSecond())
	var affect int64
	if id == 0 {
		json.ShopId = pub.GetRedisShopId(c)
		json.CreatedAt = pub.StrToNum(pub.GetSecond())

		affect = model.AddMenu(&json)
		if affect == 0 {
			c.JSON(http.StatusBadRequest, pub.Fail("添加失败"))
			return
		}
		c.JSON(http.StatusOK, pub.Success("添加成功"))
		return
	}
	// 否则更新
	affect = model.EditMenu(id, &json)
	if affect == 0 {
		c.JSON(http.StatusBadRequest, pub.Fail("编辑失败"))
		return
	}
	c.JSON(http.StatusOK, pub.Success("编辑成功"))
}

// 分页列表
func List(c *gin.Context) {
	var ps = new(pub.PageSize)
	if err := c.ShouldBindJSON(&ps); err != nil {
		log.Debug(err)
		c.JSON(http.StatusBadRequest, pub.Fail("参数错误"))
		return
	}
	ps.ShopId = pub.GetRedisShopId(c)
	var list []model.Menus
	list, count := model.MenuPage(ps)
	c.JSON(http.StatusOK, pub.Success(gin.H{"list": list, "count": count}))
}

// 软删除
func Del(c *gin.Context) {
	sid := pub.StrToNum(c.Query("id"))
	effect := model.DelMenu(sid)
	if effect == 0 {
		c.JSON(http.StatusOK, pub.Fail("删除失败"))
		return
	}
	c.JSON(http.StatusOK, pub.Success("删除成功"))
}
