package shop

import (
	"OrderMeal/data/db"
	"OrderMeal/model"
	"OrderMeal/data/log"
	"OrderMeal/pub"
)



// 新开店面
func Add(shop model.Shop) int64 {
	o := db.GetOrm()
	affect, err := o.Insert(&shop)
	if err != nil {
		log.Error("insert err: ", err)
		return 0
	}
	return affect
}




// 餐桌添加
func AddTable(tab *model.Table) bool {
	orm := db.GetOrm()
	if _, err := orm.Insert(tab); err != nil {
		return false
	}
	return true
}

// 餐桌编辑
func EditTable(tab *model.Table) bool {
	orm := db.GetOrm()
	_, err := orm.ID(tab.Id).Cols("name", "updated_at").Update(tab)
	if err != nil {
		return false
	}
	return true
}

// 餐桌列表
func ListTable(ps *pub.PageSize) []model.Table {
	orm := db.GetOrm()
	list := make([]model.Table, 0)
	err := orm.Where("shop_id=? AND is_del=0", ps.ShopId).Find(&list)
	if err != nil {
		log.Error("获取餐桌错误： ", err)
		panic(err)
	}
	return list
}

// 餐桌软删除
func DelTable(id int64) bool {
	orm := db.GetOrm()
	row := new(model.Table)
	row.IsDel = true
	row.UpdatedAt = pub.StrToNum(pub.GetSecond())
	_, err := orm.ID(id).Cols("is_del", "updated_at").Update(row)
	if err != nil {
		return false
	}
	return true
}