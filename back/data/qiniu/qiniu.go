package qiniu

import (
	"OrderMeal/conf"
	"OrderMeal/data/log"
	"context"
	"OrderMeal/pub"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// 定义返回结构图
type PutRet struct {
	Key   string
	Hash  string
	Fsize int
	Name  string
}

var accessKey = conf.Cfg("qiniu", "accessKey")
var secretKey = conf.Cfg("qiniu", "secretKey")
var bucket = conf.Cfg("qiniu", "bucket")
var prefix = conf.Cfg("qiniu", "prefix")

// 文件上传，文件路径
func Uploader(fpath string) *PutRet {
	date := pub.GetDate()
	// 加密文件名
	fname := fmt.Sprintf("%s/%s/%s", prefix, date, pub.GetMd5(pub.GetSecond()))
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key": "$(key)", "hash": "$(etag)", "fsize": $(fsize), "name": "$(name)"}`,
	}
	// 凭证有效期，默认1小时
	// putPolicy.Expires = 600
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	// cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	// cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	// cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"name": fname,
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, fname, fpath, &putExtra)
	if err != nil {
		log.Error("文件保存失败：", err)
		return nil
	}
	return &ret
}

// 空间文件删除
func Deleter(fname string) bool {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(bucket, fname)
	if err != nil {
		log.Error("删除文件失败：", err)
		return false
	}
	return true
}
