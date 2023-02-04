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
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password:     viper.GetString("redis.password"),
		DB:           0,
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conn"),
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
