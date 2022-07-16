package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
	TODO
	1. 第三方http库gin
	2. 不同第三方库协调工作；使用zap记录关于http服务器的日志---中间件
**/
func main() {

	//===gin===
	// r := gin.Default()
	// r.GET("/request", func(ctx *gin.Context) { //为/request注册处理函数
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "reply", "error": "0"})
	// })
	// r.Run()

	//====using middleware====
	//使用zap统一记录日志（结构化日志）
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// r := gin.Default()
	r := gin.New()

	//使用中间件
	r.Use(func(ctx *gin.Context) { //Context实现go标准库的Context接口；中间件中可以向Context写内容，然后在后面的HandleFunc中使用
		// 记录访问路径、返回码、响应时间
		s := time.Now()
		// zapLogger.Info("incoming request", zap.String("path", ctx.Request.URL.Path), zap.Int("code", ctx.Writer.Status()))
		// zapLogger.Info("incoming request", ))
		//请求之前
		fmt.Println("Hello Before Next;")
		ctx.Next()
		fmt.Println("Hello After Next;")

		//请求之后
		// zapLogger.Info("incoming request", zap.Int("code", ctx.Writer.Status()))
		zapLogger.Info("incoming request", zap.String("path", ctx.Request.URL.Path), zap.Int("code", ctx.Writer.Status()), zap.Duration("elapsed", time.Since(s)))

	}, func(ctx *gin.Context) {
		//第二个HandleFunc
		ctx.Set("requestId", rand.Int())
		ctx.Next()

	})

	//设置HandleFunc
	r.GET("/request", func(ctx *gin.Context) {
		fmt.Println("Where am I; request")
		h := gin.H{
			"message": "reply",
		}
		if rid, exist := ctx.Get("requestId"); exist { //判断是否存在requestId字段
			h["requestId"] = rid
		}
		ctx.JSON(http.StatusOK, h)

		// ctx.JSON(http.StatusOK, gin.H{"message": "reply", "error": "0"})

	})
	//尝试取中间件写入的数据

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Gin")
		fmt.Println("Where am I; Hello")

	})
	r.Run()

}
