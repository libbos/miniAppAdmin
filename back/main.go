package main

import (
	"OrderMeal/route"
	"io"
	"os"
	"os/signal"
	"net/http"
	"context"
	"time"
	"github.com/gin-gonic/gin"
	"OrderMeal/data/log"
	Log "log"
)

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.SetMode(gin.DebugMode)
	r := route.Route()
	// 优雅重启服务
	srv := &http.Server{
		Addr: ":8888",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("listen: ", err)
			Log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	Log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		Log.Fatal("Server Shutdown: ", err)
	}

	Log.Println("Server exiting")
}
