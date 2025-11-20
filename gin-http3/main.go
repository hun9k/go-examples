package main

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func main() {
	// 1. 初始化Handler（Gin 引擎）
	handler := gin.Default()

	// 2. 定义路由
	handler.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello "+c.Request.Proto)
	})

	// 证书和密钥文件地址
	certFile, keyFile := "localhost.pem", "localhost-key.pem"

	// 同时启动多服务监听，使用wg管理goroutine
	wg := &sync.WaitGroup{}

	// 启动 HTTP/3 服务QUIC(TLS)
	wg.Go(func() {
		defer wg.Done()
		if err := http3.ListenAndServeQUIC(":443", certFile, keyFile, handler); err != nil {
			slog.Error(err.Error())
		}
	})

	// 启动 HTTP/1.1/2 服务TCP+TLS
	wg.Go(func() {
		if err := http.ListenAndServeTLS(":443", certFile, keyFile, handler); err != nil {
			slog.Error(err.Error())
		}
	})

	// 启动 HTTP/1.1/2 服务TCP
	wg.Go(func() {
		if err := http.ListenAndServe(":80", handler); err != nil {
			slog.Error(err.Error())
		}
	})

	wg.Wait()
}
