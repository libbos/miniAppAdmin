package model

/*
* 1、此文件为数据库数据表
* 2、符号（-：不映射，->：只写不读，<-只读不写）
*/

// 菜系
type Category struct {
	Id 					int64	 		`json:"id", xorm:"pk autoincr"`
	ShopId 			int64	 		`json:"shop_id", xorm:"->"`
	Name 				string 		`json:"name"`
	IsDel 			bool 			`json:"is_del", xorm:"notnull default 0 ->"`
	CreatedAt 	int64			`json:"created_at"`
	UpdatedAt 	int64			`json:"updated_at"`
}


// 菜品类型
type Greens struct {
	Id       		int64	   		`json:"id", xorm:"pk autoincr"`
	ShopId   		int64	   		`json:"shop_id", xorm:"->"`                                 // 店id
	Name     		string  		`json:"name", binding:"required", xorm:"unique(shop_id)"` // 菜单名称
	Price    		float64 		`json:"price", binding:"required"`               // 价格
	CategoryId 	int64	   		`json:"category_id", xorm:"index"`                        // 分类
	Remarks  		string  		`json:"remarks"`                                 // 描述
	Putaway  		bool    		`json:"putaway"`                                 // 是否上架
	Cost     		float64 		`json:"cost"`                                    // 成本
	Picture  		string  		`json:"picture"`                                 // 图片地址
	IsDel    		bool    		`json:"is_del", xorm:"notnull default 0 ->"`  	// 是否删除，软删除
	CreatedAt 	int64			  `json:"created_at"`       			// 创建时间
	UpdatedAt 	int64			  `json:"updated_at"`       			// 更新时间
}

// 门店
type Shop struct {
	Id       		int64	 		`json:"id", xorm:"pk autoincr"`
	Name     		string 		`json:"name"`                   							// 店名
	Address  		string 		`json:"address", xorm:"unique"` 							// 地址
	Picture  		string 		`json:"picture"`                							// 门店图片
	Logo     		string 		`json:"logo"`                   							// 图标
	LongLat  		string 		`json:"long_lat"`               							// 经纬度
	Phone    		string 		`json:"phone"`                  							// 联系电话
	Remarks  		string 		`json:"remarks"`                							// 描述信息
	Valid    		int64  		`json:"valid"`                  							// 有效期，结束时间戳
	CreatedAt 	int64		  `json:"created_at"`        		// 创建时间
	UpdatedAt 	int64		  `json:"updated_at"`        		// 更新时间
}

// 餐桌
type Table struct {
	Id 					int64 		`json:"id"`
	Name 				string 		`json:"name"`
	ShopId 			int64 		`json:"shop_id"`
	IsDel 			bool 			`json:"is_del", xorm:"notnull default 0 ->"`
	CreatedAt 	int64 		`json:"created_at"`
	UpdatedAt 	int64 		`json:"updated_at"`
}



// 自动创建数据表
/*
func init(){
	orm := db.GetOrm()
	err := orm.Sync2(new(Category), new(Greens), new(Shop))
	if err != nil {
		log.Error("初始化数据库失败： ", err)
		panic(err)
	}
}
*/