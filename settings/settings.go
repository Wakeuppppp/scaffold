/*
 * -*- coding = utf-8 -*-
 * Author: _谷安
 * @Time : 2023/2/4 20:21
 * @Project_Name : scaffold
 * @File : settings.go
 * @Software :GoLand
 */

package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	fmt.Println("Config init success")
	return
}
