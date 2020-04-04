package greens

import (
	"OrderMeal/data/db"
	"OrderMeal/model"
	"OrderMeal/pub"
	"OrderMeal/data/log"
	"fmt"
)

// 综合类型
type Menus struct {
	model.Greens 		`xorm:"extends"`
	Name 	string 		`json:"cname"`
}

// 所有菜单列表
func MenuAll(shopId string) []Menus {
	orm := db.GetOrm()
	var list []Menus
	err := orm.Sql("SELECT g.*,c.name as cname FROM greens AS g, category AS c ON g.category = c.id WHERE g.is_del <> 0 AND c.is_del <> 0 AND g.shop_id = "+shopId).Find(&list)
	if err != nil {
		log.Debug("SQL ERR: ", err)
		panic(err)
	}
	return list
}

// 分页查询列表
func MenuPage(ps *pub.PageSize) ([]Menus, int64) {
	orm := db.GetOrm()
	var list []Menus
	cond := `g.is_del = 0 AND (g.name like "%` + ps.Sear + `%" or g.remarks like "%` + ps.Sear + `%")`
	if ps.Sear == "" {
		cond = "g.is_del = 0"
	}
	condition := fmt.Sprintf("%s AND g.shop_id = %d", cond, ps.ShopId)

	count, err := orm.Table("greens").Alias("g").
		Join("INNER", []string{"category", "c"}, "g.category_id = c.id").
		Where(condition).Limit(ps.Size, (ps.Page-1)*(ps.Size)).
		Desc("g.id").FindAndCount(&list)

	if err != nil {
		log.Debug("SQL ERR: ", err)
		panic(err)
	}

	return list, count
}

// 菜单添加
func AddMenu(menu *model.Greens) int64 {
	orm := db.GetOrm()
	affect, err := orm.Table("greens").Insert(menu)
	if err != nil {
		log.Debug("SQL ERR: ", err)
		panic(err)
	}
	return affect
}

// 更新菜单
func EditMenu(id int64, menu *model.Greens) int64 {
	orm := db.GetOrm()
	// 根据当前菜单是否被删除进行更新
	affect, err := orm.Table("greens").ID(id).Cols("name","price","category_id","remarks","putaway","cost","picture","updated_at").Update(menu)
	if err != nil {
		log.Debug("SQL ERR: ", err)
		panic(err)
	}
	return affect
}

// 删除
func DelMenu(id int64) int64 {
	orm := db.GetOrm()
	m := new(model.Greens)
	m.IsDel = true
	m.UpdatedAt = pub.StrToNum(pub.GetSecond())
	effect, err := orm.Table("greens").ID(id).Cols("is_del","updated_at").Update(m)
	if err != nil {
		log.Debug("SQL ERR: ", err)
		panic(err)
	}
	return effect
}

/* 菜系分类 */
// 添加菜系
func AddCategory(data *model.Category) int64 {
	orm := db.GetOrm()
	affect, err := orm.Table("category").Insert(data)
	if err != nil {
		log.Error("SQL ERR: ", err)
		panic(err)
	}
	return affect
}
// 编辑菜系
func EditCategory(data *model.Category) int64 {
	orm := db.GetOrm()
	affect, err := orm.Table("category").ID(data.Id).Cols("name").Update(data)
	if err != nil {
		log.Error("SQL ERR: ", err)
		panic(err)
	}
	return affect
}
// 菜系列表
func ListCategory(ps *pub.PageSize) []model.Category {
	orm := db.GetOrm()
	result := make([]model.Category, 0)
	err := orm.Table("category").Where("shop_id = ? AND is_del = 0", ps.ShopId).Find(&result)
	if err != nil {
		log.Error("SQL ERR: ", err)
		panic(err)
	}
	return result
}

// 删除菜系
func DelCategory(id int64) bool {
	orm := db.GetOrm()
	session := orm.NewSession()
	defer session.Close()
	// 开启事务
	err := session.Begin()
	// 先获取菜系信息
	cate := new(model.Category)
	_, err = session.Table("category").ID(id).Get(cate)
	// 软删除菜品
	gre := new(model.Greens)
	gre.IsDel = true
	gre.UpdatedAt = pub.StrToNum(pub.GetSecond())
	_, err = session.Table("greens").Cols("is_del","updated_at").Where("category_id=? AND is_del=0", cate.Id).Update(gre)
	if err != nil {
		log.Debug("软删除菜品错误: ", err)
		session.Rollback()
		return false
	}
	// 软删除菜系
	cate.IsDel = true
	cate.UpdatedAt = pub.StrToNum(pub.GetSecond())
	_, err = session.Table("category").ID(id).Cols("is_del").Update(cate)
	if err != nil {
		log.Debug("软删除菜系错误: ", err)
		session.Rollback()
		return false
	}
	// 提交事务
	err = session.Commit()
	if err != nil {
		log.Debug("软删除菜系提交错误: ", err)
		session.Rollback()
		return false
	}
	return true
}
