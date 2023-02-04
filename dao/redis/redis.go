/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2023/2/4 21:08
 * @Project_Name : scaffold
 * @File : redis.go
 * @Software :GoLand
 */

package redis

import (
	"demo/settings"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init(config *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConn,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("Redis ping failed", zap.Error(err))
		return
	}
	fmt.Println("Redis init success")
	return
}

func Close() {
	_ = rdb.Close()
}
