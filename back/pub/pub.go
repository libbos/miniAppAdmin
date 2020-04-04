package pub

import (
	"github.com/gin-gonic/gin"
	"OrderMeal/data/redis"
	"OrderMeal/data/log"
	"strconv"
)

// 分页类型
type PageSize struct {
	Page 		int    `json:"page"`
	Size 		int    `json:"size"`
	Sear 		string `json:"sear"`
	ShopId 	int64
}

func Fail(msg string) gin.H {
	return gin.H{"code": 1000, "msg": msg}
}

func Success(data interface{}) gin.H {
	return gin.H{"code": 0, "data": data}
}


// 获取redis缓存
func GetRedisShopId(c *gin.Context) int64 {
	Token := c.Request.Header["Token"]
	if len(Token) == 0 {
		log.Debug("Token不存在")
		panic("header token not exist!")
	}

	data, err := redis.Get(Token[0])
	if err != nil {
		log.Debug("获取redis失败：", err)
		panic(err)
	}
	sid, err := strconv.ParseInt(data, 0, 64)
	if err != nil {
		log.Error("转换错误：", err)
		panic(err)
	}
	return sid
}
