/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2023/2/4 20:55
 * @Project_Name : scaffold
 * @File : mysql.go
 * @Software :GoLand
 */

package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("MySQL init failed", zap.Error(err))
		return
	}

	err = db.Ping()
	if err != nil {
		zap.L().Error("MySQL ping failed", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conn"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conn"))
	fmt.Println("MySQL init success")
	return
}
func Close() {
	_ = db.Close()
}
