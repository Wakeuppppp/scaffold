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
	"demo/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(config *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
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

	db.SetMaxOpenConns(config.MaxOpenConn)
	db.SetMaxIdleConns(config.MaxIdleConn)
	fmt.Println("MySQL init success")
	return
}
func Close() {
	_ = db.Close()
}
