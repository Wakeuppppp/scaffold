/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2023/2/4 20:02
 * @Project_Name : scaffold
 * @File : main.go
 * @Software :GoLand
 */

package main

import (
	"context"
	"demo/dao/mysql"
	"demo/dao/redis"
	"demo/logger"
	"demo/router"
	"demo/settings"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("config init failed:", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Println("logger init failed:", err)
		return
	}
	zap.L().Debug("logger init success")
	zap.L().Sync()

	// 3. 初始化 MySQL
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Println("MySQL init failed:", err)
		return
	}
	zap.L().Debug("MySQL init success")
	defer mysql.Close()

	// 4. 初始化 Redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Println("Redis init failed:", err)
		return
	}
	zap.L().Debug("Redis init success")
	defer redis.Close()

	// 5. 注册路由
	r := router.Setup()

	// 6. 启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
}
