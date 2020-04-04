package common

import (
	// "OrderMeal/data/qiniu"
	"OrderMeal/pub"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"OrderMeal/data/log"
	"path/filepath"
)

// 暂时将文件保存在本地服务器
func UploadFile(c *gin.Context){
	file, err := c.FormFile("file")
	if err != nil {
		log.Warn("上传菜单失败：", err)
		c.JSON(http.StatusBadRequest, pub.Fail("上传图片失败"))
		return
	}
	fname := pub.GetMd5(file.Filename + pub.GetSecond())
	// 判断目录是否存在，不存在则创建目录
	fdir := filepath.Join("runtime", "public")
	if err := pub.DirExist(fdir); err != nil {
		pub.CreateDir(fdir)
	}
	fdir = filepath.Join(fdir, pub.GetMonth())
	if err := pub.DirExist(fdir); err != nil {
		pub.CreateDir(fdir)
	}
	dst := fmt.Sprintf("%s%s", fname, filepath.Ext(file.Filename))
	fdir = filepath.Join(fdir, dst)

	if err := c.SaveUploadedFile(file, fdir); err != nil {
		log.Warn(err)
		c.JSON(http.StatusBadRequest, pub.Fail("图片保存失败"))
		return
	}
	c.JSON(http.StatusOK, pub.Success(fdir))
}


